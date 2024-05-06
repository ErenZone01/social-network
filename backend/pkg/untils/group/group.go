package group

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	responses "main/pkg/untils/Responses"
	"main/pkg/untils/register"
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
	register.SaveImage(Group.GroupImage, Group.GroupImageFile, "group")

	// Appeler la fonction CreateGroup pour enregistrer le groupe dans la base de données
	sqlite.CreateNewGroup(Group)
	var Data Struct.AllData
	var Allgroup = sqlite.GetAllGroup()
	Allgroup = GetMembersGroup(Allgroup)
	Data.UncknowGroup, Data.MyGroup = GroupCouldBeFollow(user, Allgroup)
	fmt.Println("New group added by : ", user.Nickname)

	responses.SendResponsesHome(w, r, "response succesfully", Data)
})

func GetMembersGroup(Allgroup []Struct.Group) []Struct.Group {
	for i, v := range Allgroup {
		var tmp = sqlite.GetAllMemberGroup(v.ID_Group)
		if len(tmp) != 0 {
			for _, m := range tmp {
				Allgroup[i].IdMember = append(Allgroup[i].IdMember, m.ID_User)
			}
		}
	}
	return Allgroup
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
	var receive Struct.User
	if Group.ID_User != 0 {
		receive = sqlite.GetUserById(Group.ID_User)
	} else {
		receive = sqlite.GetUserById(ProfilGroup.IdCreator)
	}
	Group.ID_User = user.Id
	Group.ID_Post = 0
	// enregistrer le nouveau membre dans la base de données
	sqlite.SetNotif(" Want to intregrate : ", receive.Nickname, user.Nickname, "pending", "group", user.Avatar, Group.ID_Group)
	var Data Struct.AllData
	var myNotif = sqlite.GetMyNotif(user.Nickname)
	Data.Allnotif = myNotif
	responses.SendResponsesHome(w, r, "response succesfully", Data)
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
			UncknowGroup = append(UncknowGroup, v)
		} else {
			AlreadyFollowedGroup = append(AlreadyFollowedGroup, v)
			actif = true
		}
	}
	return UncknowGroup, AlreadyFollowedGroup
}
