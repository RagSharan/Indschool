package main

import(
	"fmt"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "All Users endpoint hits")
}

func NewUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Create New User EndPoint")
}

func DeleteUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Delete User EndPoint")
}
func UpdateUser(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Update User EndPoint")
}