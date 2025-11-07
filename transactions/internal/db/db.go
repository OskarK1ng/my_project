package db

import (
	"bank_account_api/internal/config"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB(c *config.Config) error {
	dns := fmt.Sprintf(
		"postgresql://%v:%v@%v:%v/%v?sslmode=disable",
		c.DBUser, c.DBPassword, c.DBHost, c.DBPort, c.DBName,
	)

	var err error
	DB, err = pgxpool.New(context.Background(), dns)
	if err != nil {
		return fmt.Errorf("error creating connection pool: %w", err)
	}

	if err = DB.Ping(context.Background()); err != nil {
		return fmt.Errorf("unable to ping database: %w", err)
	}

	log.Println("Connected to PostgreSQL")

	return nil
}
