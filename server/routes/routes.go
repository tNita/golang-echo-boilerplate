package routes

import (
	"echo-demo-project/server"
	"echo-demo-project/server/handlers"
	"github.com/labstack/echo"
	"net/http"
)

func ConfigureRoutes(server *server.Server) {
	postHandler := handlers.NewPostHandler(server)

	server.Echo.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	server.Echo.POST("/posts", postHandler.CreatePost())
	server.Echo.GET("/posts", postHandler.GetPosts())
	server.Echo.DELETE("/posts/:id", postHandler.DeletePost())
	server.Echo.PUT("/posts/:id", postHandler.UpdatePost())
}
