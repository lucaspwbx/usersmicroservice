package service

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type UserService struct {
}

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Errorf("Error decoding JSON")
	}

	user_href := fmt.Sprintf("/users/%d", user.Id)
	w.Header().Set("Location", user_href)
	w.WriteHeader(http.StatusCreated)
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	user := User{1, "Lucas", 29}

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Errorf("Error encoding JSON")
	}
	w.Write(data)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	users := []User{User{1, "Lucas", 29}, User{2, "Majoe", 30}}

	data, err := json.Marshal(users)
	if err != nil {
		fmt.Errorf("Error encoding JSON")
	}
	w.Write(data)
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//id := params["id"]
	//TODO
	//delete from database
	w.WriteHeader(http.StatusNoContent)
}

func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	//params := mux.Vars(r)
	//id := params["id"]

	w.WriteHeader(http.StatusNoContent)
}
