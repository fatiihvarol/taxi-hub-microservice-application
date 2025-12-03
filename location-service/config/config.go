package config

import (
    "fmt"
    "log"
    "os"
    "github.com/joho/godotenv"
    "github.com/go-redis/redis/v8"
)

var RedisClient *redis.Client
var AppPort string
var NearbyRadius float64
var NearbyCount int

func LoadEnv() {
    err := godotenv.Load()
    if err != nil {
        log.Println(".env file not found, using system environment variables")
    }

    AppPort = getEnv("APP_PORT", "8081")
    NearbyRadius = getEnvFloat("NEARBY_RADIUS_KM", 5)
    NearbyCount = getEnvInt("NEARBY_COUNT", 20)
}

func ConnectRedis() {
    RedisClient = redis.NewClient(&redis.Options{
        Addr:     fmt.Sprintf("%s:%s", getEnv("REDIS_HOST", "localhost"), getEnv("REDIS_PORT", "6379")),
        Password: getEnv("REDIS_PASSWORD", ""),
        DB:       getEnvInt("REDIS_DB", 0),
    })
}

func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}

func getEnvInt(key string, fallback int) int {
    if value, exists := os.LookupEnv(key); exists {
        var v int
        fmt.Sscanf(value, "%d", &v)
        return v
    }
    return fallback
}

func getEnvFloat(key string, fallback float64) float64 {
    if value, exists := os.LookupEnv(key); exists {
        var v float64
        fmt.Sscanf(value, "%f", &v)
        return v
    }
    return fallback
}
