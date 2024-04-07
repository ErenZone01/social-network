package sqlite

import (
	"fmt"
	Struct "main/pkg/struct"
)

func GetMyFollowers(user Struct.User) []Struct.Follow { // recuperer les gens qui me suive
	var follows []Struct.Follow
	query := `SELECT * FROM Follow WHERE ID_Receiver = ? AND Operation = ?`
	rows, err := DB.Query(query, user.Id, "true")
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var follow Struct.Follow
		err := rows.Scan(&follow.ID_Follow, &follow.ID_User, &follow.ID_Receiver, &follow.Privacy, &follow.Operation, &follow.Privacy, &follow.Types, &follow.ID_Group)
		if err != nil {
			return nil
		}
		follows = append(follows, follow)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return follows
}

func GetMyFollowing(user Struct.User) []Struct.Follow { //recuperer les gens que je suis
	var follows []Struct.Follow
	query := `SELECT * FROM Follow WHERE ID_User = ? AND Operation = ?`
	rows, err := DB.Query(query, user.Id, "true")
	if err != nil {
		return nil
	}
	defer rows.Close()
	for rows.Next() {
		var follow Struct.Follow
		err := rows.Scan(&follow.ID_Follow, &follow.ID_User, &follow.ID_Receiver, &follow.Privacy, &follow.Operation, &follow.Privacy, &follow.Types, &follow.ID_Group)
		if err != nil {
			return nil
		}
		follows = append(follows, follow)
	}
	if err := rows.Err(); err != nil {
		return nil
	}
	return follows
}
func SetFollowing(follower Struct.User, following Struct.User, Operation string, Type string, ID_Group int) {
	var query = "INSERT INTO Follow (ID_User, ID_Receiver, Privacy, Operation, Types, ID_Group) VALUES(?,?,?,?,?,?)"
	_, err := DB.Exec(query, follower.Id, following.Id, following.Privacy, Operation, Type, ID_Group)
	if err != nil {
		fmt.Println("SetFollow : ", err)
		return
	}
}
