package main

import (
	"api-golang/controllers"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		panic("$PORT not set")
	}

	// Echo instance
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", controllers.HomePage)
	e.GET("/radio", controllers.Radio)

	// Start server
	e.Logger.Fatal(e.Start(":" + port))
}
