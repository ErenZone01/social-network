package event

import (
	"encoding/json"
	"fmt"
	"main/pkg/app/session"
	Struct "main/pkg/app/struct"
	responses "main/pkg/app/untils/Responses"
	Errors "main/pkg/app/untils/error"
	"main/pkg/app/untils/notification"
	"main/pkg/db/sqlite"
	"net/http"
)

var Event = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var myaccount = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if myaccount.Error {
		return
	}
	var event Struct.EventGroup
	var err = json.NewDecoder(r.Body).Decode(&event)
	if err != nil {
		fmt.Println("Event : erreur de decodage json", err)
		return
	}

	var mygroup, err2 = sqlite.GetGroupsByGroupID(event.ID_Group)
	if err2 != nil {
		Errors.SendError(w, r, 400, "Group doesn't exist !")
		return
	}
	event.ID_User = myaccount.Id
	sqlite.CreateNewEvent(event)

	var newMember []int
	var tab = sqlite.GetAllMemberGroup(mygroup.ID_Group, 0)
	for _, v := range tab {
		newMember = append(newMember, v.ID_User)
	}
	fmt.Println("les members : ", newMember)
	var myEventGlob = sqlite.GetEventByUsers(mygroup.ID_Group)
	var myEvent = myEventGlob[len(myEventGlob)-1]
	mygroup.IdMember = newMember
	for _, v := range mygroup.IdMember {
		var user = sqlite.GetUserById(v)
		if user.Id != myaccount.Id {
			sqlite.SetNotif(myaccount.Nickname+" invites you to participate in the "+mygroup.GroupName+" group in the "+event.Title+"event ", user.Nickname, myaccount.Nickname, "pending", "event", myaccount.Avatar, mygroup.ID_Group, myEvent.ID_Event)
		}
	}
	if myaccount.Id != mygroup.IdCreator {
		var creator = sqlite.GetUserById(mygroup.IdCreator)
		sqlite.SetNotif(myaccount.Nickname+" invites you to participate in the "+mygroup.GroupName+" group in the "+event.Title+"event ", creator.Nickname, myaccount.Nickname, "pending", "event", myaccount.Avatar, mygroup.ID_Group, myEvent.ID_Event)
	}
	if event.Option == "Going" {
		var members Struct.Members
		members.ID_Event = myEvent.ID_Event
		members.ID_Group = mygroup.ID_Group
		members.ID_Post = 0
		members.ID_User = myaccount.Id
		sqlite.CreateNewMember(members)
	}
	var events = notification.GetMembersEvent(mygroup.ID_Group)
	responses.SendResponsesHome(w, r, "responses succesfully", events)
})
