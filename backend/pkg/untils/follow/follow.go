package follow

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	responses "main/pkg/untils/Responses"
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
	var operation string;
	if following.Privacy == "public"{
		operation = "true"
	}else{
		operation = "pending"
	}
	//tempFollowing := sqlite.GetMyFollowing(myaccount)


	sqlite.SetFollowing(myaccount, following, operation, "person", 0)
	sqlite.SetNotif("following you", following.Nickname, myaccount.Nickname, operation, "person", myaccount.Avatar,0)
	var data Struct.AllData
	var Allfollower = sqlite.GetMyFollowing(myaccount)
	var Allnotif = sqlite.GetMyNotif(myaccount.Nickname)
	var followings []Struct.User
	for _, v := range Allfollower {
		followings = append(followings, sqlite.GetUserById(v.ID_Receiver))
	}
	data.Allfollowing =followings
	data.Allnotif = Allnotif
	responses.SendResponsesHome(w,r,"responses succesfully",data)
})

func IsAlreadyFollowing(Allfollowing []Struct.Follow, User Struct.User, Sender Struct.User){}