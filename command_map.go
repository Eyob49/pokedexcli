package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func commandMap(cfg *config, args ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != nil {
		url = *cfg.next
	}

	var body []byte

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

	var locationResp LocationAreaResponse
	if err := json.Unmarshal(body, &locationResp); err != nil {
		return err
	}

	for _, area := range locationResp.Results {
		fmt.Println(area.Name)
	}

	cfg.next = locationResp.Next
	cfg.previous = locationResp.Previous

	return nil
}


func commandMapb(cfg *config, args ...string) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *cfg.previous

	var body []byte

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

	var locationResp LocationAreaResponse
	if err := json.Unmarshal(body, &locationResp); err != nil {
		return err
	}

	for _, area := range locationResp.Results {
		fmt.Println(area.Name)
	}

	cfg.next = locationResp.Next
	cfg.previous = locationResp.Previous

	return nil
}