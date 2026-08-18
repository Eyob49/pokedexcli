package main

import (
	"time"
	"github.com/Eyob49/pokedexcli/internal/pokecache"
)

func main() {
    cfg := &config{
		commands: getCommands(),
		cache:    pokecache.NewCache(5 * time.Second),
		pokedex:  make(map[string]Pokemon),
	}
	
	startRepl(cfg)
}
