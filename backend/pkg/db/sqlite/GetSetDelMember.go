package sqlite

import (
	Struct "main/pkg/app/struct"
)

func CreateNewMember(membres Struct.Members) {
	query := `INSERT INTO Members (ID_Post, ID_Group, ID_User, ID_Event) VALUES (?,?,?,?)`
	//Inserer des utilisateurs dans notre table User
	_, err := DB.Exec(query, membres.ID_Post, membres.ID_Group, membres.ID_User, membres.ID_Event)
	if err != nil {
		return
	}
}

func GetAllMemberPost(ID_Post int) []Struct.Members {
	query := "SELECT * FROM Members WHERE ID_Post = ?"
	rows, err := DB.Query(query, ID_Post)
	if err != nil {
		return []Struct.Members{}
	}
	defer rows.Close()
	var MemberLists []Struct.Members
	for rows.Next() {
		var membres Struct.Members
		if err := rows.Scan(&membres.ID_Member, &membres.ID_Post, &membres.ID_Group, &membres.ID_User, &membres.ID_Event); err != nil {
			continue
		}
		MemberLists = append(MemberLists, membres)
	}
	if err := rows.Err(); err != nil {
		return []Struct.Members{}
	}
	return MemberLists
}
func GetAllMemberGroup(ID_Group int, ID_Event int) []Struct.Members {
	query := "SELECT * FROM Members WHERE ID_Group = ? AND ID_Event = ?"
	rows, err := DB.Query(query, ID_Group, ID_Event)
	if err != nil {
		return []Struct.Members{}
	}
	defer rows.Close()
	var MemberLists []Struct.Members
	for rows.Next() {
		var membres Struct.Members
		if err := rows.Scan(&membres.ID_Member, &membres.ID_Post, &membres.ID_Group, &membres.ID_User, &membres.ID_Event); err != nil {
			continue
		}
		MemberLists = append(MemberLists, membres)
	}
	if err := rows.Err(); err != nil {
		return []Struct.Members{}
	}
	return MemberLists
}

func GetAllMemberOfAnyGroup(ID_Group int) []int {
	query := "SELECT * FROM Members WHERE ID_Group = ?"
	rows, err := DB.Query(query, ID_Group)
	if err != nil {
		return []int{}
	}
	defer rows.Close()
	var ID_MemberLists []int
	for rows.Next() {
		var membres Struct.Members
		if err := rows.Scan(&membres.ID_Member, &membres.ID_Post, &membres.ID_Group, &membres.ID_User, &membres.ID_Event); err != nil {
			continue
		}
		ID_MemberLists = append(ID_MemberLists, membres.ID_User)
	}
	if err := rows.Err(); err != nil {
		return []int{}
	}
	return ID_MemberLists
}
