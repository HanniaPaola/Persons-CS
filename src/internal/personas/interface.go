package personas

import "API/src/internal/models"

type PersonaService interface {
    AddPersona(p models.Persona) error
    CountGender() (int, int, error)
}
