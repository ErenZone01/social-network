package login

import (
	"encoding/json"
	"fmt"
	"main/pkg/app/session"
	Struct "main/pkg/app/struct"
	responses "main/pkg/app/untils/Responses"
	Errors "main/pkg/app/untils/error"
	"main/pkg/db/sqlite"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func MiddlewareLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err, _ = Login(w, r)
		if err {
			next.ServeHTTP(w, r)
			return
		}
		// Si l'utilisateur n'est pas authentifié, renvoyer une réponse d'erreur 401 Unauthorized
		Errors.SendError(w, r, http.StatusOK, "Login or Password is incorrect")
		//http.Error(w, "l'utilisateur 'existe pas", http.StatusUnauthorized)
	})
}

func Login(w http.ResponseWriter, r *http.Request) (bool, Struct.User) {
	var newUser Struct.User
	decoder := json.NewDecoder(r.Body)
	errs := decoder.Decode(&newUser)
	if errs != nil {
		fmt.Println("Erreur de décodage JSON")
		return false, Struct.User{}
	}
	Alluser := sqlite.GetAllUser()
	var err, user = IfUserExist(Alluser, newUser)
	if !err {
		return err, user
	}
	var state = session.Createsession(w, user)
	if !state {
		SessionValue := session.Sessions[user.Id]
		//supprimer la session des maps locals
		delete(session.SessionUser, SessionValue)
		delete(session.Sessions, user.Id)
		//supprimer la session de la BD
		sqlite.DeleteSession(SessionValue)
		session.Createsession(w, user)
	}
	user.Actif = "true"
	sqlite.UpdateUser(user)
	fmt.Println(user.Nickname, " is connected !")
	return err, user
}

func IfUserExist(Alluser []Struct.User, login Struct.User) (bool, Struct.User) {
	for _, v := range Alluser {
		if v.Email == login.Email && bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(login.Password)) == nil {
			return true, v
		}
	}
	return false, Struct.User{}
}

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // Définir le type de contenu de la réponse comme JSON
	responses.SendResponsesHome(w, r, "connection succesfully", nil)
})
