package controllers

import (
	"net/http"
	"recu/src/persona/application"
	"time"

	"github.com/gin-gonic/gin"
)

type GetNewPersonIsAddedController struct {
	useCase *application.GetNewPersonIsAddedUc
}

func NewGetNewPersonIsAddedController(useCase *application.GetNewPersonIsAddedUc)*GetNewPersonIsAddedController{
	return &GetNewPersonIsAddedController{useCase: useCase}

}

func (controller *GetNewPersonIsAddedController)Execute(ctx *gin.Context){
	pollingInterval := 5 * time.Second
	timeout := time.After(30 * time.Second) // Tiempo de espera máximo para evitar la espera infinita

	ticker := time.NewTicker(pollingInterval)
	defer ticker.Stop()

	// Bucle de short polling
	for {
		select {
		case <-timeout:
			// Si se excede el tiempo máximo, respondemos con un error
			ctx.JSON(http.StatusRequestTimeout, gin.H{"error": "Polling timed out"})
			return
		case <-ticker.C:
			// Verifica si hay una nueva persona añadida
			newPersonAdded, err := controller.useCase.Execute()
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			// Si se encuentra una nueva persona añadida, respondemos al cliente
			if newPersonAdded {
				ctx.JSON(http.StatusOK, gin.H{"new_person_added": true})
				return
			}
		}
	}
}