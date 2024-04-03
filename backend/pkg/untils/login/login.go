package login

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	Struct "main/pkg/struct"
	"net/http"
)

func MiddlewareLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err, user = Login(w, r)
		if err {
			fmt.Println("le marshal start")
			 // Convertissez les données d'utilisateurs en JSON
			 usersJSON, err := json.Marshal(user)
			 if err != nil {
				 http.Error(w, "Erreur de conversion en JSON", http.StatusInternalServerError)
				 return
			 }
			 fmt.Println("le user est : ", user)
			 // Envoyez la réponse JSON
			 w.Header().Set("Content-Type", "application/json")
			 w.Write(usersJSON)

			next.ServeHTTP(w, r)
			return
		}
		// Si l'utilisateur n'est pas authentifié, renvoyer une réponse d'erreur 401 Unauthorized
		fmt.Println("l'utilisateur 'existe pas")
		http.Error(w, "l'utilisateur 'existe pas", http.StatusUnauthorized)
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
	Alluser := sqlite.GetAllUser()
	var err, user = IfUserExist(Alluser, newUser)
	if !err {
		return err, user
	}
	return err, user
}

func IfUserExist(Alluser []Struct.User, login Struct.User) (bool, Struct.User) {
	for _, v := range Alluser {
		if v.Email == login.Email && v.Password ==  login.Password{
			return true, v
		} else if v.Nickname == login.Email && v.Password ==  login.Password{
			return true, v
		}
	}
	return false, Struct.User{}
}
