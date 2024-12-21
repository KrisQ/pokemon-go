package main

import (
	"fmt"

	"github.com/KrisQ/pokemon-go/internal/pokeapi"
)

func commandMap(c *Config) error {
	url := "https://pokeapi.co/api/v2/location-area?limit=20"
	if c.NextUrl != "" {
		url = c.NextUrl
	}
	l, err := pokeapi.GetLocationAreas(url)
	if err != nil {
		return err
	}
	c.NextUrl = l.Next
	c.PreviousUrl = l.Previous
	for _, area := range l.Results {
		fmt.Println(area.Name)
	}
	return nil
}
