package main

import (
	"github.com/paulacgates/api-rest-go/database"
	"github.com/paulacgates/api-rest-go/routes"
)

func main() {
	database.ConnectDb()
	println("Iniciando")
	routes.HandleRequest()
}
