package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(10)

	// makes http GET request to random word API
	response, err := http.Get("https://random-word-api.herokuapp.com/word?number=10")
	if err != nil {
		fmt.Printf("The HTTP Get request failed with error %v\n", err)
	}
	defer response.Body.Close()

	// saving the http response in the variable body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	// creating array unmarshall json
	var words []string
	// unmarshalling json
	_ = json.Unmarshal([]byte(body), &words)
	log.Printf("Unmarshalled: %v", words)

	// choosing a random word from the array
	log.Printf(words[randomIndex])

}
