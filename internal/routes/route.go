package routes

import (
	"github.com/Jagadwp/link-easy-go/internal/controllers"
	"github.com/Jagadwp/link-easy-go/internal/middleware"
	"github.com/Jagadwp/link-easy-go/internal/shared/dto"
	"github.com/labstack/echo/v4"
)

func RegisterUserPath(e *echo.Echo, userController *controllers.UserController) {
	if userController == nil {
		panic("Controller parameter cannot be nil")
	}

	e.POST("/login", userController.Login)
	e.POST("/register", userController.Register)

	//user with Versioning endpoint
	user := e.Group("users")
	user.Use(middleware.IsAuthenticated)
	user.GET("", userController.GetAllUsers)
	user.GET("/:id", userController.GetUserById)
	user.PUT("/:id", userController.UpdateUser)
	user.DELETE("/:id", userController.DeleteUser)

	//health check
	e.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})

	e.GET("/", func(c echo.Context) error {
		return dto.SuccessResponse(c, 200, "When Link is Easy!", nil)
	})

	// e.GET("/test", userController.test)

}

func UrlUserPath(e *echo.Echo, urlController *controllers.UrlController) {
	if urlController == nil {
		panic("Controller parameter cannot be nil")
	}
	
	e.POST("urls/generate", urlController.CreateShortUrl)
	
	url := e.Group("urls")
	url.Use(middleware.IsAuthenticated)

	url.GET("", urlController.GetAllUrlsByUserID)
	url.POST("", urlController.InsertUrl)
	url.GET("/:id", urlController.GetUrlUserById)
	url.PUT("/:id", urlController.UpdateUrl)
	url.DELETE("/:id", urlController.DeleteUrl)
}

func PublicPath(e *echo.Echo, urlController *controllers.UrlController) {
	if urlController == nil {
		panic("Controller parameter cannot be nil")
	}

	public := e.Group("public")
	public.GET("/:short_link", urlController.GetUrlPublicByShortLink)
}