package main

import (
	"log"

	"github.com/bhumong/lemonilo/app/database"
	"github.com/bhumong/lemonilo/app/router"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
)

func main() {
	var app = fiber.New()
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}
	router.SetupApi(app)
	app.Listen("127.0.0.1:8000")
}
