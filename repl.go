package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			line := scanner.Text()

			words := strings.Fields(line)
			commandName := words[0]

			args := words[1:]
			command, exists := commands[commandName]

			if exists {
				err := command.callback(cfg, args)

				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}

		}
	}
}
