package domain

type Person struct {
    Nombre string `json:"nombre"`
    Edad   int    `json:"edad"`
    Sexo   bool   `json:"sexo"` // true = masculino, false = femenino
}