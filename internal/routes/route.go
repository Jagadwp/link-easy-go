package routes

import (
	"github.com/Jagadwp/link-easy-go/internal/controllers"
	"github.com/labstack/echo/v4"
)

func RegisterUserPath(e *echo.Echo, userController *controllers.UserController) {
	if userController == nil {
		panic("Controller parameter cannot be nil")
	}

	//authentication with Versioning endpoint
	// auth := e.Group("auth")
	// auth.POST("/login", authController.Login)

	//user with Versioning endpoint
	// user := e.Group("users")
	// user.POST("", userController.InsertUser)
	// user.GET("", userController.FindAllUser)
	// user.GET("/:id", userController.FindUserByID)
	// user.PUT("/:id", userController.UpdateUser)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
}