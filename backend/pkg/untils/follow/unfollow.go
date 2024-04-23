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

var UnFollowHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
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
	var followings []Struct.User
	var follow = sqlite.GetFollowsByUsers(myaccount.Id, following.Id)
	sqlite.DeleteFollow(follow.ID_Follow)
	var Allfollowing = sqlite.GetMyFollowing(myaccount)
	for _, v := range Allfollowing {
		followings = append(followings, sqlite.GetUserById(v.ID_Receiver))
	}
	responses.SendResponsesHome(w, r, "responses succesfully", followings)
})
