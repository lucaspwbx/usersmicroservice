package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

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
	db.DropTableIfExists(User{})
	db.CreateTable(User{})
	return db, err
}

func (s *UserService) Run() {
	db, err := s.getDb()
	if err != nil {
		log.Println("Error getting database")
		return
	}

	resource := &UserResource{db: db}

	r := mux.NewRouter()
	r.HandleFunc("/users", resource.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users", resource.CreateUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", resource.GetUserHandler).Methods("GET")
	r.HandleFunc("/users/{id}", resource.DeleteUserHandler).Methods("DELETE")
	r.HandleFunc("/users/{id}", resource.UpdateUserHandler).Methods("PUT")
	http.Handle("/", r)
	http.ListenAndServe(":8080", r)
}

func (ur *UserResource) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
	}

	ur.db.Create(&user)
	user_href := fmt.Sprintf("/users/%d", user.Id)
	w.Header().Set("Location", user_href)
	w.WriteHeader(http.StatusCreated)
}

func (ur *UserResource) GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	ur.db.Where(&User{Id: id}).First(&user)

	data, err := json.Marshal(user)
	if err != nil {
		fmt.Errorf("Error encoding JSON")
	}
	w.Write(data)
}

func (ur *UserResource) GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	//users := []User{User{1, "Lucas", 29}, User{2, "Majoe", 30}}
	var users []User

	ur.db.Find(&users)
	data, err := json.Marshal(users)
	if err != nil {
		fmt.Errorf("Error encoding JSON")
	}
	w.Write(data)
}

func (ur *UserResource) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user User

	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	ur.db.Where(&User{Id: id}).First(&user)
	ur.db.Delete(&user)

	w.WriteHeader(http.StatusNoContent)
}

func (ur *UserResource) UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	var user, newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	ur.db.Where(&User{Id: id}).First(&user)
	ur.db.Model(&user).Updates(newUser)

	w.WriteHeader(http.StatusNoContent)
}
