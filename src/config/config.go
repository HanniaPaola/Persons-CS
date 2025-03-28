package config

import (
    "log"

    "github.com/joho/godotenv"
)

func LoadEnv() {
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error al cargar el archivo .env: %v", err)
    }
}
