package routes

import (
	"effective_mobile/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(router *gin.Engine) {
	router.POST("/people", controllers.AddPerson)
	router.GET("/people", controllers.GetPeople)
	router.PUT("/people/:id", controllers.UpdatePerson)
	router.DELETE("/people/:id", controllers.DeletePerson)
}
