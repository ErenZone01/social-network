package session

import (
	"fmt"
	"log"
	"main/pkg/db/sqlite"
	Struct "main/pkg/struct"
	Errors "main/pkg/untils/error"
	"net/http"
	"sync"
	"time"

	"github.com/gofrs/uuid/v5"
)

var (
	sessionMutex sync.Mutex
	Sessions     = make(map[int]string)
	SessionUser  = make(map[string]Struct.User)
)

func actualiseSession() {
	var Allsession, err = sqlite.DataSession()
	if err != nil {
		fmt.Println("session have problems : ", err)
		return
	}
	for i := 0; i < len(Allsession); i++ {
		SessionUser[Allsession[i].Value] = sqlite.GetUserById(Allsession[i].Users_id)
		Sessions[sqlite.GetUserById(Allsession[i].Users_id).Id] = Allsession[i].Value
	}
}
func Createsession(w http.ResponseWriter, users Struct.User) bool {
	sessionMutex.Lock()
	userID_s := sqlite.GetUser(users.Email)
	userID := userID_s.Id
	actualiseSession()
	_, exists := Sessions[userID]
	if exists {
		sessionMutex.Unlock()
		return false
	}
	users.Actif = "true"
	//Créez une nouvelle session
	sessionID := generateSessionID()
	SessionUser[sessionID] = users
	Sessions[userID] = sessionID
	//enregistrer la session dans la base de donnée
	var newSessionBD Struct.Session
	newSessionBD.Users_id = userID
	newSessionBD.Value = sessionID
	sqlite.NewSession(newSessionBD)
	sessionMutex.Unlock()
	expiration := time.Now().Add(2 * time.Hour)
	//All_Forum.AllSessions = SessionUser
	//Stockez l'identifiant de session dans un cookie
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   sessionID,
		Expires: expiration,
	})
	return true
}
func generateSessionID() string {
	//Create a Version 4 UUID.
	u2, err := uuid.NewV4()
	if err != nil {
		log.Fatalf("failed to generate UUID: %v", err)
	}
	//log.Printf("generated Version 4 UUID %v", u2)
	u3, err := uuid.FromString(u2.String())
	if err != nil {
		log.Fatalf("failed to parse UUID %q: %v", u2.String(), err)
	}
	//log.Printf("successfully parsed UUID %v", u3)
	return u3.String()
}
func DeleteCookies(w http.ResponseWriter, r *http.Request) {
	//nom du cookie
	var userOnline = Myaccount(w, r)
	userOnline.Actif = "false"
	sqlite.UpdateUser(userOnline)
	cookieName := "session"
	//Créez un cookie avec une date deja expiré
	expiration := time.Now().AddDate(0, 0, -1)
	deletecookie := http.Cookie{
		Name:    cookieName,
		Value:   "",
		Expires: expiration,
		Path:    "/",
	}
	http.SetCookie(w, &deletecookie)
	fmt.Println("Cookie supprimé avec succés!")
}
func Myaccount(w http.ResponseWriter, r *http.Request) Struct.User {
	// Récupérer le cookie de session
	session, _ := r.Cookie("session")

	// Vérifier si le cookie de session est vide
	if session == nil || session.Value == "" {
		Errors.SendError(w, r, "Le cookie de session est vide")
		user := Struct.User{}
		user.Error = true
		return user
	}

	// Vérifier si le cookie de session est valide
	actualiseSession()
	_, exists := SessionUser[session.Value]
	if !exists {
		fmt.Println("Le cookie de session n'est pas valide")
		Errors.SendError(w, r, "Le cookie de session n'est pas valide")

		user := Struct.User{}
		user.Error = true
		return user
	}

	// Le cookie de session est valide, récupérer les informations de l'utilisateur
	var user = SessionUser[session.Value]
	user.Error = false
	user.Actif = "true"

	// Renouveler la durée de vie du cookie de session
	expiration := time.Now().Add(2 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   session.Value,
		Expires: expiration,
	})

	return user
}
