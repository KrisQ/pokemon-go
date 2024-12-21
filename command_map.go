package main

import (
	"fmt"

	"github.com/KrisQ/pokemon-go/internal/pokeapi"
)

func commandMap(c *Config) error {
	l, err := pokeapi.GetLocationAreas("https://pokeapi.co/api/v2/location-area?limit=20")
	if err != nil {
		return err
	}
	for _, area := range l.Results {
		fmt.Println(area.Name)
	}
	return nil
}
