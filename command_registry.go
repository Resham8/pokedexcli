package main 

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

var commands = map[string]cliCommand{	
	"help" : {
		name: "help",
		description: "Displays a help message",
		callback: commandHelp,
	},
	"map" : {
		name:"map",
		description: "Displays list of areas",
		callback: commandMap,
	},
	"mapb" : {
		name:"mapb",
		description: "Displayes the previous list",
		callback: commandMapback,
	},
	"explore" : {
		name:"explore {location_area}",
		description: "Displayes the list of pokemon's in the location area",
		callback: commandExplore,
	},
	"catch" : {
		name:"catch {pokemon_name}",
		description: "catch's the pokemon",
		callback: commandCatch,
	},
	"inspect":{
		name: "inspect {pokemon_name}",
		description: "Displays details about the pokemon",
		callback: commandInspect,
	},
	"pokedex" :{
		name: "pokedex",
		description: "Displays all the pokemon's caught",
		callback: commandPokedex,
	},
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	},
}
