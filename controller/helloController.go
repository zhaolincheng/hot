package controller

import (
	"fmt"
	"hot/utils"
	"net/http"
)
func helloGet(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println("username:", r.Form["username"])
	fmt.Println("password:", r.Form["password"])
	_, err := fmt.Fprintln(w, len(r.Form["username"]))
	if err != nil {
		utils.Error.Println(err)
	}
}

func helloPost(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method)
	fmt.Println(r.FormValue("name"))
	_, err := fmt.Fprintln(w, "hello")
	if err != nil {
		utils.Error.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}