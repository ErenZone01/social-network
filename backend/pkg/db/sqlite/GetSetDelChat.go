package sqlite

import (
	"fmt"
	Struct "main/pkg/app/struct"
)

func CreateChat(chat Struct.Chat) {
	query := `INSERT INTO Chat(Content, ID_Group,ID_Receiver,ID_User, Types) VALUES  (?,?,?,?,?)`
	//Inserer des utilisateurs dans notre table User
	_, err := DB.Exec(query, chat.Content, chat.ID_Group, chat.ID_Receiver, chat.ID_User, chat.Types)
	if err != nil {
		fmt.Println("Set Chat : ", err)
		return
	}
	// s := "hello"
	// socket.Broadcast <- []byte(s)
}

//	func DeletePost(id int) {
//		_, err := DB.Exec(`
//		DELETE FROM Post WHERE ID_Post = ?
//		`, id)
//		if err != nil {
//			fmt.Println(err)
//			return
//		}
//	}
func GetTheLatestChat(sender int, types string) Struct.User {
	var userID int
	// query := "SELECT * FROM Chat where ID_Receiver=?  OR  ID_User=?  ORDER BY ID_Chat DESC"
	query := "SELECT * FROM Chat WHERE (ID_Receiver = ? OR ID_User = ?) AND Types = ? ORDER BY ID_Chat DESC LIMIT 1"

	rows, err := DB.Query(query, sender, sender, types)
	if err != nil {
		fmt.Println("Error from Get latest Chat: ", err)
		return Struct.User{}
	}
	defer rows.Close()
	var chat Struct.Chat
	for rows.Next() {
		if err := rows.Scan(&chat.ID_Chat, &chat.ID_User, &chat.ID_Receiver, &chat.Content, &chat.Types, &chat.ID_Group); err != nil {
			fmt.Println("Get latest chat: Error scanning row: ", err)
			continue
		}
		// log.Println("latest ", chat)
		// Chats = append(Chats, chat)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("GetAllChat : Error iterating rows: ", err)
		return Struct.User{}
	}
	if sender != chat.ID_User {
		userID = chat.ID_User
	} else {
		userID = chat.ID_Receiver
	}
	return GetUserById(userID)

}

func GetChatsGroup(idGroup int, types string) []Struct.Chat {
	query := "SELECT * FROM Chat WHERE ID_Receiver=? AND Types=?"
	rows, err := DB.Query(query, idGroup, types)
	if err != nil {
		fmt.Println("Error from GetAllChatGroup: ", err)
		return []Struct.Chat{}
	}
	defer rows.Close()
	var Chats []Struct.Chat
	for rows.Next() {
		var chat Struct.Chat
		if err := rows.Scan(&chat.ID_Chat, &chat.ID_User, &chat.ID_Receiver, &chat.Content, &chat.Types, &chat.ID_Group); err != nil {
			fmt.Println("GetAllChatGroup: Error scanning row: ", err)
			continue
		}
		// log.Println("one chat", chat)
		Chats = append(Chats, chat)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("GetAllChatGroup : Error iterating rows: ", err)
		return []Struct.Chat{}
	}
	return Chats
}

func GetAllChat(sender, receiver int, types string) []Struct.Chat {

	query := "SELECT * FROM Chat WHERE (ID_Receiver=? AND ID_User=?) OR (ID_Receiver=? AND ID_User=?) AND  Types=?"
	rows, err := DB.Query(query, sender, receiver, receiver, sender, types)
	if err != nil {
		fmt.Println("Error from GetAllChat: ", err)
		return []Struct.Chat{}
	}
	defer rows.Close()
	var Chats []Struct.Chat
	for rows.Next() {
		var chat Struct.Chat
		if err := rows.Scan(&chat.ID_Chat, &chat.ID_User, &chat.ID_Receiver, &chat.Content, &chat.Types, &chat.ID_Group); err != nil {
			fmt.Println("GetAllChat: Error scanning row: ", err)
			continue
		}
		// log.Println("one chat", chat)
		Chats = append(Chats, chat)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("GetAllChat : Error iterating rows: ", err)
		return []Struct.Chat{}
	}
	return Chats
}
