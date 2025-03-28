package dependencies

import (
	"log"
	"recu/src/helpers"
	"recu/src/persona/application"
	"recu/src/persona/infrastructure"
	"recu/src/persona/infrastructure/controllers"
)

var (
	mySQL infrastructure.MySQL
)

func Init() {
	db, err := helpers.ConnMySQL()

	if err != nil {
		log.Fatalf("Error al conectar a la base de datos: %v", err)
	}
	mySQL = *infrastructure.NewMySQL(db)

}


func GetAddPersonController()*controllers.AddPersonController{
	useCase := application.NewAddPersonUc(&mySQL)
	return controllers.NewAddPersonController(useCase)
}