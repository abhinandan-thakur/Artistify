package database

import (
	"context"
	// "fmt"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// func ConnectDB() (*pgx.Conn, error) {

// }

func ConnectDB() (*pgx.Conn, error){

	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	var DB_URL string

	DB_URL = "postgres://"+os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+"/"+os.Getenv("DB_NAME")

	conn, err := pgx.Connect(
		context.Background(),
		DB_URL,
	)

	if err != nil {
		return nil, err
	}

	return conn, err
}
