package repository

import (
	"github.com/jackc/pgx/v5"
	"artistify/internal/models"
	"context"
)

func Register(conn *pgx.Conn, user models.Users) (models.Users, error) {
		
	err := conn.QueryRow(context.Background(),
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

func RegisterWithRole(conn *pgx.Conn, user models.Users) (models.Users, error) { 
		
		err := conn.QueryRow(
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


func Login(conn *pgx.Conn, user models.Users) (models.Users, error) {
			
	err := conn.QueryRow(context.Background(),
								"SELECT id, username, password, type FROM users WHERE email = $1",
								user.Email,
							).Scan(
									&user.ID,
									&user.Username,
									&user.Password,
									&user.Type,
						)
	
	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}

