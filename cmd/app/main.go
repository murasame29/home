package main

import (
	"log"

	"github.com/murasame29/home/internal/router"
	"github.com/murasame29/home/internal/server"
	"github.com/sirupsen/logrus"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	logger := logrus.New()

	handler := router.New(logger)

	server.New(handler).RunWithGracefulShutdown()
	return nil
}
