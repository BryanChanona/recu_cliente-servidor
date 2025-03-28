package routes

import (
	"recu/src/persona/infrastructure/dependencies"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	routes := router.Group("/person")
	getAddPersonController := dependencies.GetAddPersonController().Execute

	routes.POST("/addPerson", getAddPersonController)
	routes.GET("/newPersonIsAdded")
	routes.GET("/CountGender")

}