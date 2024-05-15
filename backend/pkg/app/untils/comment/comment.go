package comment

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/app/session"
	Struct "main/pkg/app/struct"
	responses "main/pkg/app/untils/Responses"
	"main/pkg/app/untils/register"
	"net/http"
	"time"
)

var Comment = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if user.Error {
		return
	}
	var NewComment Struct.Comment
	NewComment.ID_User = user.Id
	var names = user.Firstname + " " + user.Lastname
	NewComment.Names = &names
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	NewComment.CreatedComment = &formattedTime

	var err = json.NewDecoder(r.Body).Decode(&NewComment)
	if err != nil {
		fmt.Println("Follow : erreur de decodage json", err)
		return
	}
	if NewComment.Images != "" {
		register.SaveImage(NewComment.Images, NewComment.ImageData, "comment")
	}

	// Appeler la fonction CreateNewPost pour enregistrer le post dans la base de données
	sqlite.CreateNewComment(NewComment)
	var Allcomment = sqlite.GetAllComments()
	fmt.Println("allcomments", Allcomment)

	responses.SendResponsesHome(w, r, "response succesfully", Allcomment)
})
