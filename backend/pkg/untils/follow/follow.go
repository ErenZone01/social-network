package follow

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

var FollowHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var myaccount = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if myaccount.Error {
		return
	}
	var following Struct.User
	var err = json.NewDecoder(r.Body).Decode(&following)
	if err != nil {
		fmt.Println("Follow : erreur de decodage json", err)
		return
	}
	var operation, msg string
	if following.Privacy == "public" {
		operation = "true"
		msg = "started following you"
	} else {
		operation = "pending"
		msg = "want to follow you"
	}
	//tempFollowing := sqlite.GetMyFollowing(myaccount)
	if IsAlreadyFollowing(myaccount, following) {
		fmt.Println("You have already follow : ", following.Nickname)
		Errors.SendError(w, r, "You have already follow : "+following.Nickname)
		return
	}
	sqlite.SetFollowing(myaccount, following, operation, "person", 0)
	sqlite.SetNotif(msg, following.Nickname, myaccount.Nickname, operation, "person", myaccount.Avatar, 0)
	var data Struct.AllData
	var Allnotif = sqlite.GetMyNotif(myaccount.Nickname)
	var Allfollowing = sqlite.GetMyFollowing(myaccount)
	var Allfollower = sqlite.GetMyFollowers(myaccount)
	var followings []Struct.User
	var followers = []Struct.User{}
	for _, v := range Allfollower {
		followers = append(followers, sqlite.GetUserById(v.ID_User))
	}
	for _, v := range Allfollowing {
		followings = append(followings, sqlite.GetUserById(v.ID_Receiver))
	}
	data.Alluser = home.UserCouldBeFollow(myaccount, sqlite.GetAllUser(), followings)
	data.Allfollowing = followings
	data.Allnotif = Allnotif
	responses.SendResponsesHome(w, r, "responses succesfully", data)
})

func IsAlreadyFollowing(myUser Struct.User, following Struct.User) bool {
	return sqlite.GetFollowsByUsers(myUser.Id, following.Id) != Struct.Follow{}
}
