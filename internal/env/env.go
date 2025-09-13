package env

import (
	"os"
	"strconv"
)

func GetEnvString(key, defulatValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defulatValue
}

func GetEnvInt(key string, defulatValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defulatValue
}
