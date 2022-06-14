package routes

import (
	"github.com/Jagadwp/link-easy-go/internal/controllers"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
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
	user := e.Group("users")
	user.POST("", userController.InsertUser)
	user.GET("", userController.GetAllUsers)
	// user.GET("/:id", userController.FindUserByID)
	// user.PUT("/:id", userController.UpdateUser)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
	e.GET("/", func(c echo.Context) error {
		return dto.MessageResponse(c, 200, "When Link is Easy!")
	})

	// e.GET("/test", userController.test)
}

func UrlUserPath(e *echo.Echo, urlController *controllers.UrlController) {
	if urlController == nil {
		panic("Controller parameter cannot be nil")
	}

	e.POST("/urls", urlController.InsertUrl)
	e.GET("/urls/:user_id", urlController.GetAllUrlsByUserID)
}
