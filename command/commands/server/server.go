package server

import (
	"context"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"time"
	"user_center/app"
	"user_center/command"
	"user_center/config"
)

var CMDuimsServer = &command.Command{
	UsageLine: "server [port]",
	Short:     "创建一个UIMS HTTP API服务",
	Long: `
server 子命令会创建一个UIMS HTTP API服务应用程序。
`,
	Run:    createUIMSapiServer,
}

var (
	//address string
	port string

	g errgroup.Group
)

func init() {
	//CMDuimsServer.Flag.StringVar(&address, "a", "0.0.0.0", "Listen address")
	CMDuimsServer.Flag.StringVar(&port, "p", "", "Listen port")
	command.CMD.Register(CMDuimsServer)
}

// createUIMSapiServer 创建一个 UIMS API 服务器
func createUIMSapiServer(cmd *command.Command, args []string) int {
	var err error
	if len(args) > 0 {
		err = cmd.Flag.Parse(args[1:])
		if err != nil {
			log.Println(err)
		}
	}
	if len(port) == 0 {
		port = config.Host
	} else {
		port = ":" + strings.TrimPrefix(port, ":")
	}

	router1 := app.GetEngineRouter()
	server1 := &http.Server{
		Addr:         port,
		Handler:      router1,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	g.Go(func() error {
		return server1.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		cmd.Error(err.Error())
		log.Fatal(err)
	}

	//err = router.Run(port)
	//if err != nil {
	//	cmd.Error(err.Error())
	//}

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	cmd.Info("Shutdown uims api server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server1.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")

	return 0
}
