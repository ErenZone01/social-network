package home

import (
	"encoding/json"
	"fmt"
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
	//get All followers
	var Allfollower = sqlite.GetMyFollowers(user)
	var followers[]Struct.User
	for _, v := range Allfollower {
		followers = append(followers, sqlite.GetUserById(v.ID_Receiver))
	}
	fmt.Println("Allfollower : ", Allfollower)
	data.Alluser = Alluser
	data.Allpost = Allpost
	data.Allfollowers = followers
	w.Header().Set("Content-Type", "application/json") // Définir le type de contenu de la réponse comme JSON
	json.NewEncoder(w).Encode(data)
})
