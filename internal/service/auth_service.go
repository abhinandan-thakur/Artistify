package service

import (
	"artistify/internal/models"
	"artistify/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Register(pool *pgxpool.Pool, user models.Users) (models.Users, error) {

	user, err := repository.Register(pool, user)

	if err != nil {
		return models.Users{}, nil
	}

	return user, nil
}

func RegisterWithRole(pool *pgxpool.Pool, user models.Users) (models.Users, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.Users{}, nil
	}

	user.Password = string(hashedPassword)

	user, err = repository.RegisterWithRole(pool, user)

	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}

func Login(pool *pgxpool.Pool, input models.Users) (string, string, error) {

	user, err := repository.Login(pool, input)

	if err != nil {
		return "", "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))

	if err != nil {
		return "", "", err
	}

	expires := time.Now().Add(time.Hour)
	claims := jwt.MapClaims{
		"id":   user.ID,
		"role": user.Type,
		"exp":  expires.Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("super-secret-key"))

	return tokenString, user.Type, nil
}
