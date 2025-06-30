// package pokeapi

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io"
// 	"net/http"
// )

// // type Client struct {
// // 	Next     string
// // 	Previous string
// // }

// type LocationArea struct {
// 	Name string `json:"name"`
// }

// type APIResponse struct {
// 	Count    int            `json:"count"`
// 	Next     *string        `json:"next"`
// 	Previous *string        `json:"previous"`
// 	Results  []LocationArea `json:"results"`
// }

// func (c *Client) GetLocations(url string) ([]LocationArea, error) {
// 	res, err := http.Get(url)
// 	if err != nil {
// 		return nil, fmt.Errorf("HTTP request failed: %v", err)
// 	}
// 	defer res.Body.Close()

// 	body, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, fmt.Errorf("failed to read response: %v", err)
// 	}

// 	if res.StatusCode > 299 {
// 		return nil, fmt.Errorf("API request failed with status %d: %s", res.StatusCode, body)
// 	}

// 	var apiResponse APIResponse
// 	if err := json.Unmarshal(body, &apiResponse); err != nil {
// 		return nil, fmt.Errorf("JSON parse error: %v", err)
// 	}

// 	// Update client state
// 	if apiResponse.Next != nil {
// 		c.Next = *apiResponse.Next
// 	} else {
// 		c.Next = ""
// 	}

// 	if apiResponse.Previous != nil {
// 		c.Previous = *apiResponse.Previous
// 	} else {
// 		c.Previous = ""
// 	}

// 	return apiResponse.Results, nil
// }