package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"mvc/services"
	"net/http"
	"strconv"
)

func GetMessage() string {

	fmt.Println("1. performing HTTP GET...")
	resp, err := http.Get("http://constantv9:8081")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()
	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println("API Response as String:\n" + bodyString)

	return bodyString
}

func GetUser(resp http.ResponseWriter, req *http.Request) {

	body := GetMessage()

	userId, err := strconv.ParseInt(req.URL.Query().Get("user_id"), 10, 64)
	fmt.Fprintf(resp, body)
	log.Printf("about to process user_id %v", userId)

	if err != nil {
		// just return the BadRequest to the client
		return
	}

	user, err := services.GetUser(userId)
	if err != nil {
		// handle the error and return to the client
		return
	}

	// return user to client
	jsonValue, _ := json.Marshal(user)
	resp.Write(jsonValue)
}
