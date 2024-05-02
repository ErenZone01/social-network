package comment

import (
	"main/pkg/session"
	"net/http"
)

var Comment = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	var user = session.Myaccount(w, r)
	//si le compte n'existe pas quitte
	if user.Error {
		return
	}
	
})
