package handlers

import (
	"encoding/json"
	"golang-mail-service/internal/email"
	"net/http"
)

type EmailRequest struct {
	To      string `json:"to"`
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

func SendEmailHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	var req EmailRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Error al parsear JSON", http.StatusBadRequest)
		return
	}

	if err := email.SendMail(req.To, req.Subject, req.Body); err != nil {
		http.Error(w, "Error al enviar correo: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Correo enviado exitosamente"))
}
