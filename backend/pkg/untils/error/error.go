package Errors

import (
	"encoding/json"
	"fmt"
	Struct "main/pkg/struct"
	"net/http"
)

func SendError(w http.ResponseWriter, r *http.Request, msg string) {
	fmt.Println("l'erreur est : ", msg)
	// Convertissez les données d'utilisateurs en JSON
	var msgFetch Struct.FetchMsg
	msgFetch.Types = "Error"
	msgFetch.Msg = msg
	msgFetch.Data = ""

	// Envoyez la réponse JSON
	w.Header().Set("Content-Type", "application/json") // Définir le type de contenu de la réponse comme JSON
	json.NewEncoder(w).Encode(msgFetch)
}
