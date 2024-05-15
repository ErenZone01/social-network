package follow

import (
	"encoding/json"
	"fmt"
	"main/pkg/app/session"
	Struct "main/pkg/app/struct"
	responses "main/pkg/app/untils/Responses"
	"main/pkg/db/sqlite"
	"net/http"
)

var UnFollowHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var myaccount = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if myaccount.Error {
		return
	}
	var following Struct.User
	var err = json.NewDecoder(r.Body).Decode(&following)
	if err != nil {
		fmt.Println("UnFollow : erreur de decodage json", err)
		return
	}
	var followings []Struct.User
	var follow = sqlite.GetFollowsByUsers(myaccount.Id, following.Id)
	sqlite.DeleteFollow(follow.ID_Follow)
	var Allfollowing = sqlite.GetMyFollowing(myaccount)
	for _, v := range Allfollowing {
		followings = append(followings, sqlite.GetUserById(v.ID_Receiver))
	}
	fmt.Println(myaccount.Nickname, " unfollow ", following.Nickname)
	responses.SendResponsesHome(w, r, "responses succesfully", followings)
})
