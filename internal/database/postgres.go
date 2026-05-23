package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDB() (*pgxpool.Pool, error) {

	env := os.Getenv("APP_ENV")

	if env == "" {
		env = "local"
	}

	envFile := ".env." + env

	err := godotenv.Load(envFile)

	if err != nil {
		return nil, err
	}

	var DB_URL string

	var pool *pgxpool.Pool

	DB_URL = "postgres://" +
		os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + "/" +
		os.Getenv("DB_NAME") +
		"?sslmode=disable"

	for i := 0; i < 10; i++ {

		pool, err = pgxpool.New(
			context.Background(),
			DB_URL,
		)

		if err == nil {
			fmt.Println("Connected to Database!")
			return pool, nil
		}
		fmt.Println("Retrying to connect to Database...")
		time.Sleep(3 * time.Second)
	}

	return nil, err

}
