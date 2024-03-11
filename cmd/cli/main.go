package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/q-sw/go-pokedexcli/internal/pokeApi"
	pokecache "github.com/q-sw/go-pokedexcli/internal/pokeCache"
)

func Start() {
	var st state
	st.PokeCache = pokecache.NewCache(time.Duration(time.Second * 10))
	st.PokemonCatch = make(map[string]pokeApi.Pokemon)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		cmd := GetCommand()
		cmdName := strings.Split(scanner.Text(), " ")
		if c, ok := cmd[strings.ToLower(cmdName[0])]; ok {
			c.callback(&st, cmdName...)
		} else {
			fmt.Println("Command not found")
			cmd["help"].callback(&st)
		}
	}
}
