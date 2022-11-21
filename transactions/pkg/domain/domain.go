package domain

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func catchError(text string) {
	if r := recover(); r != nil {
		message := fmt.Sprintf("%s is not started yet", text)
		fmt.Println(message)
	}
}

func FetchUserID(id int) (map[string]interface{}, error) {
	defer catchError("user server")

	userServer := "http://localhost:8080"
	var client = &http.Client{}
	var user map[string]interface{}
	endPoint := fmt.Sprintf("%s/user/%d/", userServer, id)
	req, err := http.NewRequest("GET", endPoint, nil)

	if err != nil {
		panic(err.Error())
	}

	response, err := client.Do(req)

	if err != nil {
		panic(err.Error())
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&user)

	return user, err
}

func FetchBookID(id int) (map[string]interface{}, error) {
	defer catchError("book server")

	userServer := "http://localhost:8090"
	var client = &http.Client{}
	var user map[string]interface{}
	endPoint := fmt.Sprintf("%s/book/%d/", userServer, id)
	req, err := http.NewRequest("GET", endPoint, nil)

	if err != nil {
		panic(err.Error())
	}

	response, err := client.Do(req)

	if err != nil {
		panic(err.Error())
	}

	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&user)

	return user, err
}
