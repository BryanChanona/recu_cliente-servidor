package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"recu/src/persona/domain"
)

type MySQL struct {
	db *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{db: db}
}



func (sql *MySQL) SavePerson(person domain.Person) error{

	query, err:= sql.db.Prepare("INSERT INTO personas (nombre,edad,sexo) VALUES (?,?,?)")
	if err != nil {
		return fmt.Errorf("error preparando consulta SQL: %v", err)
	}
	defer query.Close() // Aquí sí se usa defer correctamente

	_, err = query.Exec(&person.Nombre,&person.Edad,&person.Sexo)
	if err != nil {
		return fmt.Errorf("error ejecutando la consulta SQL: %v", err)
	}
	log.Println("Oxígeno guardado correctamente en la base de datos.")
	return nil

}