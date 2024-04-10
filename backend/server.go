package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	Struct "main/pkg/struct"
	"main/pkg/untils/event"
	"main/pkg/untils/follow"
	"main/pkg/untils/home"
	"main/pkg/untils/login"
	"main/pkg/untils/register"
	"net/http"
)

func MiddlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Autoriser l'origine spécifique de votre application frontend
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		// Autoriser les méthodes spécifiées
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
		// Autoriser les en-têtes spécifiés, y compris Authorization
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		// Autoriser l'envoi de cookies
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		// Poursuivre le traitement de la demande
		next.ServeHTTP(w, r)
	})
}

func handlerFunction() {
	http.Handle("/", MiddlewareCors(home.Home))
	http.Handle("/Register", MiddlewareCors(register.MiddlewareRegister(register.RegisterHandler)))
	http.Handle("/Login", MiddlewareCors(login.MiddlewareLogin(login.LoginHandler)))
	http.HandleFunc("/Post", handlerPost)
	http.HandleFunc("/Comment", Comment)
	http.Handle("/Follow", MiddlewareCors(follow.FollowHandler))
	http.HandleFunc("/Event", MiddlewareEvent)
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func SendResponses(w http.ResponseWriter, r *http.Request, data interface{}) {
	var msgFetch Struct.FetchMsg
	msgFetch.Types = "Success"
	msgFetch.Msg = "connection succesfully"
	msgFetch.Data = data
	w.Header().Set("Content-Type", "application/json") // Définir le type de contenu de la réponse comme JSON
	json.NewEncoder(w).Encode(data)
}

func MiddlewareEvent(w http.ResponseWriter, r *http.Request) {
	event.Event(w, r)
}

func Comment(w http.ResponseWriter, r *http.Request) {

}

func handlerPost(w http.ResponseWriter, r *http.Request) {

}

var DB *sql.DB

func main() {
	DB = sqlite.CheckDB()
	handlerFunction()
}
