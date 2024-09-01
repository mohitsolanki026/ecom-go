package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	PublicHost string
	Port       string
	DBUser     string
	DBPassword string
	DBAddress  string
	DBName     string
	JWTExpirationInSecond int64
	JWTSecretKey string
}

var Envs = initConfig()

func initConfig() Config {
	godotenv.Load()
	return Config{
		PublicHost: getEnv("PUBLIC_HOST", "http://localhost"),
		Port:       getEnv("PORT", "8000"),
		DBUser:     getEnv("DB_USER", "root"),
		DBPassword: getEnv("DB_PASSWORD", "Mohit@1234"),
		DBAddress:  fmt.Sprintf("%s:%s", getEnv("DB_HOST", "127.0.0.1"), getEnv("DB_PORT", "3306")),
		DBName:     getEnv("DB_NAME", "ecom"),
		JWTExpirationInSecond: getEnvAsInt("JWT_EXPIRATION_IN_SECOND", 7*24*3600),
		JWTSecretKey: getEnv("JWTSecretKey","1234"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func getEnvAsInt(key string,fallback int64) int64 {
	if value, ok := os.LookupEnv(key); ok {
		i,err := strconv.ParseInt(value,10,64)
		if err != nil {
			return fallback
		}
		return i
	}
	return fallback
}