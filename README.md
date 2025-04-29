# Pokedex CLI

![Pokémon Logo](https://upload.wikimedia.org/wikipedia/commons/thumb/9/98/International_Pok%C3%A9mon_logo.svg/320px-International_Pok%C3%A9mon_logo.svg.png)

## 📖 About

Pokedex CLI is a command-line application that allows you to explore the world of Pokémon right from your terminal. Built with Go, this application interacts with the [PokeAPI](https://pokeapi.co/) to fetch information about different Pokémon, their locations, and more.

**This project was developed as a learning exercise to practice and demonstrate proficiency in Go development.**

## 💡 Learning Objectives & Skills Demonstrated

Through building this project, I focused on learning and applying the following concepts in Go:

*   **Command-Line Interface (CLI) Development:** Creating an interactive REPL (Read-Eval-Print Loop) using standard Go libraries (`bufio`, `os`, `strings`).
*   **API Integration:** Consuming a third-party REST API (PokeAPI) using Go's `net/http` package.
*   **Data Handling:** Marshalling and unmarshalling JSON data.
*   **Concurrency:** (If applicable, add details here if concurrency was used, e.g., for API calls).
*   **Caching:** Implementing a simple in-memory cache (`pokecache`) to reduce API calls and improve performance.
*   **Project Structure:** Organizing code into logical packages (`internal/api`, `internal/cli`, `internal/config`, etc.).
*   **Dependency Management:** Using Go modules (`go.mod`).
*   **Error Handling:** Implementing robust error checking and reporting.

## ✨ Features

*   Browse Pokémon location areas
*   View detailed information about Pokémon
*   Catch Pokémon and build your collection
*   Interactive command-line interface
*   Simple and intuitive commands

## 🚀 Getting Started

### Prerequisites

*   Go (version 1.x or later)

### Installation & Running

```bash
# Clone the repository
git clone https://github.com/twomotive/pokedex.git

# Navigate to project directory
cd pokedex

# Build the application
go build -o pokedex ./cmd/pokedex

# Run the application
./pokedex
```

## 🎮 Usage

*   `help`: Displays help information about available commands.
*   `map`: Shows the next page of location areas.
*   `mapb`: Shows the previous page of location areas.
*   `explore <LocationAreaName>`: Lists Pokémon found in the specified location area.
*   `catch <PokemonName>`: Attempts to catch the specified Pokémon.
*   `inspect <PokemonName>`: Displays detailed information about a caught Pokémon.
*   `pokedex`: Lists all Pokémon caught in your Pokedex.
*   `exit`: Exits the application.




