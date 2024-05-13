package sqlite

import (
	
	"fmt"
	Struct "main/pkg/struct"
)


func CreateNewEvent(event Struct.EventGroup) {
	query := `INSERT INTO EventGroup (id_event, eventdescription, title, eventdays, option, id_user, id_group) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := DB.Exec(query, event.ID_Event, event.EventDescription, event.Title, event.EventDays, event.Option, event.ID_User, event.ID_Group)
	if err != nil {
		fmt.Println("CreateNewevent error:", err)
		return
	}
}

