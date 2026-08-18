# Pokedex CLI

A command-line Pokedex built in Go that lets you explore Pokemon locations, catch Pokemon, and build your collection using real-time data from the [PokeAPI](https://pokeapi.co/).

## Features

- **Location Exploration**: Browse through location areas with pagination support (`map` and `mapb` commands)
- **Pokemon Discovery**: Explore specific locations to see which Pokemon are found there
- **Catch Mechanics**: Attempt to catch Pokemon with probability-based difficulty (stronger Pokemon are harder to catch)
- **Pokedex Management**: Build and view your collection of caught Pokemon
- **Inspect Details**: View detailed stats, types, height, and weight of your caught Pokemon
- **Smart Caching**: Reduces API calls with concurrent-safe, time-expiring cache layer
- **Persistent State**: Your caught Pokemon list persists for the duration of your session

## Installation

### Prerequisites
- Go 1.20 or higher

### Build

```bash
git clone https://github.com/yourusername/pokedexcli.git
cd pokedexcli
go build
```

### Run

```bash
./pokedexcli
```

Or run directly:
```bash
go run .
```

## Usage

Once the REPL starts, you can use the following commands:

### Core Commands

- **`help`** — Display available commands and usage information
- **`map`** — Show the next 20 location areas (paginated)
- **`mapb`** — Show the previous 20 location areas
- **`explore <location-name>`** — List all Pokemon found in a location area
  ```
  Pokedex > explore pastoria-city-area
  ```
- **`catch <pokemon-name>`** — Attempt to catch a Pokemon (success rate based on base experience)
  ```
  Pokedex > catch pikachu
  Throwing a Pokeball at pikachu...
  pikachu was caught!
  ```
- **`pokedex`** — Display all Pokemon you've caught
- **`inspect <pokemon-name>`** — View detailed information about a caught Pokemon
  ```
  Pokedex > inspect pikachu
  Name: pikachu
  Height: 4
  Weight: 60
  Stats:
    -hp: 35
    -attack: 55
    -defense: 40
    ...
  Types:
    - electric
  ```
- **`exit`** — Close the Pokedex

## Technical Highlights

### Architecture & Patterns

- **Command Registry Pattern**: Flexible dispatch system for extensible command handling
- **Shared Config State**: Central `config` struct passed through the application for maintaining session state
- **Separation of Concerns**: Organized into logical packages and files (REPL logic, commands, API interaction, caching)

### Concurrency & Performance

- **Concurrent-Safe Caching**: `sync.Mutex`-protected cache with background goroutine for automatic entry expiration
- **Goroutine-Based Cleanup**: Time-based cache eviction running in the background without blocking user interaction
- **Smart API Reuse**: Caches all API responses, so exploring the same location twice is instantaneous

### API Integration

- **Real-Time Data**: Integrates with the free PokeAPI for live Pokemon and location data
- **Error Handling**: Graceful handling of network errors, rate limits, and malformed responses
- **Pagination Support**: Maintains state across multiple requests to handle large result sets

### Go Fundamentals Demonstrated

- JSON unmarshaling with nested structs
- Variadic function parameters for flexible command arguments
- Map manipulation and safe concurrent access
- HTTP client requests with proper error handling
- Time-based operations and goroutines
- Pointer receivers and method organization

## Project Structure

```
pokedexcli/
├── main.go                 # Entry point, REPL initialization
├── repl.go                 # REPL loop, command registry, core structs
├── command_map.go          # map/mapb commands for location exploration
├── command_explore.go      # explore command for location details
├── command_catch.go        # catch command with probability mechanics
├── command_inspect.go      # inspect command for Pokemon details
├── command_pokedex.go      # pokedex command for viewing collection
├── internal/pokecache/     # Internal caching package
│   ├── pokecache.go        # Cache implementation
│   └── pokecache_test.go   # Cache tests
└── go.mod                  # Go module definition
```

## What I Learned

This project reinforced fundamental backend engineering concepts through practical application:

1. **API Client Design**: Building reliable HTTP clients with proper error handling and timeout management
2. **Caching Strategies**: Implementing time-expiring caches with concurrent safety
3. **State Management**: Designing shared state patterns for application data
4. **Goroutines & Concurrency**: Understanding Go's lightweight concurrency model for background tasks
5. **Test-Driven Development**: Writing tests before implementation for better code design
6. **Clean Architecture**: Organizing code for readability and maintainability as features grow
7. **Go Idioms**: Learning Go conventions (comma-ok patterns, defer, variadic parameters, etc.)

## Testing

Run all tests:
```bash
go test ./...
```

The project includes unit tests for the cache package, verifying both basic add/get functionality and the time-based reaping logic.

## Future Enhancements

- Persistent storage (save caught Pokemon to disk between sessions)
- Battle mechanics between caught Pokemon
- Trade system between players
- More advanced filtering in explore (by type, stats, etc.)
- Move history and learning mechanics

## License

This project is provided as-is for educational purposes.

## Acknowledgments

- [PokeAPI](https://pokeapi.co/) for providing free Pokemon data
- [Boot.dev](https://boot.dev/) for the project guidance and structure
