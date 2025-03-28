package main

import (
	"recu/src/persona/infrastructure/dependencies"
	"recu/src/persona/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	dependencies.Init()

	r := gin.Default()
	routes.Routes(r)

	r.Run(":8081")

}