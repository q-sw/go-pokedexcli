package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func StartPrompt() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		cmd := commands()

		if c, ok := cmd[strings.ToLower(scanner.Text())]; ok {
			c.callback()
		} else {
			fmt.Println("Command not found")
			cmd["help"].callback()
		}
	}
}
