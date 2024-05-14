package home

import (
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	responses "main/pkg/untils/Responses"
	"main/pkg/untils/group"
	"main/pkg/untils/post"
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
	var AllUtilisateur = Alluser
	var Allpost = sqlite.GetAllPost()
	Allpost = post.GetMembers(Allpost)
	var Allfollower = sqlite.GetMyFollowers(user)
	var Allfollowing = sqlite.GetMyFollowing(user)
	var Allnotif = sqlite.GetMyNotif(user.Nickname)
	var followers []Struct.User
	var followings []Struct.User
	for _, v := range Allfollower {
		followers = append(followers, sqlite.GetUserById(v.ID_User))
	}
	for _, v := range Allfollowing {
		followings = append(followings, sqlite.GetUserById(v.ID_Receiver))
	}
	data.Alluser = UserCouldBeFollow(user, Alluser, followings)
	var Allgroup = sqlite.GetAllGroup()
	Allgroup = group.GetMembersGroup(Allgroup)
	data.UncknowGroup, data.MyGroup = group.GroupCouldBeFollow(user, Allgroup)
	data.Allpost = Allpost
	data.Allfollowers = followers
	data.Allfollowing = followings
	data.AllUtilisateur = AllUtilisateur
	data.Allcomment = sqlite.GetAllComments()
	data.Allnotif = Allnotif
	data.Myaccount = user
	responses.SendResponsesHome(w, r, "response succesfully", data)
})

func UserCouldBeFollow(myaccount Struct.User, Alluser []Struct.User, Followings []Struct.User) []Struct.User {
	var newAlluser []Struct.User
	var actif = true
	for _, v := range Alluser {
		if v.Id == myaccount.Id {
			continue
		}
		for _, v2 := range Followings {
			if v.Id == v2.Id {
				actif = false
				break
			}
		}
		if actif {
			var user = v
			var tmp = sqlite.GetFollowsByUsers(myaccount.Id, v.Id)
			if (tmp != Struct.Follow{}){ user.States = tmp.Operation}
			newAlluser = append(newAlluser, user)
		} else {
			actif = true
		}
	}
	return newAlluser
}
