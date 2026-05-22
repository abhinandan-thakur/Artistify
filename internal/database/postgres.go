package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// func ConnectDB() (*pgx.Conn, error) {

// }

func ConnectDB() (*pgx.Conn, error) {

	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	var DB_URL string

	var conn *pgx.Conn

	DB_URL = "postgres://" +
		os.Getenv("DB_USER") + ":" +
		os.Getenv("DB_PASSWORD") + "@" +
		os.Getenv("DB_HOST") + ":" +
		os.Getenv("DB_PORT") + "/" +
		os.Getenv("DB_NAME") +
		"?sslmode=disable"

	for i := 0; i < 10; i++ {

		conn, err = pgx.Connect(
			context.Background(),
			DB_URL,
		)

		if err == nil {
			fmt.Println("Connected to Database!")
			return conn, nil
		}
		fmt.Println("Retrying to connect to Database...")
		time.Sleep(3 * time.Second)
	}

	return nil, err

}
