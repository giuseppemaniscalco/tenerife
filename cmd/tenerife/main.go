package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/giuseppemaniscalco/tenerife/internal/application"
	"github.com/giuseppemaniscalco/tenerife/internal/diagnostics"
)

func main() {
	logger := logrus.New()
	logger.SetOutput(os.Stdout)

	logger.Infof(
		"Starting the application %s %s %s",
		diagnostics.Version,
		diagnostics.Commit,
		diagnostics.Buildtime,
	)

	port := os.Getenv("PORT")
	if port == "" {
		logger.Fatal("Port is not provided")
	}

	r := mux.NewRouter()
	r.HandleFunc("/", application.HomeHandler(logger))
	r.HandleFunc("/healthz", diagnostics.LivenessHandler(logger))
	r.HandleFunc("/readyz", diagnostics.ReadynessHandler(logger))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	shutdown := make(chan error, 1)

	server := http.Server{
		Addr:    net.JoinHostPort("", port),
		Handler: r,
	}

	go func() {
		err := server.ListenAndServe()
		shutdown <- err
	}()

	select {
	case killSignal := <-interrupt:
		switch killSignal {
		case os.Interrupt:
			logger.Print("Got SIGINT...")
		case syscall.SIGTERM:
			logger.Print("Got SIGTERM...")
		}
	case <-shutdown:
		logger.Info("Got an error...")
	}

	logger.Infof("The service is stopping...")
	err := server.Shutdown(context.Background())
	if err != nil {
		logger.Info("Error on shutting down")
	}
}
