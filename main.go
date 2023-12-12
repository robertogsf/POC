package main

import (
	"github.com/robertogsf/POC/database"
	"github.com/robertogsf/POC/handlers"
)

func main() {
	database.ConnectDB()
	handlers.Handlerr()
}
