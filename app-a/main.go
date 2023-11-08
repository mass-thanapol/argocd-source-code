package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	app.Get("/hello-a/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		response := map[string]string{"status": "success", "result": name + " (a)"}
		fmt.Println(response)
		return c.JSON(response)
	})
	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}
}
