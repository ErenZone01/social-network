package sqlite

import (
	"fmt"
	Struct "main/pkg/struct"
)

func CreateNewEvent(event Struct.EventGroup) {
	query := `INSERT INTO EventGroup ( EventDescription, Title, EventDays, Options, ID_User, ID_Group) VALUES (?, ?, ?, ?, ?, ?)`
	_, err := DB.Exec(query, event.EventDescription, event.Title, event.EventDays, event.Option, event.ID_User, event.ID_Group)
	if err != nil {
		fmt.Println("CreateNewevent error:", err)
		return
	}
}

func GetEventByID(ID_Event int) (Struct.EventGroup, error) {
	query := `SELECT * FROM EventGroup WHERE ID_Event = ?`
	var event Struct.EventGroup
	err := DB.QueryRow(query, ID_Event).Scan(&event.ID_Event , &event.EventDescription, &event.Title, &event.EventDays, &event.Option, &event.ID_User, &event.ID_Group)
	if err != nil {
		return Struct.EventGroup{}, err
	}
	return event, nil
}

func UpdateEvent(event Struct.EventGroup) {
	_, err := DB.Exec("UPDATE EventGroup SET Options = ? WHERE ID_Group = ?", event.Option, event.ID_Group)
	if err != nil {
		fmt.Println("Error event :", err)
		return
	}
}

func GetEventByUsers(ID_Group int) []Struct.EventGroup { //recuperer les events que je suis
	var events []Struct.EventGroup
	query := "SELECT * FROM EventGroup WHERE ID_Group = ? "
	rows, err := DB.Query(query,ID_Group)
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var event Struct.EventGroup
		err := rows.Scan(&event.ID_Event, &event.EventDescription, &event.Title, &event.EventDays, &event.Option, &event.ID_User, &event.ID_Group)
		if err != nil {
			return nil
		}
		events = append(events, event)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return events
}
