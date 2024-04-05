package Struct

type User struct {
	Id        int    `json:"ID"`
	Email     string `json:"Email"`
	Nickname  string `json:"Nickname"`
	Password  string `json:"Password"`
	Firstname string `json:"Firstname"`
	Lastname  string `json:"Lastname"`
	Birth     string `json:"Birth"`
	Avatar    string `json:"Avatar"`
	About     string `json:"About"`
	Privacy   string `json:"Privacy"`
	Error     string
}

type Session struct {
	Id       int
	Users_id int
	Value    string
}
