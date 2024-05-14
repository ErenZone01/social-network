package post

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	responses "main/pkg/untils/Responses"
	"main/pkg/untils/register"
	"net/http"
	"time"
)

var Post = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if user.Error {
		return
	}
	var NewPost Struct.Post
	var NewMember Struct.Members
	NewPost.ID_User = user.Id
	var names = user.Firstname + " " + user.Lastname
	NewPost.Names = &names
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	NewPost.CreatedPost = &formattedTime

	var err = json.NewDecoder(r.Body).Decode(&NewPost)
	if err != nil {
		fmt.Println("Follow : erreur de decodage json", err)
		return
	}
	if NewPost.Image != "" {
		register.SaveImage(NewPost.Image, NewPost.ImageData, "post")
	}
	// Appeler la fonction CreateNewPost pour enregistrer le post dans la base de données
	sqlite.CreateNewPost(NewPost)
	var Allpost []Struct.Post
	if NewPost.Types == "Post" {
		Allpost = sqlite.GetAllPost()
		for _, v := range NewPost.MembersPost {
			NewMember.ID_User = v
			NewMember.ID_Post = Allpost[0].Id
			NewMember.ID_Group = 0
			sqlite.CreateNewMember(NewMember)
		}
		Allpost = GetMembers(Allpost)
	} else {
		Allpost = sqlite.GetAllPostByGroup(NewPost.ID_Group)
	}

	fmt.Println("New post added by : ", user.Nickname)

	responses.SendResponsesHome(w, r, "response succesfully", Allpost)
})

func GetMembers(Allpost []Struct.Post) []Struct.Post {
	for i, v := range Allpost {
		var tmp = sqlite.GetAllMemberPost(v.Id)
		if len(tmp) != 0 {
			for _, m := range tmp {
				Allpost[i].MembersPost = append(Allpost[i].MembersPost, m.ID_User)
			}
		}
	}
	return Allpost
}
