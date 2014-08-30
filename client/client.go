package client

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/lucasweiblen/usersmicroservice/models"
)

type UserClient struct {
	Host string
}

func (uc *UserClient) CreateUser(name string, age int) error {
	user := models.User{Name: name, Age: age}
	url := fmt.Sprintf("%s/users", uc.Host)

	b, err := json.Marshal(user)
	if err != nil {
		log.Print("Error marshaling JSON")
		return errors.New("marshaling Json")
	}
	body := bytes.NewBuffer(b)

	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		log.Print("Error creating request")
		return errors.New("creating request")
	}
	req.Header.Set("content-type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Print("Error making request")
		return errors.New("making request")
	}
	return nil
}

func (uc *UserClient) GetUser(id int) (*models.User, error) {
	url := fmt.Sprintf("%s/users/%d", uc.Host, id)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Print("Error creating request")
		return nil, errors.New("creating request")
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("Error making request")
		return nil, errors.New("making request")
	}
	var user models.User
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Printf("Error reading response")
		return nil, errors.New("reading response")
	}
	err = json.Unmarshal(body, user)
	if err != nil {
		log.Printf("Error unmarshaling JSON")
		return nil, errors.New("unmarshaling json")
	}
	return user, nil
}
