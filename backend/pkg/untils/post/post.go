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
	register.SaveImage(NewPost.Image, NewPost.ImageData, "post")

	// Appeler la fonction CreateNewPost pour enregistrer le post dans la base de données
	sqlite.CreateNewPost(NewPost)
	var Allpost = sqlite.GetAllPost()
	fmt.Println("allposts", Allpost)

	responses.SendResponsesHome(w, r, "response succesfully", Allpost)
})
