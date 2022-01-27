package config

import (
	"log"
	"os"
	"sync"
)

const (
	DefaultHostPort = "8080"
	DefaultHostName = "localhost"
)

type Config struct {
	HostPort string
	HostName string
}

var instance *Config
var once sync.Once

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{
			HostPort: getEnv("HOST_PORT", DefaultHostPort),
			HostName: getEnv("HOST_NAME", DefaultHostName),
		}
	})

	return instance
}

func getEnv(key string, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		value = defaultValue
	}
	log.Printf("Environment Variable: %s, %s", key, value)

	return value
}
