package main

import (
	"database/sql"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/untils/event"
	"main/pkg/untils/follow"
	"main/pkg/untils/home"
	"main/pkg/untils/login"
	"main/pkg/untils/register"
	"net/http"
)

var RegisterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("Bienvenue sur la page d'accueil"))
	fmt.Println("fin")
})
var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("fin")
	// w.Write([]byte("Bienvenue sur la page d'accueil"))
})

func handlerFunction() {
	http.HandleFunc("/", handler)
	http.Handle("/Register", register.MiddlewareRegister(RegisterHandler))
	http.Handle("/Login", login.MiddlewareLogin(LoginHandler))
	http.HandleFunc("/Home", Login)
	http.HandleFunc("/Post", handlerPost)
	http.HandleFunc("/Comment", handler)
	http.HandleFunc("/Follow", handler)
	http.HandleFunc("/Event", handler)
	fmt.Println("http://localhost:8080/")
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	home.Home(w, r)
}
func Event(w http.ResponseWriter, r *http.Request) {
	event.Event(w, r)
}
func Follow(w http.ResponseWriter, r *http.Request) {
	follow.Follow(w, r)
}
func Comment(w http.ResponseWriter, r *http.Request) {

}

func Login(w http.ResponseWriter, r *http.Request) {

}

func handlerPost(w http.ResponseWriter, r *http.Request) {

}

var DB *sql.DB

func main() {
	DB = sqlite.CreateBD()
	handlerFunction()
}
