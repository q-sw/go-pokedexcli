package cli

import (
	"errors"
	"fmt"
	"math/rand"
	"os"

	"github.com/q-sw/go-pokedexcli/internal/pokeApi"
)

func GetCommand() map[string]command {
	return map[string]command{
		"help": {
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
	}
}

func commandExit(st *state, args ...string) error {
	fmt.Println("Goodbye !!!")
	os.Exit(0)
	return nil
}

func commandHelp(st *state, args ...string) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex! \nUsage:")

	fmt.Println()

	for _, v := range GetCommand() {
		fmt.Printf("%v: %v \n", v.name, v.description)
	}
	fmt.Println()
	return nil
}

func commandMap(st *state, args ...string) error {
	fmt.Println("Location Area in Pokemon World:")
	fmt.Println()
	loc, err := pokeApi.GetLocation(st.LocationNextUrl, st.PokeCache)

	if err != nil {
		return err
	}

	for l := range loc.Results {
		fmt.Printf("\t- %s\n", loc.Results[l].Name)

	}
	st.LocationNextUrl = loc.Next
	st.LocationPrevtUrl = loc.Previous
	fmt.Println()
	return nil
}

func commandBMap(st *state, args ...string) error {
	fmt.Println("Location Area in Pokemon World:")
	fmt.Println()
	loc, err := pokeApi.GetLocation(st.LocationPrevtUrl, st.PokeCache)

	if err != nil {
		return err
	}

	for l := range loc.Results {
		fmt.Printf("\t- %s\n", loc.Results[l].Name)

	}
	st.LocationNextUrl = loc.Next
	st.LocationPrevtUrl = loc.Previous
	fmt.Println()
	return nil
}
func commandExplore(st *state, args ...string) error {
	fmt.Println("Pokemon in Location:")
	fmt.Println()
	if len(args) == 1 {
		return errors.New("no location provided to explore")

	}
	location := args[1]

	loc, err := pokeApi.GetLocationDetails(location, st.PokeCache)

	if err != nil {
		return errors.New("error during get locationdetails")
	}
	for _, l := range loc.PokemonEncounters {
		fmt.Printf("\t- %s\n", l.Pokemon.Name)

	}
	fmt.Println()
	return nil
}

func commandCatch(st *state, args ...string) error {
	fmt.Println("Catch Pokemon:")
	fmt.Println()
	if len(args) == 1 {
		return errors.New("no pokemon provided to cacth it")

	}
	pokemon := args[1]

	poke, err := pokeApi.GetPokemon(pokemon, st.PokeCache)

	if err != nil {
		return errors.New("error during get pokemon")
	}

	fmt.Printf("%v, %v\n", poke.Name, poke.BaseExperience)
	chance := rand.Intn(poke.BaseExperience)
	fmt.Printf("my chance %v\n", chance)
	if chance < poke.BaseExperience/2 {
		fmt.Printf("no Chance you failed to catch %s\n", pokemon)
		return errors.New("no Chance you failed to catch pokemon")

	}
	fmt.Printf("Congrats you caught %s\n", pokemon)
	st.PokemonCatch[pokemon] = poke
	fmt.Println()

	fmt.Printf("My pokemon caught %v\n", st.PokemonCatch)

	return nil
}

func commandInspect(st *state, args ...string) error {
	if len(args) == 1 {
		return errors.New("no pokemon provided to get information")
	}

	poke, ok := st.PokemonCatch[args[1]]
	if !ok {
		return errors.New("you do not have this pokemon")
	}

	fmt.Println(poke.Name)
	fmt.Println(poke.Height)
	fmt.Println(poke.Weight)
	fmt.Println("Stats:")
	for _, stat := range poke.Stats {
		fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range poke.Types {
		fmt.Printf("-%s", t.Type.Name)
	}
	return nil
}
