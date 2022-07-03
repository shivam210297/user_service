package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"os/signal"
	"syscall"
	"user-service/server"
)

func main() {

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	srv := server.SrvInit()
	go srv.Start()

	<-done
	logrus.Info("Graceful shutdown")
	srv.Stop()

}
