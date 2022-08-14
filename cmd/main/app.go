package main

import (
	"github.com/julienschmidt/httprouter"
	"golang-rest-api/internal/user"
	"golang-rest-api/pkg/logging"
	"net"
	"net/http"
	"time"
)

func main() {
	logger := logging.GetLogger()
	logger.Info("create router")
	router := httprouter.New()

	logger.Info("register user handler")
	handler := user.NewHandler(*logger)
	handler.Register(router)

	start(router)
}

func start(router *httprouter.Router) {
	logger := logging.GetLogger()

	logger.Info("start application")

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
	}

	logger.Info("server is listening 0.0.0.0:1234")
	logger.Fatal(server.Serve(listener))
}
