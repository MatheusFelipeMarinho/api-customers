package main

import (
	"api-gin/database"
	"api-gin/routes"
)


func main() {
	database.ConectaBanco()
	routes.HandleRequests()
}

