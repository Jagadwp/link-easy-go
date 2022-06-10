package main

import (
	"github.com/Jagadwp/link-easy-go/db"
	"github.com/Jagadwp/link-easy-go/internal/controller"
)

func main() {

	db.DatabaseInit()

	e := controller.Init()

	e.Logger.Fatal(e.Start(":8080"))

}
