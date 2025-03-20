# Pokedex CLI

![Pokémon Logo](https://upload.wikimedia.org/wikipedia/commons/thumb/9/98/International_Pok%C3%A9mon_logo.svg/320px-International_Pok%C3%A9mon_logo.svg.png)

## 📖 About

Pokedex CLI is a command-line application that allows you to explore the world of Pokémon right from your terminal. Built with Go, this application interacts with the [PokeAPI](https://pokeapi.co/) to fetch information about different Pokémon, their locations, and more.

## ✨ Features

- Browse Pokémon location areas
- View detailed information about Pokémon
- Catch Pokémon and build your collection
- Interactive command-line interface
- Simple and intuitive commands


### Using Go

```bash
# Clone the repository
git clone https://github.com/twomotive/pokedexCLI.git

# Navigate to project directory
cd pokedexCLI

# Build the application
go build -o pokedex

# Run the application
./pokedex
```

## 🎮 Usage

help - Displays help information about available commands.  

map - Shows a list of location areas.  

explore [LocationArea] - Allows you to explore Pokémon found in the specified location area.  

catch [PokemonName] - Attempts to catch the specified Pokémon.  

inspect [PokemonName] - Displays detailed information about the specified Pokémon.  

pokedex - Lists the Pokémon in your collection.  

exit - Exits the application.




