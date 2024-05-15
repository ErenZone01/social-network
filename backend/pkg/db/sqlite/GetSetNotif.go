package sqlite

import (
	"fmt"
	Struct "main/pkg/app/struct"
)

func SetNotif(msg string, Receiver string, Sender string, State string, Type string, AvatarSender string, IdGroup int, ID_Event int) {
	var query = "INSERT INTO Notifications (Messages, Receiver, Types, States, Sender, AvatarSender, ID_Group, ID_Event) VALUES(?,?,?,?,?,?,?,?)"
	_, err := DB.Exec(query, msg, Receiver, Type, State, Sender, AvatarSender, IdGroup, ID_Event)
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
	query := "SELECT * From Notifications WHERE Receiver = ? "
	rows, err := DB.Query(query, Receiver)
	if err != nil {
		fmt.Println("Error from GetAllUser: ", err)
		return []Struct.Notif{}
	}
	defer rows.Close()
	var notifList []Struct.Notif
	for rows.Next() {
		var notif Struct.Notif
		if err := rows.Scan(&notif.ID_Notif, &notif.Messages, &notif.Receiver, &notif.Types, &notif.States, &notif.Sender, &notif.AvatarSender, &notif.ID_Group, &notif.ID_Event); err != nil {
			fmt.Println("GetAllUser : Error scanning row: ", err)
			continue
		}
		notifList = append(notifList, notif)
	}
	if err := rows.Err(); err != nil {
		return []Struct.Notif{}
	}
	return notifList
}

func GetNotifPersonnal(Receiver string, Sender string, IdGroup int)Struct.Notif {
	var notif Struct.Notif
	query := "SELECT * From Notifications WHERE ID_Group = ? AND Receiver = ? AND Sender = ?"
	err := DB.QueryRow(query, IdGroup , Receiver, Sender).Scan(&notif.ID_Notif, &notif.Messages, &notif.Receiver, &notif.Types, &notif.States, &notif.Sender, &notif.AvatarSender, &notif.ID_Group, &notif.ID_Event)
	if err != nil {
		return Struct.Notif{}
	}
	return notif
}
func GetAllPersonnalNotif(Receiver string, Sender string, IdGroup int)Struct.Notif {
	var notif Struct.Notif
	query := "SELECT * From Notifications WHERE ID_Group = ? AND Receiver = ? AND Sender = ? AND States = ?"
	err := DB.QueryRow(query, IdGroup , Receiver, Sender, "pending").Scan(&notif.ID_Notif, &notif.Messages, &notif.Receiver, &notif.Types, &notif.States, &notif.Sender, &notif.AvatarSender, &notif.ID_Group, &notif.ID_Event)
	if err != nil {
		return Struct.Notif{}
	}
	return notif
}

func GetNotifById(Id int) Struct.Notif {
	var notif Struct.Notif
	query := "SELECT * From Notifications  WHERE ID_Notification = ?"
	err := DB.QueryRow(query, Id).Scan(&notif.ID_Notif, &notif.Messages, &notif.Receiver, &notif.Types, &notif.States, &notif.Sender, &notif.AvatarSender, &notif.ID_Group, &notif.ID_Event)
	if err != nil {
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

func DeleteNotifByReceiver(Receiver string) {
	_, err := DB.Exec(`
DELETE FROM Notifications WHERE Receiver = ? AND  States = ?
`, Receiver, "true")
	if err != nil {
		fmt.Println(err)
		return
	}
}

