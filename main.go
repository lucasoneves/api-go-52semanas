package main

import (
	"api/models"
	"api/routes"
	"fmt"
)

func main() {
	models.Player = []models.Players{
		{Name: "Weverton", History: "Historia", ShirtNumber: 21},
		{Name: "Piquerez", History: "História do Piquerez", ShirtNumber: 22},
	}
	fmt.Println("Server is listening on")
	routes.HandleRequest()
}
