package main

import (
	"fmt"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

func commandMapB(config *configURL) error{
	fmt.Println()
	fmt.Printf("Current URL: %s\n", config.Previous)
	res, err := http.Get(config.Previous)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}

	type LocationArea struct {
		Name string `json:"name"`
	}
	type APIResponse struct {
		Count    int            `json:"count"`
		Next     *string        `json:"next"`
		Previous *string        `json:"previous"`
		Results  []LocationArea `json:"results"`
	}

	var apiResponse APIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		fmt.Println(err)
	}
    if apiResponse.Next != nil {
		config.Next = *apiResponse.Next
        fmt.Println("Next url:", config.Next)
    } else {
        fmt.Println("Next url: nil")
    }

    if apiResponse.Previous != nil {
		config.Previous = *apiResponse.Previous
        fmt.Println("Previous url:", config.Previous)
		fmt.Println()
    } else {
        fmt.Println("Previous url: nil")
		fmt.Println()
    }

	for _, area := range apiResponse.Results {
		fmt.Println(area.Name)
	}

	return nil
}