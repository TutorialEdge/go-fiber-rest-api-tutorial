package main

import (
	"github.com/elliotforbes/go-fiber-tutorial/internal/database"
	"github.com/elliotforbes/go-fiber-tutorial/internal/transport"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func main() {
	app := transport.Setup()
	database.InitDatabase()
	app.Listen(3000)
}
