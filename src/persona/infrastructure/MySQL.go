package infrastructure

import (
	"database/sql"
	"fmt"
	"log"
	"recu/src/persona/domain"
)

type MySQL struct {
	db *sql.DB
	lastCount int
}

func NewMySQL(db *sql.DB) *MySQL {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM personas").Scan(&count)
	if err != nil {
		log.Fatal("Error obteniendo el conteo inicial: ", err)
	}

	// Inicializamos el repositorio con el conteo inicial
	return &MySQL{
		db:        db,
		lastCount: count, // Establecemos el conteo inicial
	}
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

func (sql *MySQL)GetnewPersonIsAdded() (bool, error){

	var count int
	err := sql.db.QueryRow("SELECT COUNT(*) FROM personas").Scan(&count)
	if err != nil {
		return false, fmt.Errorf("error obteniendo el conteo de personas: %v", err)
	}

	if count > sql.lastCount {
		// Si el conteo actual es mayor que el anterior, significa que se ha agregado una persona
		// Actualizamos lastCount para mantener el conteo más reciente
		sql.lastCount = count
		return true, nil
	}

	return false, nil
}