package home

import (
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	responses "main/pkg/untils/Responses"
	"net/http"
)

var Home = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if user.Error {
		return
	}
	var data Struct.AllData
	var Alluser = sqlite.GetAllUser()
	var Allpost = sqlite.GetAllPost()
	var Allfollower = sqlite.GetMyFollowers(user)
	var Allfollowing = sqlite.GetMyFollowing(user)
	var followers []Struct.User
	var followings []Struct.User
	for _, v := range Allfollower {
		followers = append(followers, sqlite.GetUserById(v.ID_Receiver))
	}
	for _, v := range Allfollowing {
		followings = append(followings, sqlite.GetUserById(v.ID_Receiver))
	}
	data.Alluser = Alluser
	data.Allpost = Allpost
	data.Allfollowers = followers
	data.Allfollowing = followings
	data.Myaccount = user
	responses.SendResponsesHome(w, r, data)
})
