package sqlite

import (
	"fmt"
	Struct "main/pkg/struct"
)

func GetGroupsByUserID(userID int) ([]Struct.Group, error) {
	var Groups []Struct.Group
	query := `SELECT * FROM Groupe WHERE ID_Group = ?`
	rows, err := DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var Group Struct.Group
		err := rows.Scan(&Group.ID_Group , &Group.GroupName, &Group.GroupImage, &Group.GroupDescription, &Group.IdMember, &Group.IdCreator)
		if err != nil {
			return nil, err
		}
		Groups = append(Groups, Group)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return Groups, nil
}

func CreateNewGroup(group Struct.Group) {
	query := `INSERT INTO Groupe (GroupName, GroupImage, GroupDescription, IdMember, IdCreator) VALUES  (?,?,?,?,?)`
	//Inserer des utilisateurs dans notre table User
	_, err := DB.Exec(query, group.GroupName, group.GroupImage, group.GroupDescription, group.IdMember, group.IdCreator)
	if err != nil {
		fmt.Println("SetGroup : ", err)
		return
	}
}

func GetAllGroup() []Struct.Group {
	query := "SELECT * FROM Groupe ORDER BY ID_Group DESC"
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("Error from GetAllGroup: ", err)
		return []Struct.Group{}
	}
	defer rows.Close()
	var GroupLists []Struct.Group
	for rows.Next() {
		var Group Struct.Group
		if err := rows.Scan(&Group.ID_Group , &Group.GroupName, &Group.GroupImage, &Group.GroupDescription, &Group.IdMember, &Group.IdCreator); err != nil {
			fmt.Println("GetAllGroup : Error scanning row: ", err)
			continue
		}
		GroupLists = append(GroupLists, Group)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("GetAllGroup : Error iterating rows: ", err)
		return []Struct.Group{}
	}
	return GroupLists
}
