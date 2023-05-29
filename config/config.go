package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	//DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	//RealiseMode indicates service mode is release
	RealiseMode = "realise"
)

type Config struct {
	ServiceName string
	ServiceHost string 
	ServicePort string

	Environment string //debug, test, realise 
	Version string

	OrderServiceHost string
	OrderServicePort string

	PostgresHost string
	PostgresPort string 
	PostgresUser string
	PostgresPassword string
	PostgresDatabase string

	PostgresMaxConnections int32
}


func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME ", "golang_psql"))
	config.ServiceHost = cast.ToString(getOrReturnDefaultValue("SERVICE_HOST", "localhost"))
	config.ServicePort = cast.ToString(getOrReturnDefaultValue("SERVICE_PORT", ":8081"))

	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.OrderServiceHost = "localhost"

	config.OrderServicePort = cast.ToString(getOrReturnDefaultValue("ORDER_SERVICE_PORT", ":8080"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToString(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "islombek"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "lib1234"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "bicycle"))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))


	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}