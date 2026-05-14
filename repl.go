package main

import (
	"fmt"
	"github.com/chzyer/readline"
	"strings"
)

func startREPL(cfg *config) {

	rl, err := readline.New("Pokedex > ")
	if err != nil {
		panic(err)
	}
	defer rl.Close()

	for {
		line, err := rl.Readline()

		if err != nil {
			break
		}

		line = strings.TrimSpace(line)

		if line == "" {
			continue
		}

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
