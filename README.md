# go-pokedexcli

To improve my knowledge and dev skills, I subscribe to [boot.dev](https://www.boot.dev/).
I followed the Golang Course and realized the project PokedexCLI.

This repository is the result of this project.

## Use the CLI

Build the project

```shell
make build
```

Build and run the binary

```shell
make run
```

### Command information

```json
  {
   name:        "help",
   description: "Displays a help message",
   callback:    commandHelp,
  },
  "exit": {
   name:        "exit",
   description: "Exit the Pokedex",
   callback:    commandExit,
  },
  "map": {
   name:        "map",
   description: "Get Location area in Pokemon World",
   callback:    commandMap,
  },
  "bmap": {
   name:        "bmap",
   description: "Get previous Location area in Pokemon World",
   callback:    commandBMap,
  },
  "explore": {
   name:        "explore {Location Name}",
   description: "Explore location and found Pokemon",
   callback:    commandExplore,
  },
  "catch": {
   name:        "catch {Pokemon Name}",
   description: "Cacth a new Pokemon",
   callback:    commandCatch,
  },
  "inspect": {
   name:        "inpect {Pokemon Name}",
   description: "Get information about caught pokemon",
   callback:    commandInspect,
  },
  "pokedex": {
   name:        "Pokedex",
   description: "Get all caught pokemon",
   callback:    commandPokedex,
  },
```
