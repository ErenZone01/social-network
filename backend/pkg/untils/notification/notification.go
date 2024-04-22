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
	var Sender = sqlite.GetUser(sqlite.GetNotifById(notif.ID_Notif).Sender)
	var Receiver = sqlite.GetUser(sqlite.GetNotifById(notif.ID_Notif).Receiver)
	var follow = sqlite.GetFollowByUsers(Sender.Id, Receiver.Id)

	if notif.States == "Decline" {
		sqlite.DeleteFollow(follow.ID_Follow)
		sqlite.DeleteNotif(notif.ID_Notif)
	} else {
		follow.Operation = "true"
		sqlite.UpdateFollow(follow)
		notif.States = "true"
		notif.Messages = "started following you"
		sqlite.UpdateNotif(notif)
	}
	var data Struct.AllData;
	var follower []Struct.User
	var newNotif = sqlite.GetMyNotif(myaccount.Nickname)
	var newFollower =sqlite.GetMyFollowers(myaccount)
	for _, v := range newFollower {
		follower = append(follower, sqlite.GetUserById(v.ID_Receiver))
	}
	data.Allnotif =newNotif
	data.Allfollowers = follower
	responses.SendResponsesHome(w, r, "responses succesfully", data)
})
