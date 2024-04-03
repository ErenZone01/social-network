package home

import (
	"fmt"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	var template, err = template.ParseFiles("pkg/untils/home/index.html")
	if err != nil {
		fmt.Println("error of template")
		return 
	}
	template.Execute(w,nil)
}
