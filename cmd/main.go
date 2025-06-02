package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"golang-mail-service/internal/handlers"
	"log"
	"net/http"
)

func main() {
	_ = godotenv.Load()

	http.HandleFunc("/send", handlers.SendEmailHandler)

	fmt.Println("Servidor escuchando en :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
