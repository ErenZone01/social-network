package login

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	Errors "main/pkg/untils/error"
	"net/http"
)

func MiddlewareLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err, _ = Login(w, r)
		if err {
			next.ServeHTTP(w, r)
			return
		}
		// Si l'utilisateur n'est pas authentifié, renvoyer une réponse d'erreur 401 Unauthorized
		fmt.Println("l'utilisateur 'existe pas")
		Errors.SendError(w, r, "Login or Password is incorrect")
		//http.Error(w, "l'utilisateur 'existe pas", http.StatusUnauthorized)
	})
}

func Login(w http.ResponseWriter, r *http.Request) (bool, Struct.User) {
	var newUser Struct.User
	decoder := json.NewDecoder(r.Body)
	errs := decoder.Decode(&newUser)
	if errs != nil {
		fmt.Println("Erreur de décodage JSON")
		//http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
		return false, Struct.User{}
	}
	fmt.Println("user : ", newUser)
	Alluser := sqlite.GetAllUser()
	var err, user = IfUserExist(Alluser, newUser)
	if !err {
		return err, user
	}
	session.Createsession(w, user)
	return err, user
}

func IfUserExist(Alluser []Struct.User, login Struct.User) (bool, Struct.User) {
	for _, v := range Alluser {
		if v.Email == login.Email && v.Password == login.Password {
			return true, v
		} else if v.Nickname == login.Email && v.Password == login.Password {
			return true, v
		}
	}
	return false, Struct.User{}
}

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("le marshal start")
	// Convertissez les données d'utilisateurs en JSON
	var msgFetch Struct.FetchMsg
	msgFetch.Types= "Reponses" 
	msgFetch.Msg= "Vous êtes connecté avec succé"
	msgFetch.Data=""
	fmt.Println("tout baigne")
	w.Header().Set("Content-Type", "application/json") // Définir le type de contenu de la réponse comme JSON
	json.NewEncoder(w).Encode(msgFetch)
})
