package notification

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	responses "main/pkg/untils/Responses"
	"net/http"
)

var Invitation = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var myaccount = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if myaccount.Error {
		return
	}
	var notif Struct.Notif
	var err = json.NewDecoder(r.Body).Decode(&notif)
	if err != nil {
		fmt.Println("Follow : erreur de decodage json", err)
		return
	}
	sqlite.UpdateNotif(notif)
	var Sender = sqlite.GetUser(sqlite.GetNotifById(notif.ID_Notif).Sender)
	var Receiver = sqlite.GetUser(sqlite.GetNotifById(notif.ID_Notif).Receiver)
	var follow = sqlite.GetFollowByUsers(Sender.Id, Receiver.Id)
	fmt.Println("le follow : ", follow)

	if notif.States == "Decline" {
		sqlite.DeleteFollow(follow.ID_Follow)
	} else {
		follow.Operation = "true"
		sqlite.UpdateFollow(follow)
	}
	responses.SendResponsesHome(w, r, "responses succesfully", nil)
})
