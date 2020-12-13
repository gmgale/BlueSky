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
	"github.com/gmgale/BlueSky/ratelimit"

	"github.com/gorilla/mux"
)

func main() {
	var flagPort string
	var flagHost string

	flag.StringVar(&flagPort, "port", "9090", "Port for server setup.")
	flag.StringVar(&flagHost, "host", "localhost", "Host IP for server setup.")
	flag.StringVar(&ratelimit.GlobalRateLimit, "limit", "-1", "Rate limit per minute.")

	flag.Parse()

	if ratelimit.GlobalRateLimit == "-1" {
		fmt.Printf("WARNING: Rate-limiting is switched off.\n")
		fmt.Printf("Use commang line flag '-limit' to set.\n")
	}
	sm := mux.NewRouter()
	err := os.MkdirAll("data", os.ModePerm)
	if err != nil {
		fmt.Printf("Error building data folder - try running as administrator.")
		fmt.Printf("%v\n", err)
		fmt.Printf("Warning: Rate limiting may be disabled.")
	}

	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/currentweather/{city}/{imgSize:[a-z]+}", handlers.GetImage)
	getRouter.Use(handlers.RatelimitMiddleware)
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
		err = s.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	log.Println("Received terminate, gracefully shutting down", sig)

	// Clean up when shutting down
	err = os.RemoveAll("data")
	if err != nil {
		fmt.Printf("Error removing data folder.")
		fmt.Printf("%v\n", err)
	}

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)

	s.Shutdown(tc)
}
