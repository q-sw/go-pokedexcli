package cli

import (
	"errors"
	"fmt"
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
