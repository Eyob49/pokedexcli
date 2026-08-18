package main

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
}

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		fmt.Println("no name of Pokemon provided")
		return nil
	}

	url := "https://pokeapi.co/api/v2/pokemon/" + args[0]

	var body []byte

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	if val, ok := cfg.cache.Get(url); ok {
		body = val
	} else {
		res, err := http.Get(url)
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode > 299 {
			return fmt.Errorf("bad response status: %d", res.StatusCode)
		}

		body, err = io.ReadAll(res.Body)
		if err != nil {
			return err
		}

		cfg.cache.Add(url, body)
	}

	var pokemon Pokemon
	if err := json.Unmarshal(body, &pokemon); err != nil {
		return err
	}

	randNum := rand.Intn(pokemon.BaseExperience)
	catchThreshold := 40

	if randNum > catchThreshold {
		fmt.Println(pokemon.Name + " escaped!")
	} else {
		fmt.Println(pokemon.Name + " was caught!")
		cfg.pokedex[pokemon.Name] = pokemon
	}

	return nil
}