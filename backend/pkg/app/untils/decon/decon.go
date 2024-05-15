package decon

import (
	"fmt"
	"main/pkg/app/session"
	responses "main/pkg/app/untils/Responses"
	"main/pkg/db/sqlite"
	"net/http"
)

var Decon = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var Myaccount = session.Myaccount(w, r)
	if Myaccount.Error {
		return
	}
	Myaccount.Actif = "false"
	session.DeleteCookies(w, r)
	SessionValue := session.Sessions[Myaccount.Id]
	//supprimer la session des maps locals
	delete(session.SessionUser, SessionValue)
	delete(session.Sessions, Myaccount.Id)
	//supprimer la session de la BD
	sqlite.DeleteSession(SessionValue)
	fmt.Println(Myaccount.Nickname, " is deconnected successfuly")
	responses.SendResponsesHome(w, r, "Vous êtes déconnecté avec succé", nil)
})
