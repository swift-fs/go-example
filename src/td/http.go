package td

import (
	"os"

	"github.com/gofiber/fiber/v2"
)

func HelloHandler(c *fiber.Ctx) error {
	name := c.Query("name")
	os.WriteFile("abc.txt", []byte("123"), 0666)
	if name == "" {
		return c.JSON(fiber.Map{"message": "Hello, World!"})
	}
	return c.JSON(fiber.Map{"message": "Hello, World!", "name": name})
}

func SetupWebServer() *fiber.App {
	app := fiber.New()
	app.Get("/hello", HelloHandler)
	return app
}
