package main

import (
    "auth-service/config"
    "auth-service/database"
    "auth-service/routes"
    "github.com/gofiber/fiber/v2"
    _ "auth-service/docs"
    swagger "github.com/gofiber/swagger"
)

func main() {
    // Env ve DB
    config.LoadEnv()
    database.ConnectMongo()

    app := fiber.New()

    // Swagger UI
    app.Get("/docs/*", swagger.HandlerDefault)

    // Auth routes
    routes.AuthRoutes(app)

    app.Listen(":" + config.GetEnv("PORT"))
}
