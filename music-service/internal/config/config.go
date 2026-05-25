package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv string

	DBHost string
	DBPort string
	DBUser string
	DBName string
	DBPassword string

	URLPort string
	URLHost string
	
	RedisHost string
	RedisPort string
	
	GRPCHost string
	GRPCPort string
	
	JWTSecret string
	
	URL string
}

func LoadConfig() *Config {

	env := os.Getenv("APP_ENV")
	
	if env == "" { env = "local"}
	err := godotenv.Load(".env." + env)

	if err != nil { log.Println("No env file found")}
	
	if (os.Getenv("JWT_SECRET") == "") {log.Fatal("No JWT Secret key")}

	url := "http://"+os.Getenv("URL_HOST")+":"+os.Getenv("URL_PORT")
	return &Config{
		AppEnv: env,

		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword:	os.Getenv("DB_PASSWORD"),
		DBName:	os.Getenv("DB_NAME"),

		RedisHost: os.Getenv("REDIS_HOST"),
		RedisPort: os.Getenv("REDIS_PORT"),

		GRPCPort: os.Getenv("GRPC_PORT"),
		GRPCHost: os.Getenv("GRPC_HOST"),

		JWTSecret: os.Getenv("JWT_SECRET"),
		
		URLPort: os.Getenv("URL_PORT"),
		URLHost: os.Getenv("URL_HOST"),
		URL: url,
	}
}