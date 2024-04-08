package sqlite

import (
	"fmt"
	Struct "main/pkg/struct"
)

func GetUser(login string) Struct.User {
	var user Struct.User
	query := "SELECT ID_User, Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy From User  WHERE Email = ? OR Nickname = ?"
	err := DB.QueryRow(query, login, login).Scan(&user.Id, &user.Email, &user.Nickname, &user.Password, &user.Firstname, &user.Lastname, &user.Birth, &user.Avatar, &user.About, &user.Privacy)
	if err != nil {
		fmt.Println("Error form GetUser", err)
		return Struct.User{}
	}
	return user
}

func GetUserById(Id int) Struct.User {
	var user Struct.User
	query := "SELECT ID_User, Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy From User  WHERE ID_User = ?"
	err := DB.QueryRow(query, Id).Scan(&user.Id, &user.Email, &user.Nickname, &user.Password, &user.Firstname, &user.Lastname, &user.Birth, &user.Avatar, &user.About, &user.Privacy)
	if err != nil {
		fmt.Println("Error form GetUser", err)
		return Struct.User{}
	}
	return user
}


func CreateNewUser(user Struct.User) {
	query := `INSERT INTO User (Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy) VALUES (?,?,?,?,?,?,?,?,?)`
	//Inserer des utilisateurs dans notre table User
	_, err := DB.Exec(query, user.Email, user.Nickname, user.Password, user.Firstname, user.Lastname, user.Birth, user.Avatar, user.About, user.Privacy)
	if err != nil {
		fmt.Println("SetUser : ", err)
		return
	}
}

func DeleteUser(id int) {
	_, err := DB.Exec(`
	DELETE FROM User WHERE ID_User = ?
	`, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func UpdateUser(user Struct.User) {
	_, err := DB.Exec("UPDATE User SET Actif = ? WHERE ID_User = ?", user.Actif, user.Id)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

func GetAllUser() []Struct.User {
	query := "SELECT ID_User, Email, Nickname, Passwords, Firstname, Lastname, Birth, Avatar, About, Privacy From User"
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("Error from GetAllUser: ", err)
		return []Struct.User{}
	}
	defer rows.Close()
	var usersList []Struct.User
	for rows.Next() {
		var user Struct.User
		if err := rows.Scan(&user.Id, &user.Email, &user.Nickname, &user.Password, &user.Firstname, &user.Lastname, &user.Birth, &user.Avatar, &user.About, &user.Privacy); err != nil {
			fmt.Println("GetAllUser : Error scanning row: ", err)
			continue
		}
		usersList = append(usersList, user)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("GetAllUser : Error iterating rows: ", err)
		return []Struct.User{}
	}
	return usersList
}
