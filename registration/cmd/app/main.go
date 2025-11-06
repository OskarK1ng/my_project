package main

import (
	"log"
	"net/http"
	"registration/internal/config"
	"registration/internal/db"
	"registration/internal/handlers"
)

func main() {
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatalf("failed to init config: %v", err)
	}

	// подключаемся к БД
	err = db.ConnectDB(cfg)
	if err != nil {
		log.Fatalf("failed to connect to DB: %v", err)
	}

	defer db.DB.Close()

	http.HandleFunc("/register", handlers.RegisterHandler)

	port := cfg.ServerPort
	log.Printf("Starting server on localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
