package config

import (
	"log"
	"os"
	"strconv"

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

	GetRateLimit float64
	GetRateRefill float64
	PostRateLimit float64
	PostRateRefill float64
	DeleteRateLimit float64
	DeleteRateRefill float64
	PutRateLimit float64
	PutRateRefill float64
}

func LoadConfig() *Config {

	env := os.Getenv("APP_ENV")
	
	if env == "" { env = "local"}
	err := godotenv.Load(".env." + env)

	if err != nil { log.Println("No env file found")}
	
	if (os.Getenv("JWT_SECRET") == "") {log.Fatal("No JWT Secret key")}

	url := "http://"+os.Getenv("URL_HOST")+":"+os.Getenv("URL_PORT")

	getRateLimit, _ := strconv.ParseFloat(os.Getenv("GET_RATE_LIMIT"),64)
	getRateRefill, _ := strconv.ParseFloat(os.Getenv("GET_RATE_REFILL"),64)
	postRateLimit, _ := strconv.ParseFloat(os.Getenv("POST_RATE_LIMIT"),64)
	postRateRefill, _ := strconv.ParseFloat(os.Getenv("POST_RATE_REFILL"),64)
	deleteRateLimit, _ := strconv.ParseFloat(os.Getenv("DELETE_RATE_LIMIT"),64)
	deleteRateRefill, _ := strconv.ParseFloat(os.Getenv("DELETE_RATE_REFILL"),64)
	putRateLimit, _ := strconv.ParseFloat(os.Getenv("PUT_RATE_LIMIT"),64)
	putRateRefill, _ := strconv.ParseFloat(os.Getenv("PUT_RATE_REFILL"),64)

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


		GetRateLimit: getRateLimit,
		GetRateRefill: getRateRefill,
		PostRateLimit: postRateLimit,
		PostRateRefill: postRateRefill,
		DeleteRateLimit: deleteRateLimit,
		DeleteRateRefill: deleteRateRefill,
		PutRateLimit: putRateLimit,
		PutRateRefill: putRateRefill,
	}
}