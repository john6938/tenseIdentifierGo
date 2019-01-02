package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestServer(t *testing.T) {
	go main()
	// Wait for server to load
	time.Sleep(time.Second * 1)

	// Sent example sentence to api
	exampleSentence := "I'm a robot"
	req, err := http.Get("http://localhost:8080/api/sentence/tense/" + exampleSentence)
	if err != nil {
		t.Fatal(err.Error())
	}

	// Read all bytes from body
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		t.Fatal(err.Error())
	}
	var responseData struct {
		TenseID int `json:"tenseID"`
	}
	responseData.TenseID = -1
	// Deserialize it
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		t.Fatal(err)
	}

	// Check if tenseID was in response
	if responseData.TenseID == -1 {
		t.Fatal("Response does not contain sentence ID")
	}
}
