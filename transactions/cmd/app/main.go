package main

import (
	"log"
	"net/http"
	"transactions/internal/config"
	"transactions/internal/db"
	"transactions/internal/handlers"
	"transactions/internal/middleware"
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

	http.HandleFunc("/deposit", middleware.AuthMiddleware(cfg, handlers.DepositHandler))
	http.HandleFunc("/withdraw", middleware.AuthMiddleware(cfg, handlers.WithdrawHandler))
	http.HandleFunc("/getBalance", middleware.AuthMiddleware(cfg, handlers.GetBalanceHandler))

	port := cfg.ServerPort
	log.Printf("Starting server on localhost%s", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
