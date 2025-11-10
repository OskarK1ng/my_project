package main

import (
	"log"
	"net/http"
	"transactions/internal/config"
	"transactions/internal/db"
	"transactions/internal/handlers"
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

	http.HandleFunc("/deposit", handlers.DepositHandler)
	http.HandleFunc("/withdraw", handlers.WithdrawHandler)

	port := cfg.ServerPort
	log.Printf("Starting server on localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
