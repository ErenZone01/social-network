package main

import (
	"database/sql"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	"main/pkg/untils/comment"
	"main/pkg/untils/decon"
	Errors "main/pkg/untils/error"
	"main/pkg/untils/event"
	"main/pkg/untils/follow"
	"main/pkg/untils/group"
	"main/pkg/untils/home"
	"main/pkg/untils/login"
	"main/pkg/untils/notification"
	"main/pkg/untils/post"
	"main/pkg/untils/profil"
	"main/pkg/untils/register"
	"net/http"
)

var DB *sql.DB

func MiddlewareCors(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Autoriser l'origine spécifique de votre application frontend
        w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
        // Autoriser les méthodes spécifiées, y compris OPTIONS
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        // Autoriser les en-têtes spécifiés, y compris Authorization
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        // Autoriser l'envoi de cookies
        w.Header().Set("Access-Control-Allow-Credentials", "true")
        // Autoriser les codes de statut personnalisés
        w.Header().Set("Access-Control-Expose-Headers", "Status-Code")
        
        // Répondre à la pré-vérification (preflight) OPTIONS avec un OK
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        // Poursuivre le traitement de la demande
        next.ServeHTTP(w, r)
    })
}


func MiddlewareMethodGet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			next.ServeHTTP(w, r)
		} else {
			Errors.SendError(w, r, http.StatusMethodNotAllowed, "Methods Not Allowed")
		}
	})
}
func MiddlewareMethodPost(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			next.ServeHTTP(w, r)
		} else {
			Errors.SendError(w, r, http.StatusMethodNotAllowed, "Methods Not Allowed")
		}
	})
}

func handlerFunction() {
	http.Handle("/", MiddlewareCors(MiddlewareMethodGet(home.Home)))
	http.Handle("/Register", MiddlewareCors(register.MiddlewareRegister(register.RegisterHandler)))
	http.Handle("/Login", MiddlewareCors(login.MiddlewareLogin(login.LoginHandler)))
	http.Handle("/Post", MiddlewareCors(MiddlewareMethodPost(post.Post)))
	http.Handle("/Group", MiddlewareCors(MiddlewareMethodPost(group.Group)))
	http.Handle("/InvitationGroup", MiddlewareCors(MiddlewareMethodPost(group.InvitationGroup)))
	http.Handle("/AddMemberGroup", MiddlewareCors(MiddlewareMethodPost(notification.AddMemberGroup)))
	http.Handle("/Comment", MiddlewareCors(MiddlewareMethodPost(comment.Comment)))
	http.Handle("/Follow", MiddlewareCors(MiddlewareMethodPost(follow.FollowHandler)))
	http.Handle("/UnFollow", MiddlewareCors(MiddlewareMethodPost(follow.UnFollowHandler)))
	http.Handle("/Event", MiddlewareCors(event.Event))
	http.Handle("/Decon", MiddlewareCors(MiddlewareMethodPost(decon.Decon)))
	http.Handle("/Invitation", MiddlewareCors(MiddlewareMethodPost(notification.Invitation)))
	http.Handle("/Profil", MiddlewareCors(MiddlewareMethodPost(profil.Profil)))
	http.Handle("/ProfilGroup", MiddlewareCors(MiddlewareMethodPost(profil.ProfilGroup)))
	http.Handle("/Privacy", MiddlewareCors(MiddlewareMethodPost(profil.Privacy)))
	http.Handle("/CheckSession", MiddlewareCors(MiddlewareMethodGet(session.CheckSession)))
}

func main() {
	DB = sqlite.CheckDB()
	handlerFunction()
	fmt.Println(":8080")
	http.ListenAndServe(":8080", nil)
}
