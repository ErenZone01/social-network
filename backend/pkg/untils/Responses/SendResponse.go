package responses

import (
	"encoding/json"
	Struct "main/pkg/struct"
	"net/http"
)

func SendResponsesHome(w http.ResponseWriter, r *http.Request, data interface{}) {
	var msgFetch Struct.FetchMsg
	msgFetch.Types = "Success"
	msgFetch.Msg = "connection succesfully"
	msgFetch.Data = data
	w.Header().Set("Content-Type", "application/json") // Définir le type de contenu de la réponse comme JSON
	json.NewEncoder(w).Encode(msgFetch)
}
