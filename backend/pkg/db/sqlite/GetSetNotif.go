package sqlite

import (
	"fmt"
	Struct "main/pkg/struct"
)

func SetNotif(msg string, Receiver string, Sender string, State string, Type string, AvatarSender string, IdGroup int) {
	var query = "INSERT INTO Notifications (Messages, Receiver, Types, States, Sender, AvatarSender, ID_Group) VALUES(?,?,?,?,?,?,?)"
	_, err := DB.Exec(query, msg, Receiver, Type, State, Sender, AvatarSender, IdGroup)
	if err != nil {
		fmt.Println("les Nottif : ", err)
		return
	}
}

func UpdateNotif(notif Struct.Notif) {
	_, err := DB.Exec("UPDATE Notifications SET States = ?, Messages = ? WHERE ID_Notification = ?", notif.States, notif.Messages, notif.ID_Notif)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func GetMyNotif(Receiver string) []Struct.Notif {
	query := "SELECT ID_Notification, Messages, Receiver, Types, States, Sender, AvatarSender, ID_Group From Notifications WHERE Receiver = ? "
	rows, err := DB.Query(query, Receiver)
	if err != nil {
		fmt.Println("Error from GetAllUser: ", err)
		return []Struct.Notif{}
	}
	defer rows.Close()
	var notifList []Struct.Notif
	for rows.Next() {
		var notif Struct.Notif
		if err := rows.Scan(&notif.ID_Notif, &notif.Messages, &notif.Receiver, &notif.Types, &notif.States, &notif.Sender, &notif.AvatarSender, &notif.ID_Group); err != nil {
			fmt.Println("GetAllUser : Error scanning row: ", err)
			continue
		}
		notifList = append(notifList, notif)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("GetAllUser : Error iterating rows: ", err)
		return []Struct.Notif{}
	}
	return notifList
}

func GetNotifById(Id int) Struct.Notif {
	var notif Struct.Notif
	query := "SELECT * From Notifications  WHERE ID_Notification = ?"
	err := DB.QueryRow(query, Id).Scan(&notif.ID_Notif, &notif.Messages, &notif.Receiver, &notif.Types, &notif.States, &notif.Sender, &notif.AvatarSender, &notif.ID_Group)
	if err != nil {
		fmt.Println("Error form GetNotif", err)
		return Struct.Notif{}
	}
	return notif
}
func DeleteNotif(ID_Notif int) {
	_, err := DB.Exec(`
DELETE FROM Notifications WHERE ID_Notification = ?
`, ID_Notif)
	if err != nil {
		fmt.Println(err)
		return
	}
}
