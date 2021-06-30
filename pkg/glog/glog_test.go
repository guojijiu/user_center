package glog_test

import (
	"log"
	"testing"
	"user_center/boot"
	"user_center/config"
	"user_center/pkg/glog"
)

func TestMain(m *testing.M) {
	boot.SetInTest()
	boot.Boot()
	m.Run()
}

func TestLog(t *testing.T) {
	log.Print("log by log")
	glog.Default().Print("log by glog.Default()")
	glog.Channel("gin").Print("log by glog.Channel(\"gin\")")
}

func TestAllChannel(t *testing.T) {
	for name := range config.Logging.Channels {
		glog.Channel(name).Printf("Test")
	}
}
