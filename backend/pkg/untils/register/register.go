package register

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	Struct "main/pkg/struct"
	Errors "main/pkg/untils/error"
	"net/http"
)

func MiddlewareRegister(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err, msg = Register(w, r)
		if err {
			next.ServeHTTP(w, r)
			return
		}
		fmt.Println("err : ", msg)
		Errors.SendError(w, r, msg)
		// Si l'utilisateur n'est pas authentifié, renvoyer une réponse d'erreur 401 Unauthorized
		//http.Error(w, msg, http.StatusUnauthorized)
	})
}

func Register(w http.ResponseWriter, r *http.Request) (bool, string) {
	var newUser Struct.User
	//recuperation des données du register
	decoder := json.NewDecoder(r.Body)
	errs := decoder.Decode(&newUser)
	if errs != nil {
		fmt.Println("Erreur de décodage JSON")
		//http.Error(w, "Erreur de décodage JSON", http.StatusBadRequest)
		return false, "Erreur de decodage JSON"
	}
	fmt.Println("le user est : ", newUser)
	Alluser := sqlite.GetAllUser()
	var err, msg = UserUnique(Alluser, newUser.Email, newUser.Nickname)
	if !err {
		fmt.Println("l'erreur est ici")
		return err, msg
	}
	//AJouter l'utilisateur a la BD
	sqlite.CreateNewUser(newUser)
	return err, msg
}

func UserUnique(Alluser []Struct.User, Email string, Nickname string) (bool, string) {
	for _, v := range Alluser {
		if v.Email == Email {
			return false, "Email is already used"
		}
		if v.Nickname == Nickname {
			return false, "Nickname is already used"
		}
	}
	return true, "It's valid"
}

var RegisterHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	fmt.Println("le marshal start")
	// Convertissez les données d'utilisateurs en JSON
	var msgFetch Struct.FetchMsg
	msgFetch.Types= "Reponses" 
	msgFetch.Msg= "Votre compte a été crée avec succé"
	msgFetch.Data=""
	
	// Envoyez la réponse JSON
	w.Header().Set("Content-Type", "application/json") // Définir le type de contenu de la réponse comme JSON
	json.NewEncoder(w).Encode(msgFetch)
})
