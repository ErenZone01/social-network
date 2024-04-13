package Struct

type AllData struct {
	Myaccount    User
	Alluser      []User
	Allpost      []Post
	Allfollowers []User
	Allfollowing []User
	Allnotif     []Notif
}

type User struct {
	Id         int    `json:"ID"`
	Email      string `json:"Email"`
	Nickname   string `json:"Nickname"`
	Password   string `json:"Password"`
	Firstname  string `json:"Firstname"`
	Lastname   string `json:"Lastname"`
	Birth      string `json:"Birth"`
	Avatar     string `json:"Avatar"`
	AvatarData []byte `json:"AvatarData"`
	About      string `json:"About"`
	Privacy    string `json:"Privacy"`
	Error      bool
	Actif      string
}

type Post struct {
	Id       int    `json:"ID"`
	Content  string `json:"Content"`
	Title    string `json:"Title"`
	Images   string `json:"Images"`
	ID_User  int    `json:"ID_User"`
	Privacy  string `json:"Privacy"`
	ID_Group string `json:"ID_Group"`
	Types    string `json:"Types"`
}

type Session struct {
	Id       int
	Users_id int
	Value    string
}

type FetchMsg struct {
	Types string
	Msg   string
	Data  interface{}
}

type Follow struct {
	ID_Follow   int
	ID_User     int
	ID_Receiver int
	Privacy     string
	Operation   string
	Types       string
	ID_Group    int
}

type Notif struct {
	ID_Notif    int
	Messages    string
	ID_Receiver int
	Types       string
	States      string
	IdUser      int
	ID_Group    int
}
