package main

import (
	"fmt"
	"net/http"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

// func handleRequests() {
// 	myRouter := mux.NewRouter().StrictSlash(true)
// 	myRouter.HandleFunc("/", HelloWorld).Methods("GET")
// 	myRouter.HandleFunc("/users", AllUsers).Methods("GET")
// 	myRouter.HandleFunc("/user/{name}/{email}", NewUser).Methods("Post")
// 	myRouter.HandleFunc("/user/{name}", DeleteUser).Methods("DELETE")
// 	myRouter.HandleFunc("/user/{name}/{email}", UpdateUser).Methods("PUT")
// 	log.Fatal(http.ListenAndServe(":8081", myRouter))
// }

// func main(){
// 	fmt.Println("Go Rest Api Tutorial")

// 	handleRequests()
// }
