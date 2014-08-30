package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id   int
	Name string
	Age  int
}

type UserService struct {
}

type UserResource struct {
	db gorm.DB
}

func (s *UserService) getDb() (gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "/tmp/users.db")
	if err != nil {
		return db, errors.New("Error opening DB")
	}
	return db, err
}

func (s *UserService) Run() {
	//db, err := s.getDb()
	//if err != nil {
	//log.Println("Error getting database")
	//return
	//}

	//resource := &UserResource{db: db}

	r := mux.NewRouter()
	r.HandleFunc("/users", GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users/{id}", UpdateUserHandler).Methods("PUT")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
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
