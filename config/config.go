package config

import (
	"go-clean-arhitecture/constants"
	"os"

	"github.com/joho/godotenv"
)

var (
	// App Configuration
	APPUrl    = GetEnv("APP_URL")
	APPPort   = GetEnv("APP_PORT")
	APPPrefix = GetEnv("APP_PREFIX")
	APPKey    = GetEnv("APP_KEY")

	// Database Configuration
	DBDriver = GetEnv("DB_DRIVER")
	DBName   = GetEnv("DB_NAME")
	DBHost   = GetEnv("DB_HOST")
	DBPort   = GetEnv("DB_PORT")
	DBUser   = GetEnv("DB_USER")
	DBPass   = GetEnv("DB_PASS")
	SSLMode  = GetEnv("SSL_MODE")

	MONGOHost = GetEnv("MONGO_HOST")
	MONGOPort = GetEnv("MONGO_PORT")
	MONGODb   = GetEnv("MONGO_DB")

	// Key Configuration
	JWTKey = GetEnv("JWT_KEY")
	EKey   = GetEnv("E_KEY")
)

func GetEnv(key string, value ...string) string {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error load .env file")
	}

	if os.Getenv(key) != constants.EMPTY_VALUE {
		return os.Getenv(key)
	} else {
		if len(value) > constants.EMPTY_VALUE_INT {
			return value[constants.EMPTY_VALUE_INT]
		}
		return constants.EMPTY_VALUE
	}
}
