package personas

import "API/src/archivos/models"

type PersonaService interface {
    AddPersona(p models.Persona) error
    CountGender() (int, int, error)
	GetLatestID() (int, error)
}
