package service

import (
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/models"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/repository"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
	"golang.org/x/crypto/bcrypt"
	"time"
	// "log"
	"crypto/rand"
	"math/big"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/database"
	"github.com/abhinandan-thakur/Artistify/auth-service/internal/config"
)

func GenerateOTP(user models.Users, rdb *redis.Client) (string, error) {

	const digits = "0123456789"

	ret := make([]byte, 6)

	for i := 0; i < 6; i++ {

		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits)),),)

		if err != nil {
			return "", err
		}
		ret[i] = digits[num.Int64()]
	}

	otp :=  string(ret)
	cacheKey := "otp:"+user.Email

	err := rdb.Set(database.Ctx, cacheKey, otp, 5*time.Minute).Err()
	if err != nil {
		return "", err
	}
	return otp, nil
}

func Register(pool *pgxpool.Pool, user models.Users) (models.Users, error) {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return models.Users{}, nil
	}

	user.Password = string(hashedPassword)

	user, err = repository.Register(pool, user)

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
	config := config.LoadConfig()
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
	tokenString, err := token.SignedString([]byte(config.JWTSecret))

	if err != nil {
		return "", "", err
	}

	return tokenString, user.Type, nil
}

