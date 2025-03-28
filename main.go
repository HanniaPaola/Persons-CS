package main

import (
    "database/sql"
    "log"
    "net/http"

    "API/src/internal/routes"   
    "API/src/internal/config"   
    "API/src/internal/personas" 

    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/handlers" 
)

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

    // Configurar CORS
    headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
    originsOk := handlers.AllowedOrigins([]string{"*"})
    methodsOk := handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS"})

    log.Println("Servidor en ejecución en :8080")
    // Usa el middleware CORS
    if err := http.ListenAndServe(":8080", handlers.CORS(originsOk, headersOk, methodsOk)(r)); err != nil {
        log.Fatalf("Error al iniciar el servidor: %v", err)
    }
}
