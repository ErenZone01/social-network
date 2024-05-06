package Struct

type AllData struct {
	Myaccount      User
	Alluser        []User
	Allpost        []Post
	Allfollowers   []User
	Allfollowing   []User
	Allnotif       []Notif
	AllUtilisateur []User
	MyGroup        []Group
	UncknowGroup   []Group
	GetAllMembers  []Members
}

type Members struct {
	ID_Member int
	ID_Post   int
	ID_Group  int `json:"ID_Group"`
	ID_User   int `json:"ID_User"`
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
	Id          int     `json:"ID"`
	Names       *string `json:"Names"`
	Content     string  `json:"Content"`
	Title       string  `json:"Title"`
	Image       string  `json:"Image"`
	ImageData   []byte  `json:"ImageData"`
	ID_User     int     `json:"ID_User"`
	Privacy     string  `json:"Privacy"`
	ID_Group    int     `json:"ID_Group"`
	Types       string  `json:"Types"`
	CreatedPost *string `json:"CreatedPost"`
	MembersPost []int   `json:"MembersPost"`
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
	ID_Notif     int `json:"ID_Notif"`
	Messages     string
	Receiver     string
	Types        string `json:"Types"`
	States       string `json:"States"`
	Sender       string
	AvatarSender string
	ID_Group     int `json:"ID_Group"`
}

type Group struct {
	ID_Group         int    `json:"ID_Group"`
	GroupName        string `json:"GroupName"`
	GroupImage       string `json:"GroupImage"`
	GroupImageFile   []byte `json:"GroupIhjumageFile"`
	GroupDescription string `json:"GroupDescription"`
	IdMember         []int  `json:"IdMember"`
	IdCreator        int    `json:"IdCreator"`
}
