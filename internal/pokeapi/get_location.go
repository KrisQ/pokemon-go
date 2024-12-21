package pokeapi

// curl get https://pokeapi.co/api/v2/location-area\?limit\=20
// previous null;
// next: result.next https://pokeapi.co/api/v2/location-area?offset=10&limit=10
// previous: https://pokeapi.co/api/v2/location-area?offset=0&limit=10

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

type LocationArea struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreas(url string) (LocationArea, error) {
	var l LocationArea

	res, err := http.Get(url)
	if err != nil {
		return l, err
	}
	if res.StatusCode > 299 {
		return l, errors.New("something went wrong")
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return l, nil
	}
	if err := json.Unmarshal(body, &l); err != nil {
		return l, err
	}

	return l, nil
}
