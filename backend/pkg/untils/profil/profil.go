package profil

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	responses "main/pkg/untils/Responses"
	"net/http"
)

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
	fmt.Println("le profil user est : ", ProfilUser)
	// var Allpost = sqlite.GetAllPost()
	var Allfollower = sqlite.GetMyFollowers(ProfilUser)
	var Allfollowing = sqlite.GetMyFollowing(ProfilUser)
	var followers []Struct.User
	var followings []Struct.User
	for _, v := range Allfollower {
		followers = append(followers, sqlite.GetUserById(v.ID_User))
	}
	for _, v := range Allfollowing {
		followings = append(followings, sqlite.GetUserById(v.ID_Receiver))
	}
	// data.Alluser = Alluser
	// data.Allpost = Allpost
	data.Allfollowers = followers
	data.Allfollowing = followings
	data.Myaccount = ProfilUser

	responses.SendResponsesHome(w, r, "response succesfully", data)
})
