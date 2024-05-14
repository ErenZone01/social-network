package notification

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	responses "main/pkg/untils/Responses"
	"main/pkg/untils/group"
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
		fmt.Println("Notif : erreur de decodage json", err)
		return
	}
	var data Struct.AllData
	if notif.Types == "person" {
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
		var follower []Struct.User
		var newNotif = sqlite.GetMyNotif(myaccount.Nickname)
		var newFollower = sqlite.GetMyFollowers(myaccount)
		for _, v := range newFollower {
			follower = append(follower, sqlite.GetUserById(v.ID_Receiver))
		}
		data.Allnotif = newNotif
		data.Allfollowers = follower
	} else if notif.Types == "group" {
		if notif.States == "Decline" {
			sqlite.DeleteNotif(notif.ID_Notif)
		} else {
			fmt.Println("je suis dans 'invitation group")
			var Group Struct.Members
			var Notif = sqlite.GetNotifById(notif.ID_Notif)
			var ProfileGroup, _ = sqlite.GetGroupsByGroupID(Notif.ID_Group)
			ProfileGroup = group.GetMembersProfileGroup(ProfileGroup)
			if group.IsMember(myaccount, ProfileGroup) {
				fmt.Println("utilisateur deja dans le group")
				sqlite.DeleteNotif(notif.ID_Notif)
				data.Allnotif = sqlite.GetMyNotif(myaccount.Nickname)
				data.UncknowGroup, data.MyGroup = group.GroupCouldBeFollow(myaccount, group.GetMembersGroup(sqlite.GetAllGroup()))
				responses.SendResponsesHome(w, r, " you have already member of this group !", data)
				return
			}
			//si le cretaor est egale au receveur alors c'est une demande fait par un utilistauer hors du group
			if ProfileGroup.IdCreator == myaccount.Id {
				Group.ID_User = sqlite.GetUser(Notif.Sender).Id
			} else { //si le creator n'est pas celui qu'on a envoyé la demande alors c'est un utlisateur l'ambda le receveur donc un membre du groupe a fait la demande
				Group.ID_User = sqlite.GetUser(Notif.Receiver).Id
			}
			Group.ID_Post = 0
			Group.ID_Group = Notif.ID_Group
			sqlite.CreateNewMember(Group)
			//update notif to true
			Notif.States = "true"
			Notif.Messages = " integrate the group "
			sqlite.UpdateNotif(Notif)
		}
		data.Allnotif = sqlite.GetMyNotif(myaccount.Nickname)
		data.UncknowGroup, data.MyGroup = group.GroupCouldBeFollow(myaccount, group.GetMembersGroup(sqlite.GetAllGroup()))
		fmt.Println("New member added")
	} else if notif.Types == "event" {
		var Notif = sqlite.GetNotifById(notif.ID_Notif)
		var Event, _ = sqlite.GetEventByID(Notif.ID_Event)
		if notif.States == "Decline" {
			sqlite.DeleteNotif(notif.ID_Notif)
		} else {
			var members Struct.Members
			members.ID_Event = Event.ID_Event
			members.ID_Group = Event.ID_Group
			members.ID_Post = 0
			members.ID_User = myaccount.Id
			sqlite.CreateNewMember(members)
			//update notif to true
			Notif.States = "true"
			Notif.Messages = " participe in the event "
			sqlite.UpdateNotif(Notif)
		}
		data.Allnotif = sqlite.GetMyNotif(myaccount.Nickname)
		data.Allevent =GetMembersEvent(Event.ID_Group)
		fmt.Println("New member added to event")
	}
	responses.SendResponsesHome(w, r, "responses succesfully", data)
})

func GetMembersEvent(ID_Group int) []Struct.EventGroup {
	var Events []Struct.EventGroup
	for _, v := range sqlite.GetEventByUsers(ID_Group) {
		tmp := sqlite.GetAllMemberGroup(ID_Group, v.ID_Event)
		for _, d := range tmp {
			v.ID_Member = append(v.ID_Member, d.ID_User)
		}
		Events = append(Events, v)
	}
	return Events
}

var AddMemberGroup = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if user.Error {
		return
	}
	var Group Struct.Members

	var err = json.NewDecoder(r.Body).Decode(&Group)
	if err != nil {
		fmt.Println("group : erreur de decodage json", err)
		return
	}
	Group.ID_User = user.Id
	Group.ID_Post = 0
	// enregistrer le nouveau membre dans la base de données
	sqlite.CreateNewMember(Group)

	//sqlite.SetNotif(" Want to intregrate : ", "", user.Nickname, "pending", "group", user.Avatar, Group.ID_Group)
	var data Struct.AllData

	responses.SendResponsesHome(w, r, "response succesfully", data)
})
