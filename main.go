package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"go-gin/config"
	"go-gin/models"
	"go-gin/routes"
)

func main() {

	// Creating a connection to the database
	var err error
	c := config.NewDBConfig()
	config.DB, err = gorm.Open("mysql", c.URL())

	if err != nil {
		fmt.Println("statuse: ", err)
	}

	defer config.DB.Close()

	// run the migrations: todo struct
	config.DB.AutoMigrate(&models.Todo{})

	//setup routes
	r := routes.SetupRouter()

	// running
	r.Run()
}
