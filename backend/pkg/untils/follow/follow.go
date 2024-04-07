package follow

import (
	"encoding/json"
	"fmt"
	"main/pkg/db/sqlite"
	"main/pkg/session"
	Struct "main/pkg/struct"
	"net/http"
)

var FollowHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var myaccount = session.Myaccount(w, r)
	var following Struct.User
	var err = json.NewDecoder(r.Body).Decode(&following)
	if err != nil {
		fmt.Println("Follow : erreur de decodage json", err)
		return
	}
	sqlite.SetFollowing(myaccount, following, "true", "person", 0)
	json.NewEncoder(w).Encode(sqlite.GetMyFollowing(myaccount))
})
