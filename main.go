package main

import (
    "database/sql"
    "log"
    "net/http"

    "API/src/archivos/routes"
    "API/src/config"
    "API/src/archivos/personas"

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/handlers"
)

func optionsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "http://127.0.0.1:5500")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
    w.WriteHeader(http.StatusOK)
}

func main() {
    config.LoadEnv()

    db, err := sql.Open("mysql", "root:hannia@tcp(localhost:3306)/persona")
    if err != nil {
        log.Fatalf("Error al conectar a la base de datos: %v", err)
    }
    defer db.Close()

    repo := personas.NewRepository(db)
    service := personas.NewService(repo)
    handler := personas.NewHandler(service)

    r := routes.NewRouter(handler)

    // Agregar un handler para las solicitudes OPTIONS en todas las rutas
    r.Methods("OPTIONS").HandlerFunc(optionsHandler)

    // Configurar CORS
    headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    originsOk := handlers.AllowedOrigins([]string{"http://127.0.0.1:5500"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

    log.Println("Servidor en ejecución en :8085")

    if err := http.ListenAndServe(":8085", handlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}
