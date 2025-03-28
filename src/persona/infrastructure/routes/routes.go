package routes

import (
	"recu/src/persona/infrastructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/person")
	getAddPersonController := dependencies.GetAddPersonController().Execute
	getNewPersinIsAddedController := dependencies.GetNewPersonIsAddedController().Execute

	routes.POST("/addPerson", getAddPersonController)
	routes.GET("/newPersonIsAdded",getNewPersinIsAddedController)
	routes.GET("/CountGender")

}