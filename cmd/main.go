package main

import "github.com/gofiber/fiber/v2"

func main() {
    app := fiber.New()

    app.Get("/", func(c *fiber.Ctx) error {
        return c.SendString("This is an amazing app. Enter your name below and let me give you your match")
    })

    app.Listen(":3000")
}
