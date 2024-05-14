package sqlite

import (
	"fmt"
	Struct "main/pkg/struct"
)
//
func GetCommentsByPostID(postID int) ([]Struct.Comment, error) {
	var comments []Struct.Comment
	query := `SELECT * FROM Comment WHERE ID_Post = ?`
	rows, err := DB.Query(query, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var comment Struct.Comment
		err := rows.Scan(&comment.ID_Comment, &comment.Content, &comment.Images, &comment.ID_User, &comment.ID_Post, &comment.ID_Group, &comment.Types, &comment.CreatedComment)
		if err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return comments, nil
}

func CreateNewComment(comment Struct.Comment) {
	query := `INSERT INTO Comment (content, Images, ID_User, ID_Post, ID_Group, Types, Names) VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := DB.Exec(query, comment.Content, comment.Images, comment.ID_User, comment.ID_Post, comment.ID_Group, comment.Types,comment.Names)
	if err != nil {
		fmt.Println("CreateNewComment error:", err)
		return
	}
}

func DeleteComment(id int) {
	_, err := DB.Exec(`
		DELETE FROM Comment WHERE ID_Comment = ?
	`, id)
	if err != nil {
		fmt.Println("DeleteComment error:", err)
		return
	}
}

func GetAllComments() []Struct.Comment {
	query := "SELECT * FROM Comment ORDER BY ID_Comment DESC"
	rows, err := DB.Query(query)
	if err != nil {
		fmt.Println("GetAllComments error:", err)
		return []Struct.Comment{}
	}
	defer rows.Close()
	var commentList []Struct.Comment
	for rows.Next() {
		var comment Struct.Comment
		if err := rows.Scan(&comment.ID_Comment, &comment.Content, &comment.Images, &comment.ID_User, &comment.ID_Post, &comment.ID_Group, &comment.Types, &comment.CreatedComment,&comment.Names); err != nil {
			fmt.Println("GetAllComments: Error scanning row:", err)
			continue
		}
		commentList = append(commentList, comment)
	}
	if err := rows.Err(); err != nil {
		fmt.Println("GetAllComments: Error iterating rows:", err)
		return []Struct.Comment{}
	}
	return commentList
}
