package main

import (
	"bufio"
	"fmt"	
	"os"	
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")

		if scanner.Scan() {
			line := scanner.Text()

			command, exists := commands[line]

			if exists {
				err := command.callback()

				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("Unknown command")
			}

		}
	}
}