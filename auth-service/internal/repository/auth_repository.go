package repository

import (
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/models"
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
)

func Register(pool *pgxpool.Pool, user models.Users) (models.Users, error) {

	err := pool.QueryRow(context.Background(),
		"INSERT INTO users(username, email, password, type) VALUES($1, $2, $3, $4) RETURNING id, created_at",
		user.Username,
		user.Email,
		user.Password,
		user.Type,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return models.Users{}, nil
	}

	return user, nil
}

func RegisterWithRole(pool *pgxpool.Pool, user models.Users) (models.Users, error) {

	err := pool.QueryRow(
		context.Background(),
		"INSERT INTO users(username, email, password, type) VALUES($1, $2, $3, $4) RETURNING id, created_at",
		user.Username,
		user.Email,
		user.Password,
		user.Type,
	).Scan(
		&user.ID,
		&user.CreatedAt,
	)

	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}

func Login(pool *pgxpool.Pool, user models.Users) (models.Users, error) {
	log.Println("Enter REpository")
	err := pool.QueryRow(context.Background(),
		"SELECT id, username, password, type FROM users WHERE email = $1",
		user.Email,
	).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
		&user.Type,
	)
	log.Println("what is the error in repository", err)

	if err != nil {
		return models.Users{}, err
	}


	return user, nil
}
