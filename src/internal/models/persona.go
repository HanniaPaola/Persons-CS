package models

type Persona struct {
    ID     int    `json:"id"`
    Edad   int    `json:"edad"`
    Nombre string `json:"nombre"`
    Sexo   bool   `json:"sexo"` 
}
