package remote

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
)

type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func FetchLocationArea(offset int) ([]string, error) {
	locationAreaResponse := LocationAreaResponse{}
	url := "https://pokeapi.co/api/v2/location-area?limit=20offset=" + strconv.Itoa(offset)
	res, err := http.Get(url)
	if err != nil {
		return []string{}, err
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return []string{}, err
	}
	err = json.Unmarshal(body, &locationAreaResponse)
	if err != nil {
		return []string{}, err
	}
	var output []string
	for _, result := range locationAreaResponse.Results {
		output = append(output, result.Name)
	}
	return output, nil
}
