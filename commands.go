package main

import (
	"fmt"
	"os"
)

func commands() map[string]cliCommand {
	return map[string]cliCommand{
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
	}
}

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex! \nUsage:")

	fmt.Println()

	for _, v := range commands() {
		fmt.Printf("%v: %v \n", v.name, v.description)
	}
	fmt.Println()
	return nil
}

func commandExit() error {
	fmt.Println("Goodbye !!!")
	os.Exit(0)
	return nil
}
