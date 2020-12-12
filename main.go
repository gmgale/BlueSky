package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gmgale/BlueSky/handlers"
	"github.com/gorilla/mux"
)

func main() {
	var flagPort string
	var flagHost string
	flag.StringVar(&flagPort, "port", "9090", "Port for server setup.")
	flag.StringVar(&flagHost, "host", "localhost", "Host IP for server setup.")
	flag.Parse()

	sm := mux.NewRouter()

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/currentweather/{city}/{imgSize:[a-z]+}", handlers.GetImage)
	getRouter.Use(handlers.WeatherMiddleware)

	fmt.Printf("Starting server at %s:%s\n", flagHost, flagPort)

	s := http.Server{
		Addr:         flagHost + ":" + flagPort,
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received terminate, gracefully shutting down", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
