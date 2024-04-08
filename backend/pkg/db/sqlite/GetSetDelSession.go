package sqlite

import (
	"fmt"
	Struct "main/pkg/struct"
)

func NewSession(session Struct.Session) {
	//Inserer des sessions dans notre table Session
	_, err := DB.Exec(`
    INSERT INTO AllSession ( user_id,  Session_value) VALUES (?,?)
`, session.Users_id, session.Value)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func DataSession() ([]Struct.Session, error) {
	query := "SELECT id, user_id, Session_value FROM AllSession"
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}
	defer rows.Close()
	var sesList []Struct.Session
	for rows.Next() {
		var ses Struct.Session
		if err := rows.Scan(&ses.Id, &ses.Users_id, &ses.Value); err != nil {
			fmt.Println("Error scanning row:", err)
			continue
		}
		sesList = append(sesList, ses)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("Error iterating rows:", err)
		return nil, err
	}
	return sesList, nil
}
func DeleteSession(id int) {
	_, err := DB.Exec(`
	DELETE FROM AllSession WHERE user_id = ?
	`, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}
