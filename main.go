package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

func printState(RandomWord string, Guesses map[rune]bool) {
	for _, ch := range RandomWord {
		if Guesses[ch] == true {
			fmt.Printf("%c", ch)
		} else {
			fmt.Print("_")
		}
	}
	fmt.Println(" ")
}

func printHangman(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d.txt", hangmanState))
	if err != nil {
		panic(err)
	}
	return string(data)
}

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

	// choosing a random word from the array and storing all as lowercase
	randomWord := words[randomIndex]
	randomWord = strings.ToLower(randomWord)
	log.Printf(randomWord)

	// guessed letters
	guessedLetters := map[rune]bool{}
	// first letter
	guessedLetters[rune(randomWord[0])] = true
	// last letter
	guessedLetters[rune(randomWord[len(randomWord)-1])] = true

	// hangmanState := 0

	printState(randomWord, guessedLetters)

	fmt.Println(printHangman(5))
}
