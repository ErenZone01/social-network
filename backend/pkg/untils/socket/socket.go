package socket

import (
	"fmt"
	"log"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]Struct.User)
var Broadcast = make(chan Struct.Render)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// log.Println("r",r)
		// return true
		// Vérifie l'origine de la requête WebSocket
		origin := r.Header.Get("Origin")
		// Ici, tu peux autoriser l'origine de ton serveur Vue.js
		return strings.Contains(origin, "http://localhost:5173")
	},
}

func NumberInSlice(nbr int, slice []int) bool {
	for _, element := range slice {
		if element == nbr {
			return true
		}
	}
	return false
}

var Socket = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("%s", err)
	}
	defer ws.Close()
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	// if user.Error {
	// 	log.Println("Do you have an account")
	// 	return
	// }
	// log.Println("First step In websocket for ", user.Nickname, "!!!")
	if user.Nickname != "" {
		clients[ws] = user
	}
	// log.Println("len clients", len(clients))

	go handleMessages()

	for {
		var chatSlice Struct.Render
		err := ws.ReadJSON(&chatSlice)
		if err != nil {
			if _, ok := err.(*websocket.CloseError); ok {
				// log.Printf("connection closed with status %v due to %s", closeMsg.Code, closeMsg.Text)
			} else {
				fmt.Println("read error:", err)
			}
			delete(clients, ws)
			break
		}

		// fmt.Println("received from client:", chatSlice)
		// log.Println("chat slice", chatSlice.Payload)
		Broadcast <- chatSlice
	}

})

func handleMessages() {
	for v := range Broadcast {
		chatSlice := v
		if chatSlice.Payload == "chat" {
			for client := range clients {
				user := clients[client]
				if chatSlice.To == user.Id {
					// log.Println("for ", chatSlice.To, " by ", user.Nickname, " id ", user.Id)
					err := client.WriteJSON(chatSlice)
					if err != nil {
						client.Close()
						delete(clients, client)
					}
				}
			}
		} else if chatSlice.Payload == "group" {
			idCreator, _ := sqlite.GetGroupsByGroupID(chatSlice.To)
			members := sqlite.GetAllMemberOfAnyGroup(chatSlice.To)
			members = append(members, idCreator.IdCreator)
			for client := range clients {
				//search members
				user := clients[client]
				if user.Id != chatSlice.From {
					if NumberInSlice(user.Id, members) {
						log.Println("for in group", chatSlice.To, " by ", user.Nickname, " id ", user.Id)
						err := client.WriteJSON(chatSlice)
						if err != nil {
							client.Close()
							delete(clients, client)
						}
					}
				}
			}
		}
	}
}
