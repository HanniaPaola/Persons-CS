package routes

import (
    "github.com/gorilla/mux"
    "API/src/internal/personas"
)

func NewRouter(handler *personas.Handler) *mux.Router {
    r := mux.NewRouter()
	r.HandleFunc("/addPerson", handler.AddPerson).Methods("POST")
	r.HandleFunc("/countGender", handler.CountGender).Methods("GET")
	r.HandleFunc("/newPersonIsAdded", handler.NewPersonIsAdded).Methods("GET")
    return r
}


