package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var server = echo.New()

// Start starts the rest service
func Start(address string) error {
	return server.Start(address)
}

func init() {
	server.HideBanner = true

	// middlewares
	server.Use(middleware.Logger())
	server.Use(middleware.Recover())

	// routes
	server.GET("/about", getAbout)
	server.GET("/tasks", getTaskList)
	server.POST("/tasks", createTask)
	server.PUT("/tasks/:id", updateTask)
	server.PUT("/tasks/:id/done", updateTaskDone)
}
