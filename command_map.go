package main

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
)

func commandMap(cfg *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.next != nil {
		url = *cfg.next
	}

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("bad response status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
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


func commandMapb(cfg *config) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *cfg.previous

	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return fmt.Errorf("bad response status: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return err
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