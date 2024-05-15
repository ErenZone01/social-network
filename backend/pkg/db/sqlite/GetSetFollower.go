package sqlite

import (
	"database/sql"
	Struct "main/pkg/app/struct"
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
		err := rows.Scan(&follow.ID_Follow, &follow.ID_User, &follow.ID_Receiver, &follow.Privacy, &follow.Operation, &follow.Types, &follow.ID_Group)
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
		err := rows.Scan(&follow.ID_Follow, &follow.ID_User, &follow.ID_Receiver, &follow.Privacy, &follow.Operation, &follow.Types, &follow.ID_Group)
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
		return
	}
}
func DeleteFollow(id int) {
	_, err := DB.Exec(`
	DELETE FROM Follow WHERE ID_Follow = ?
	`, id)
	if err != nil {
		return
	}
}

// GetFollowByUsers récupère les informations de suivi pour deux utilisateurs spécifiques.
func GetFollowByUsers(Id_Sender, Id_Receiver int) Struct.Follow {
	var follow Struct.Follow
	query := "SELECT * FROM Follow WHERE (ID_User = ? AND ID_Receiver = ?) OR (ID_Receiver = ? AND ID_User = ?)"
	err := DB.QueryRow(query, Id_Sender, Id_Receiver, Id_Receiver, Id_Sender).Scan(&follow.ID_Follow, &follow.ID_User, &follow.ID_Receiver, &follow.Privacy, &follow.Operation, &follow.Types, &follow.ID_Group)
	if err != nil {
		return Struct.Follow{}
	}
	return follow
}

// GetFollowByUsers récupère les informations de suivi pour deux utilisateurs spécifiques.
func GetFollowsByUsers(Id_Sender, Id_Receiver int) Struct.Follow {
	var follow = Struct.Follow{} // Utilisez un pointeur pour pouvoir renvoyer nil
	query := "SELECT * FROM Follow WHERE (ID_User = ? AND ID_Receiver = ?)"
	err := DB.QueryRow(query, Id_Sender, Id_Receiver).Scan(&follow.ID_Follow, &follow.ID_User, &follow.ID_Receiver, &follow.Privacy, &follow.Operation, &follow.Types, &follow.ID_Group)
	if err != nil {
		if err == sql.ErrNoRows {
			// Aucune ligne trouvée, renvoie nil
			return Struct.Follow{}
		}
		return Struct.Follow{}
	}
	return follow
}

func UpdateFollow(follow Struct.Follow) {
	_, err := DB.Exec("UPDATE Follow SET Operation = ? WHERE ID_Follow = ?", follow.Operation, follow.ID_Follow)
	if err != nil {
		return
	}
}
