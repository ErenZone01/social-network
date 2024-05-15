package group

import (
	"encoding/json"
	"fmt"
	"main/pkg/app/session"
	Struct "main/pkg/app/struct"
	responses "main/pkg/app/untils/Responses"
	Errors "main/pkg/app/untils/error"
	"main/pkg/app/untils/register"
	"main/pkg/db/sqlite"
	"net/http"
)

var Group = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if user.Error {
		return
	}
	var Group Struct.Group

	var err = json.NewDecoder(r.Body).Decode(&Group)
	if err != nil {
		fmt.Println("group : erreur de decodage json", err)
		return
	}
	Group.IdCreator = user.Id
	if Group.GroupImage != "" {
		register.SaveImage(Group.GroupImage, Group.GroupImageFile, "group")
	}

	// Appeler la fonction CreateGroup pour enregistrer le groupe dans la base de données
	sqlite.CreateNewGroup(Group)
	var Data Struct.AllData
	var Allgroup = sqlite.GetAllGroup()
	Allgroup = GetMembersGroup(Allgroup)
	Data.UncknowGroup, Data.MyGroup = GroupCouldBeFollow(user, Allgroup)
	fmt.Println("New group added by : ", user.Nickname)

	responses.SendResponsesHome(w, r, "Your group has been created", Data)
})

func GetMembersGroup(Allgroup []Struct.Group) []Struct.Group {
	for i, v := range Allgroup {
		var tmp = sqlite.GetAllMemberGroup(v.ID_Group, 0)
		if len(tmp) != 0 {
			for _, m := range tmp {
				Allgroup[i].IdMember = append(Allgroup[i].IdMember, m.ID_User)
			}
		}
	}
	return Allgroup
}

func GetMembersProfileGroup(ProfileGroup Struct.Group) Struct.Group {
	var tmp = sqlite.GetAllMemberGroup(ProfileGroup.ID_Group, 0)
	if len(tmp) != 0 {
		for _, m := range tmp {
			ProfileGroup.IdMember = append(ProfileGroup.IdMember, m.ID_User)
		}
	}
	return ProfileGroup
}

var InvitationGroup = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	var ProfilGroup, _ = sqlite.GetGroupsByGroupID(Group.ID_Group)
	ProfilGroup = GetMembersProfileGroup(ProfilGroup)
	var receive Struct.User
	if Group.ID_User != 0 {
		receive = sqlite.GetUserById(Group.ID_User)
	} else {
		receive = sqlite.GetUserById(ProfilGroup.IdCreator)
	}
	if IsMember(receive, ProfilGroup) || IsNotifAlreadyStored(user.Nickname, receive.Nickname, Group.ID_Group) {
		var Data Struct.AllData
		var myNotif = sqlite.GetMyNotif(user.Nickname)
		Data.Allnotif = myNotif
		Errors.SendError(w, r, 200, "you have already sent your invitation !")
		return
	}
	Group.ID_User = user.Id
	Group.ID_Post = 0
	// enregistrer le nouveau membre dans la base de données
	sqlite.SetNotif(" Want to join : "+ProfilGroup.GroupName, receive.Nickname, user.Nickname, "pending", "group", user.Avatar, Group.ID_Group, 0)
	var Data Struct.AllData
	Data.Allnotif = sqlite.GetMyNotif(user.Nickname)
	Data.UncknowGroup, Data.MyGroup = GroupCouldBeFollow(user, GetMembersGroup(sqlite.GetAllGroup()))
	responses.SendResponsesHome(w, r, "Your request has been sent", Data)
})

func GroupCouldBeFollow(myaccount Struct.User, Allgroup []Struct.Group) ([]Struct.Group, []Struct.Group) {
	var UncknowGroup []Struct.Group
	var AlreadyFollowedGroup []Struct.Group
	var actif = true
	for _, v := range Allgroup {
		if v.IdCreator == myaccount.Id {
			AlreadyFollowedGroup = append(AlreadyFollowedGroup, v)
			continue
		}
		for _, v2 := range v.IdMember {
			if myaccount.Id == v2 {
				actif = false
				break
			}
		}
		if actif {
			var notifGroup = sqlite.GetAllPersonnalNotif(sqlite.GetUserById(v.IdCreator).Nickname, myaccount.Nickname, v.ID_Group)
			if (notifGroup != Struct.Notif{}) {
				v.States = notifGroup.States
			}
			UncknowGroup = append(UncknowGroup, v)
		} else {
			AlreadyFollowedGroup = append(AlreadyFollowedGroup, v)
			actif = true
		}
	}
	return UncknowGroup, AlreadyFollowedGroup
}

func IsMember(myaccount Struct.User, group Struct.Group) bool {
	for _, d := range group.IdMember {
		if d == myaccount.Id {
			return true
		}
	}
	return false
}

func IsNotifAlreadyStored(Sender string, Receive string, ID_Group int) bool {
	return sqlite.GetNotifPersonnal(Receive, Sender, ID_Group) != Struct.Notif{}
}
