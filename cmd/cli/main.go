package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command struct {
	name        string
	description string
	callback    func(*state) error
}

type state struct {
	LocationNextUrl  *string
	LocationPrevtUrl *string
}

func Start() {
	var st state
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		cmd := GetCommand()
		if c, ok := cmd[strings.ToLower(scanner.Text())]; ok {
			c.callback(&st)
		} else {
			fmt.Println("Command not found")
			cmd["help"].callback(&st)
		}
	}
}
