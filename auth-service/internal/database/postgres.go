package database

import (
	"context"
	"log"
	"time"

	"github.com/abhinandan-thakur/Artistify/auth-service/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(config *config.Config) (*pgxpool.Pool, error) {

	var (
		pool   *pgxpool.Pool
		err    error
		DB_URL string
	)

	DB_URL = "postgres://" +
		config.DBUser + ":" +
		config.DBPassword + "@" +
		config.DBHost + ":" +
		config.DBPort + "/" +
		config.DBName + "?sslmode=disable"

	for i := 0; i < 10; i++ {
		pool, err = pgxpool.New(context.Background(), DB_URL)

		if err == nil {
			return pool, nil
		}

		log.Println("Retrying to connect to Database...")
		time.Sleep(3 * time.Second)
	}
	return nil, err
}
