package main

import (
	"effective_mobile/config"
	"effective_mobile/controllers"
	"effective_mobile/models"
	"effective_mobile/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB
var err error

func main() {
	config.InitConfig()

	db, err = gorm.Open("postgres", config.DBConnectionString)
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}
	defer db.Close()

	db.AutoMigrate(&models.Person{})

	router := gin.Default()
	routes.SetupRoutes(router)
	controllers.InitPeopleController(db)

	router.Run(":8080")
}
