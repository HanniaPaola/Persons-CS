package personas

import (
    "database/sql"
    "API/src/archivos/models"
)

type Repository struct {
    db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
    return &Repository{db: db}
}

func (r *Repository) AddPersona(p models.Persona) error {
    _, err := r.db.Exec("INSERT INTO personas (edad, nombre, sexo) VALUES (?, ?, ?)", p.Edad, p.Nombre, p.Sexo)
    return err
}

func (r *Repository) CountGender() (int, int, error) {
    var countHombres, countMujeres int
    err := r.db.QueryRow("SELECT COUNT(*) FROM personas WHERE sexo = false").Scan(&countHombres)
    if err != nil {
        return 0, 0, err
    }
    err = r.db.QueryRow("SELECT COUNT(*) FROM personas WHERE sexo = true").Scan(&countMujeres)
    return countHombres, countMujeres, err
}

func (r *Repository) GetLatestID() (int, error) {
    var latestID int
    err := r.db.QueryRow("SELECT MAX(id) FROM personas").Scan(&latestID)
    if err != nil {
        return 0, err
    }
    return latestID, nil
}