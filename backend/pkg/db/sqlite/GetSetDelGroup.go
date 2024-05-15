package sqlite

import (
	Struct "main/pkg/app/struct"
)

func GetGroupsByGroupID(groupId int) (Struct.Group, error) {
	query := `SELECT * FROM Groupe WHERE ID_Group = ?`
	var Group Struct.Group
	err := DB.QueryRow(query, groupId).Scan(&Group.ID_Group, &Group.GroupName, &Group.GroupImage, &Group.GroupDescription, &Group.IdCreator)
	if err != nil {
		return Struct.Group{}, err
	}
	return Group, nil
}

func CreateNewGroup(group Struct.Group) {
	query := `INSERT INTO Groupe (GroupName, GroupImage, GroupDescription, IdCreator) VALUES  (?,?,?,?)`
	//Inserer des utilisateurs dans notre table User
	_, err := DB.Exec(query, group.GroupName, group.GroupImage, group.GroupDescription, group.IdCreator)
	if err != nil {
		return
	}
}

func GetAllGroup() []Struct.Group {
	query := "SELECT * FROM Groupe ORDER BY ID_Group DESC"
	rows, err := DB.Query(query)
	if err != nil {
		return []Struct.Group{}
	}
	defer rows.Close()
	var GroupLists []Struct.Group
	for rows.Next() {
		var Group Struct.Group
		if err := rows.Scan(&Group.ID_Group, &Group.GroupName, &Group.GroupImage, &Group.GroupDescription, &Group.IdCreator); err != nil {
			continue
		}
		GroupLists = append(GroupLists, Group)
	}
	if err := rows.Err(); err != nil {
		return []Struct.Group{}
	}
	return GroupLists
}
