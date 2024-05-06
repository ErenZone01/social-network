package profil

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	responses "main/pkg/untils/Responses"
	Errors "main/pkg/untils/error"
	"main/pkg/untils/home"
	"net/http"
)

var Privacy = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if user.Error {
		return
	}
	var Myaccount Struct.User
	err := json.NewDecoder(r.Body).Decode(&Myaccount)
	if err != nil {
		fmt.Println("Profil incorrect", err)
		return
	}
	if Myaccount.Privacy == "public" {
		Myaccount.Privacy = "private"
	} else {
		Myaccount.Privacy = "public"
	}
	sqlite.UpdateUser(Myaccount)

	var data Struct.AllData
	var newAccount = sqlite.GetUserById(Myaccount.Id)
	var Alluser = sqlite.GetAllUser()
	var followings []Struct.User
	for _, v := range sqlite.GetMyFollowing(user) {
		followings = append(followings, sqlite.GetUserById(v.ID_Receiver))
	}
	data.Allfollowing = followings
	data.Myaccount = newAccount
	data.Alluser = home.UserCouldBeFollow(newAccount, Alluser, followings)
	responses.SendResponsesHome(w, r, "response succesfully", data)
})

var Profil = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if user.Error {
		return
	}
	var idUser Struct.User
	err := json.NewDecoder(r.Body).Decode(&idUser)
	if err != nil {
		fmt.Println("Profil incorrect", err)
		return
	}
	var data = Struct.AllData{}
	var ProfilUser = sqlite.GetUserById(idUser.Id)
	if ProfilUser.Email == "" {
		Errors.SendError(w, r, http.StatusMethodNotAllowed, "Methods Not Allowed")
		return
	}
	var Allfollower = sqlite.GetMyFollowers(ProfilUser)
	var Allfollowing = sqlite.GetMyFollowing(ProfilUser)
	var MyPost, _ = sqlite.GetPostsByUserID(ProfilUser.Id)
	var followers = []Struct.User{}
	var followings = []Struct.User{}

	for _, v := range Allfollower {
		followers = append(followers, sqlite.GetUserById(v.ID_User))
	}
	for _, v := range Allfollowing {
		followings = append(followings, sqlite.GetUserById(v.ID_Receiver))
	}
	data.Allfollowers = followers
	data.Allfollowing = followings
	data.Allpost = MyPost
	data.Myaccount = ProfilUser

	responses.SendResponsesHome(w, r, "response succesfully", data)
})

var ProfilGroup = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if user.Error {
		return
	}
	var idGroup Struct.Group
	err := json.NewDecoder(r.Body).Decode(&idGroup)
	if err != nil {
		fmt.Println("Profil incorrect", err)
		return
	}
	var ProfilGroup, err2 = sqlite.GetGroupsByGroupID(idGroup.ID_Group)
	if err2 != nil {
		Errors.SendError(w, r, http.StatusMethodNotAllowed, "Methods Not Allowed")
		return
	}
	var newMember []int
	var tab = sqlite.GetAllMemberGroup(ProfilGroup.ID_Group)
	for _, v := range tab {
		newMember = append(newMember, v.ID_User)
	}
	ProfilGroup.IdMember = newMember
	var Data Struct.AllData
	var MyPost = sqlite.GetAllPostByGroup(idGroup.ID_Group)
	
	Data.Allpost = MyPost
	Data.MyGroup = append(Data.MyGroup, ProfilGroup) 
	responses.SendResponsesHome(w, r, "response succesfully", Data)
})
