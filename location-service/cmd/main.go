package main

import (
    "github.com/gofiber/fiber/v2"
    "location-service/config"
    "location-service/controller"
    "location-service/routes"
    "location-service/service"
    redisrepo "location-service/infrastructure/redis"
)

func main() {
    config.ConnectRedis()

    repo := redisrepo.NewRedisLocationRepository()
    service := service.NewLocationService(repo)
    controller := controller.NewLocationController(service)

    app := fiber.New()
    routes.SetupRoutes(app, controller)

    app.Listen(":8081")
}
