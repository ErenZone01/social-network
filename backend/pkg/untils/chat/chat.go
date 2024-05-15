package chat

import (
	"encoding/json"
	"fmt"
	"log"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	responses "main/pkg/untils/Responses"
	"net/http"
)

var Chat = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if user.Error {
		return
	}
	var NewChat Struct.Chat
	NewChat.ID_User = user.Id

	var err = json.NewDecoder(r.Body).Decode(&NewChat)
	if err != nil {
		fmt.Println("Follow : erreur de decodage json", err)
		return
	}
	log.Println("new chat", NewChat)
	// Appeler la fonction CreateNewChat pour enregistrer le Chat dans la base de données
	sqlite.CreateChat(NewChat)
	var AllChat = []Struct.Chat{}
	if NewChat.Types == "chat" {
		AllChat = sqlite.GetAllChat(NewChat.ID_User, NewChat.ID_Receiver, NewChat.Types)
	} else if NewChat.Types == "group" {
		AllChat = sqlite.GetChatsGroup(NewChat.ID_Receiver, "group")
	}
	// fmt.Println("New Chat added by : ", user.Nickname)
	responses.SendResponsesHome(w, r, "response succesfully", AllChat)

})

var Chats = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if user.Error {
		return
	}
	var Target Struct.User
	var err = json.NewDecoder(r.Body).Decode(&Target)
	if err != nil {
		fmt.Println("Target ID: erreur de decodage json", err)
		return
	}
	var AllChatsBetween = sqlite.GetAllChat(user.Id, Target.Id, "chat")
	responses.SendResponsesHome(w, r, "response succesfully", AllChatsBetween)

})

var ChatsGroup = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if user.Error {
		return
	}
	var Target Struct.Group
	var err = json.NewDecoder(r.Body).Decode(&Target)
	if err != nil {
		fmt.Println("Target ID: erreur de decodage json", err)
		return
	}
	var AllChatsBetween = sqlite.GetChatsGroup(Target.ID_Group, "group")
	responses.SendResponsesHome(w, r, "response succesfully", AllChatsBetween)
})
