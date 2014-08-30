package service

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user := User{Id: 1, Name: "Lucas", Age: 29}

	data, err := json.Marshal(user)
	assert.Equal(t, err, nil)

	req, err := http.NewRequest("POST", "/users", bytes.NewReader(data))
	res := httptest.NewRecorder()

	CreateUserHandler(res, req)

	assert.Equal(t, res.Code, 201)
	assert.Equal(t, res.HeaderMap["Location"][0], "/users/1")
}

func TestGetUser(t *testing.T) {
	req, err := http.NewRequest("GET", "/users/1", nil)
	assert.Equal(t, err, nil)

	res := httptest.NewRecorder()
	GetUserHandler(res, req)

	assert.Equal(t, res.Code, 200)

	var got User
	err = json.NewDecoder(res.Body).Decode(&got)
	assert.Equal(t, err, nil)
	assert.Equal(t, got.Id, 1)
	assert.Equal(t, got.Name, "Lucas")
	assert.Equal(t, got.Age, 29)
}

func TestDeleteUser(t *testing.T) {
	req, err := http.NewRequest("DELETE", "/users/1", nil)
	assert.Equal(t, err, nil)

	res := httptest.NewRecorder()
	m := mux.NewRouter()
	m.HandleFunc("/users/{id}", DeleteUserHandler)
	m.ServeHTTP(res, req)

	assert.Equal(t, res.Code, 204)
}

func TestUpdateUser(t *testing.T) {
	req, err := http.NewRequest("PUT", "/users/1", nil)
	assert.Equal(t, err, nil)

	res := httptest.NewRecorder()
	m := mux.NewRouter()
	m.HandleFunc("/users/{id}", UpdateUserHandler)
	m.ServeHTTP(res, req)

	assert.Equal(t, res.Code, 204)
}
