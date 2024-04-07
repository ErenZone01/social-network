package home

import (
	"encoding/json"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	"net/http"
)

var Home = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	var data Struct.AllData
	var Alluser = sqlite.GetAllUser()
	var Allpost = sqlite.GetAllPost()
	var Allfollower = sqlite.GetMyFollowers(user)
	data.Alluser = Alluser
	data.Allpost = Allpost
	data.Allfollowers = Allfollower
	json.NewEncoder(w).Encode(data)
})
