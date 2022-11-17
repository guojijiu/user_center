package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/sirupsen/logrus"
	"time"
	conf "user_center/config"
)

var connections = map[string]*redis.Client{}

func InitDef() (*redis.Client, error) {
	return initConn(conf.Database.Redis["default"])
}

// Def Get default connection
func Def() *redis.Client {
	return Conn("default")
}

// Conn Get connection
func Conn(name string) *redis.Client {
	if client, ok := connections[name]; ok {
		return client
	}
	if c, ok := conf.Database.Redis[name]; ok {
		connections[name], _ = initConn(c)
		return connections[name]
	}
	logrus.Panicf("Not config redis connection: %s", name)
	return Def()
}

func initConn(c conf.RedisConf) (*redis.Client, error) {
	var client *redis.Client
	client = newConn(fmt.Sprintf("%s:%d", c.Host, c.Port), c.Password, c.Database)
	_, err := client.Ping().Result()
	if err != nil {
		return client, err
	}
	return client, nil
}

func newConn(addr string, password string, db int) *redis.Client {
	var options = &redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	}
	client := redis.NewClient(options)

	err := eventually(func() error {
		return client.Ping().Err()
	}, 30*time.Second)
	if err != nil {
		logrus.Printf("redis conn err: %+v", err)
		return nil
	}

	return client
}

func Close() {
	for k, conn := range connections {
		if err := conn.Close(); err != nil {
			logrus.Printf("Close redis conn %s err: %+v", k, err)
		}
	}
}

func eventually(fn func() error, timeout time.Duration) error {
	errCh := make(chan error, 1)
	done := make(chan struct{})
	exit := make(chan struct{})

	go func() {
		for {
			err := fn()
			if err == nil {
				close(done)
				return
			}

			select {
			case errCh <- err:
			default:
			}

			select {
			case <-exit:
				return
			case <-time.After(timeout / 100):
			}
		}
	}()

	select {
	case <-done:
		return nil
	case <-time.After(timeout):
		close(exit)
		select {
		case err := <-errCh:
			return err
		default:
			return fmt.Errorf("timeout after %s without an error", timeout)
		}
	}
}
