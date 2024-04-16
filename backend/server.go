package main

import (
	"database/sql"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/untils/comment"
	"main/pkg/untils/decon"
	"main/pkg/untils/event"
	"main/pkg/untils/follow"
	"main/pkg/untils/home"
	"main/pkg/untils/login"
	"main/pkg/untils/notification"
	"main/pkg/untils/post"
	"main/pkg/untils/register"
	"net/http"
)

var DB *sql.DB

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
	http.Handle("/Post", MiddlewareCors(post.Post))
	http.Handle("/Comment", MiddlewareCors(comment.Comment))
	http.Handle("/Follow", MiddlewareCors(follow.FollowHandler))
	http.Handle("/Event", MiddlewareCors(event.Event))
	http.Handle("/Decon", MiddlewareCors(decon.Decon))
	http.Handle("/Invitation", MiddlewareCors(notification.Invitation))
}

func main() {
	DB = sqlite.CheckDB()
	handlerFunction()
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}
