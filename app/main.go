package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/Jagadwp/link-easy-go/db"
	"github.com/Jagadwp/link-easy-go/internal/controllers"
	"github.com/Jagadwp/link-easy-go/internal/repositories"
	"github.com/Jagadwp/link-easy-go/internal/routes"
	"github.com/Jagadwp/link-easy-go/internal/services"
	"github.com/Jagadwp/link-easy-go/internal/shared/config"
	echo "github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {

	db.DatabaseInit()

	DB := db.DB()

	usersRepo := repositories.NewUserRepository(DB)
	userService := services.NewUserService(usersRepo)
	userController := controllers.NewUserController(userService)

	// create echo http
	e := echo.New()

	// register API path and handler
	routes.RegisterUserPath(e, userController)

	// run server
	go func() {
		address := fmt.Sprintf("localhost:%s", config.AppPort)

		if err := e.Start(address); err != nil {
			log.Info("shutting down the server")
		}
	}()

	// wait for interrupt signal to gracefully shutdown the server with
	// https://www.rudderstack.com/blog/implementing-graceful-shutdown-in-go/
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// a timeout of 10 seconds to shutdown the server
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
}
