package tests

import (
	"github.com/Jagadwp/link-easy-go/db"
	"github.com/Jagadwp/link-easy-go/internal/controllers"
	"github.com/Jagadwp/link-easy-go/internal/repositories"
	"github.com/Jagadwp/link-easy-go/internal/routes"
	"github.com/Jagadwp/link-easy-go/internal/services"
	echo "github.com/labstack/echo/v4"
)

func InitHandler() *echo.Echo {
	db.DatabaseInit()

	DB := db.DB()

	usersRepo := repositories.NewUserRepository(DB)
	userService := services.NewUserService(usersRepo)
	userController := controllers.NewUserController(userService)

	urlsRepo := repositories.NewUrlRepository(DB)
	urlService := services.NewUrlService(urlsRepo)
	urlController := controllers.NewUrlController(urlService, userService)

	// create echo http
	e := echo.New()

	// register API path and handler
	routes.RegisterUserPath(e, userController)
	routes.UrlUserPath(e, urlController)
	routes.PublicPath(e, urlController)

	return e
}
