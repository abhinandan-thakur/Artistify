package service

import (
	"artistify/internal/models"
	"artistify/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func Register(conn *pgx.Conn, user models.Users) (models.Users, error) {

	user, err := repository.Register(conn, user)

	if err != nil {
		return models.Users{}, nil
	}

	return user, nil
}

func RegisterWithRole(conn *pgx.Conn, user models.Users) (models.Users, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.Users{}, nil
	}

	user.Password = string(hashedPassword)

	user, err = repository.RegisterWithRole(conn, user)

	if err != nil {
		return models.Users{}, err
	}

	return user, nil
}

func Login(conn *pgx.Conn, input models.Users) (string, string, error) {

	user, err := repository.Login(conn, input)

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
