package controller

import (
	"encoding/json"
	"net/http"

	"../entity"
	"../services"
)

type controller struct{}

var (
	//	userService services.UserService = services.NewUserService()  =>use dependency injectin
	userService services.UserService
)

type UserController interface {
	CreateUser(response http.ResponseWriter, request *http.Request)
	GetUser(response http.ResponseWriter, request *http.Request)
}

func NewUserController(service services.UserService) UserController {
	userService = service
	return &controller{}
}

// create user
func (*controller) CreateUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var user entity.User
	err := json.NewDecoder(request.Body).Decode(&user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}

	user, err = userService.CreateUser(user)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "User is not registered"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)

}

//get user
func (*controller) GetUser(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var userId string
	err := json.NewDecoder(request.Body).Decode(&userId)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	user, err := userService.FindUser(userId)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{"error": "User is not found"}`))
		return
	}
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(user)
}

//update user

// delete user
