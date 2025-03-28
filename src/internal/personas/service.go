package personas

import "API/src/internal/models"

type Service struct {
    repo *Repository
}

func NewService(repo *Repository) *Service {
    return &Service{repo: repo}
}

func (s *Service) AddPersona(p models.Persona) error {
    return s.repo.AddPersona(p)
}

func (s *Service) CountGender() (int, int, error) {
    return s.repo.CountGender()
}

func (s *Service) GetLatestID() (int, error) {
    return s.repo.GetLatestID()
}
