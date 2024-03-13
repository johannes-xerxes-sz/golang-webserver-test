package main

// import (
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// func main() {
// 	r := gin.Default()
// 	r.GET("/", func(c *gin.Context) {
// 		c.String(http.StatusOK, "Hello, World!")
// 	})
// 	r.Run(":8080")
// }

// func handler(c *gin.Context) {
// 	c.String(http.StatusOK, "Hello, World!")
// }

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})

	app.Listen((":8080"))
}
