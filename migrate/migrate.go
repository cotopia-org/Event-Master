package main

import (
	"github.com/cotopia-org/Event-Master/initializers"
	"github.com/cotopia-org/Event-Master/models"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectToBD()
}

func main() {
	initializers.DB.AutoMigrate(&models.User{})
	initializers.DB.AutoMigrate(&models.Event{})
}