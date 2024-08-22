package main

import (
	"fmt"
	"encoding/json"
	"net/http"
)

type Joke struct {
	Joke string `json:"joke"`
}

func getJoke() (*Joke, error) {
	// Create a new request
	req, err := http.NewRequest("GET", "https://icanhazdadjoke.com/", nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	// Set the Accept header to application/json
	req.Header.Set("Accept", "application/json")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch joke: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %v", resp.Status)
	}

	// Decode the JSON response
	var joke Joke
	err = json.NewDecoder(resp.Body).Decode(&joke)
	if err != nil {
		return nil, fmt.Errorf("failed to decode joke: %v", err)
	}

	return &joke, nil
}

func main() {
	joke, err := getJoke()
	if err != nil {
		fmt.Println("error occurred somewhere, idk")
		return
	}

	fmt.Println(joke.Joke)
}


