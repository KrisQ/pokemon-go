package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

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
