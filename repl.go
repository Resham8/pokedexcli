package main

import (
	"bufio"
	"fmt"	
	"os"	
)

func startREPL(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			line := scanner.Text()

			command, exists := commands[line]

			if exists {
				err := command.callback(cfg)

				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}

		}
	}
}