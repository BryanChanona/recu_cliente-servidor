package controllers

import (
	"net/http"
	"recu/src/persona/application"
	"recu/src/persona/domain"

	"github.com/gin-gonic/gin"
)

type AddPersonController struct {
	useCase *application.AddPersonUc
}

func NewAddPersonController(useCase *application.AddPersonUc) *AddPersonController{
	return &AddPersonController{useCase: useCase}

}

func (controller *AddPersonController)Execute(ctx *gin.Context){
	var person domain.Person

	if err := ctx.ShouldBindJSON(&person); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := controller.useCase.Execute(person)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{"message": "Registered person"})
	}
}
