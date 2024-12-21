package main

import (
	"fmt"

	"github.com/KrisQ/pokemon-go/internal/pokeapi"
)

func commandMapb(c *Config) error {
	if c.PreviousUrl == "" {
		fmt.Println("you're on the first page")
		return nil
	}
	l, err := pokeapi.GetLocationAreas(c.PreviousUrl)
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
