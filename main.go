package main

import (
	"github.com/jonathantx/api-go-gin/database"
	"github.com/jonathantx/api-go-gin/routes"
)

func main() {

	database.ConnectionDB()

	routes.HandleRequests()
}
