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
		fmt.Println("Follow : erreur de decodage json", err)
		return
	}
	Group.IdCreator = user.Id
	Group.IdMember=  ""
	register.SaveImage(Group.GroupImage, Group.GroupImageFile, "group")

	// Appeler la fonction CreateGroup pour enregistrer le groupe dans la base de données
	sqlite.CreateNewGroup(Group)
	var Data Struct.AllData
	var Allgroup = sqlite.GetAllGroup()
	Data.Allgroup = Allgroup;
	fmt.Println("New group added by : ", user.Nickname)

	responses.SendResponsesHome(w, r, "response succesfully", Data)
})
