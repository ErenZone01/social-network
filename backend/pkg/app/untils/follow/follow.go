package follow

import (
	"encoding/json"
	"fmt"
	"main/pkg/app/session"
	Struct "main/pkg/app/struct"
	responses "main/pkg/app/untils/Responses"
	Errors "main/pkg/app/untils/error"
	"main/pkg/app/untils/home"
	"main/pkg/db/sqlite"
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
		Errors.SendError(w, r, http.StatusOK, "Your request has already been sent")
		return
	}
	sqlite.SetFollowing(myaccount, following, operation, "person", 0)
	sqlite.SetNotif(msg, following.Nickname, myaccount.Nickname, operation, "person", myaccount.Avatar, 0, 0)
	var data Struct.AllData
	var Allnotif = sqlite.GetMyNotif(myaccount.Nickname)
	var Allfollowing = sqlite.GetMyFollowing(myaccount)
	var Allfollower = sqlite.GetMyFollowers(myaccount)
	var followings []Struct.User
	var followers = []Struct.User{}
	for _, v := range Allfollower {
		var tmp = sqlite.GetUserById(v.ID_User)
		followers = append(followers, tmp)
	}
	for _, v := range Allfollowing {
		followings = append(followings, sqlite.GetUserById(v.ID_Receiver))
	}
	data.Alluser = home.UserCouldBeFollow(myaccount, sqlite.GetAllUser(), followings)
	data.Allfollowing = followings
	data.Allnotif = Allnotif
	fmt.Println(myaccount.Nickname, " follow ", following.Nickname)
	responses.SendResponsesHome(w, r, "Your request has been sent", data)
})

func IsAlreadyFollowing(myUser Struct.User, following Struct.User) bool {
	return sqlite.GetFollowsByUsers(myUser.Id, following.Id) != Struct.Follow{}
}
