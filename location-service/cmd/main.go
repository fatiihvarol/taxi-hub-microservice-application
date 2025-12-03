package main

import (
    "github.com/gofiber/fiber/v2"
    "location-service/config"
    "location-service/controllers"
    "location-service/routes"
    "location-service/services"
    redisrepo "location-service/infrastructure/redis"
)

func main() {
    config.LoadEnv()
    config.ConnectRedis()

    repo := redisrepo.NewRedisLocationRepository()
    svc := services.NewLocationService(repo)
    ctrl := controllers.NewLocationController(svc)

    app := fiber.New()

    // REST ve WebSocket route’larını ekle
    routes.SetupRoutes(app, ctrl, svc)

    app.Listen(":" + config.AppPort)
}
