package sqlite

import (
	"fmt"
	Struct "main/pkg/struct"
)

func GetPostsByUserID(userID int) ([]Struct.Post, error) {
	var posts []Struct.Post
	query := `SELECT * FROM Post WHERE ID_User = ? AND Types = ?`
	rows, err := DB.Query(query, userID, "Post")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var post Struct.Post
		err := rows.Scan(&post.Id, &post.Names, &post.Content, &post.Title, &post.Image, &post.ID_User, &post.Privacy, &post.ID_Group, &post.Types, &post.CreatedPost)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return posts, nil
}

func CreateNewPost(post Struct.Post) {
	query := `INSERT INTO Post (Names, Content, Title, Imagee, ID_User, Privacy, ID_Group, Types, CreatedPost) VALUES  (?,?,?,?,?,?,?,?,?)`
	//Inserer des utilisateurs dans notre table User
	_, err := DB.Exec(query, post.Names, post.Content, post.Title, post.Image, post.ID_User, post.Privacy, post.ID_Group, post.Types, post.CreatedPost)
	if err != nil {
		fmt.Println("SetPost : ", err)
		return
	}
}

func DeletePost(id int) {
	_, err := DB.Exec(`
	DELETE FROM Post WHERE ID_Post = ?
	`, id)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func GetAllPost() []Struct.Post {
	query := "SELECT * FROM Post WHERE Types=?  ORDER BY ID_Post DESC"
	rows, err := DB.Query(query, "Post")
	if err != nil {
		fmt.Println("Error from GetAllPost: ", err)
		return []Struct.Post{}
	}
	defer rows.Close()
	var PostLists []Struct.Post
	for rows.Next() {
		var post Struct.Post
		if err := rows.Scan(&post.Id, &post.Names, &post.Content, &post.Title, &post.Image, &post.ID_User, &post.Privacy, &post.ID_Group, &post.Types, &post.CreatedPost); err != nil {
			fmt.Println("GetAllUser : Error scanning row: ", err)
			continue
		}
		PostLists = append(PostLists, post)
	}
	return PostLists
}

func GetAllPostByGroup(ID_Group int) []Struct.Post {
	query := "SELECT * FROM Post WHERE ID_Group = ? ORDER BY ID_Post DESC"
	rows, err := DB.Query(query, ID_Group)
	if err != nil {
		fmt.Println("GetAllPostGroup2 : Error scanning row: ", err)
		return []Struct.Post{}
	}
	defer rows.Close()
	var PostLists []Struct.Post
	for rows.Next() {
		var post Struct.Post
		if err := rows.Scan(&post.Id, &post.Names, &post.Content, &post.Title, &post.Image, &post.ID_User, &post.Privacy, &post.ID_Group, &post.Types, &post.CreatedPost); err != nil {
			fmt.Println("GetAllPostGroup2 : Error scanning row: ", err)
			continue
		}
		PostLists = append(PostLists, post)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("GetAllPostGroup3 : Error iterating rows: ", err)
		return []Struct.Post{}
	}
	return PostLists
}
