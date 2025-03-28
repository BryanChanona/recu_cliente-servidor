package controllers

import (
	"net/http"
	"recu/src/persona/application"
	"time"

	"github.com/gin-gonic/gin"
)

type CountGenderController struct {
	useCase *application.CountGenderUc
}

func NewCountGenderUc(useCase *application.CountGenderUc)*CountGenderController{
 return &CountGenderController{useCase: useCase}
}

func (controller *CountGenderController) Execute(c *gin.Context){
	// Obtener el parámetro de la consulta (sexo) de la solicitud
	// Obtener el parámetro de la consulta (sexo) de la solicitud
	sexo := c.Param("sexo") // "true" es el valor por defecto si no se pasa nada

	var sexoBool bool
	if sexo == "true" {
		sexoBool = true
	} else if sexo == "false" {
		sexoBool = false
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Valor de sexo inválido. Debe ser 'true' o 'false'."})
		return
	}

	// Tiempo máximo de espera para long polling (por ejemplo, 30 segundos)
	timeout := time.After(30 * time.Second)

	// Intervalo de revisión para verificar si el conteo ha cambiado
	checkInterval := time.NewTicker(5 * time.Second)
	defer checkInterval.Stop()

	// Bucle de long polling
	for {
		select {
		case <-timeout:
			// Si se excede el tiempo máximo, devolvemos un timeout
			c.JSON(http.StatusRequestTimeout, gin.H{"error": "Long polling timed out"})
			return
		case <-checkInterval.C:
			// Llamar al caso de uso para obtener el conteo basado en el sexo
			count, err := controller.useCase.Execute(sexoBool)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			// Si el conteo cambia, respondemos al cliente
			if count > 0 { // Asumiendo que el conteo mayor a 0 indica que hay datos
				c.JSON(http.StatusOK, gin.H{"conteo": count})
				return
			}
		}
	}
}