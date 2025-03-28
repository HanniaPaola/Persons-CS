package personas

import (
    "encoding/json"
    "net/http"
    "strconv"              
    "API/src/archivos/models"    
)

type Handler struct {
    service PersonaService
}

func NewHandler(service PersonaService) *Handler {
    return &Handler{service: service}
}

func (h *Handler) AddPerson(w http.ResponseWriter, r *http.Request) {
    var persona models.Persona
    if err := json.NewDecoder(r.Body).Decode(&persona); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    if err := h.service.AddPersona(persona); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusCreated)
}

func (h *Handler) CountGender(w http.ResponseWriter, r *http.Request) {
    hombres, mujeres, err := h.service.CountGender()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    response := map[string]int{"hombres": hombres, "mujeres": mujeres}
    json.NewEncoder(w).Encode(response)
}

func (h *Handler) NewPersonIsAdded(w http.ResponseWriter, r *http.Request) {
    lastID := r.URL.Query().Get("lastID")

    if lastID == "" {
        http.Error(w, `{"error": "lastID is required"}`, http.StatusBadRequest)
        return
    }

    lastIDInt, err := strconv.Atoi(lastID)
    if err != nil {
        http.Error(w, `{"error": "Invalid lastID"}`, http.StatusBadRequest)
        return
    }

    latestID, err := h.service.GetLatestID()
    if err != nil {
        http.Error(w, `{"error": "Error fetching latest ID"}`, http.StatusInternalServerError)
        return
    }

    response := map[string]bool{"newData": latestID > lastIDInt}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(response)
}

