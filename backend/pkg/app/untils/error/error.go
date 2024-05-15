package Errors

import (
	"encoding/json"
	Struct "main/pkg/app/struct"
	"net/http"
)

func SendError(w http.ResponseWriter, r *http.Request, statusCode int, msg string) {
	// Convertissez les données d'utilisateurs en JSON
	var msgFetch Struct.FetchMsg
	msgFetch.Types = "Error"
	msgFetch.Msg = msg
	msgFetch.Data = ""

	// Définir le code d'erreur HTTP
	w.WriteHeader(statusCode)

	// Envoyez la réponse JSON
	w.Header().Set("Content-Type", "application/json") // Définir le type de contenu de la réponse comme JSON
	json.NewEncoder(w).Encode(msgFetch)
}
