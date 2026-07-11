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

	MailingGRPCHost string
	MailingGRPCPort string
	
	JWTSecret string
	
	URL string

	LoginRateLimit float64
	LoginRateRefil float64
	RegisterRateLimit float64
	ReigsterRateRefil float64
	AdminRateLimit float64
	AdminRateRefil float64
}

func LoadConfig() *Config {

	env := os.Getenv("APP_ENV")
	
	if env == "" { env = "local"}
	err := godotenv.Load(".env." + env)

	if err != nil { log.Println("No env file found")}
	
	if (os.Getenv("JWT_SECRET") == "") {log.Fatal("No JWT Secret key")}

	url := "http://"+os.Getenv("URL_HOST")+":"+os.Getenv("URL_PORT")

	loginRateLimit, _ := strconv.ParseFloat(os.Getenv("LOGIN_RATE_LIMIT"),64)
	loginRateRefil, _ := strconv.ParseFloat(os.Getenv("LOGIN_RATE_REFIL"),64)
	registerRateLimit, _ := strconv.ParseFloat(os.Getenv("REGISTER_RATE_LIMIT"),64)
	registerRateRefil, _ := strconv.ParseFloat(os.Getenv("REGISTER_RATE_REFIL"),64)
	adminRateLimit, _ := strconv.ParseFloat(os.Getenv("ADMIN_RATE_LIMIT"),64)
	adminRateRefil, _ := strconv.ParseFloat(os.Getenv("ADMIN_RATE_REFIL"),64)

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

		MailingGRPCHost: os.Getenv("MAILING_GRPC_HOST"),
		MailingGRPCPort: os.Getenv("MAILING_GRPC_PORT"),

		JWTSecret: os.Getenv("JWT_SECRET"),
		
		URLPort: os.Getenv("URL_PORT"),
		URLHost: os.Getenv("URL_HOST"),
		URL: url,

		LoginRateLimit: loginRateLimit,
		LoginRateRefil: loginRateRefil,

		RegisterRateLimit: registerRateLimit,
		ReigsterRateRefil: registerRateRefil,

		AdminRateLimit: adminRateLimit,
		AdminRateRefil: adminRateRefil,
	}
}