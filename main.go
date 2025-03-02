package main
import ("strings"
	"fmt"
	"bufio"
	"os"
	"time"
	"github.com/kalinith/pokedex/internal"
	)

type config struct {
	prev		string
	next		string
	param		string
	pokedex		map[string]Pokemon
}

type cliCommand struct {
	name        string
	description string
	config		*config //pointer
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func makeHelpCommand(m map[string]cliCommand) func() error {
    return func() error {
        fmt.Println("Welcome to the Pokedex!")
        fmt.Println("Usage:")
        for _, cmd := range m {
            fmt.Printf("%s: %s\n", cmd.name, cmd.description)
        }
        return nil
    }
}

func cleanInput(text string) []string {
	lowertext := strings.ToLower(text)
	output := strings.Fields(lowertext)
	return output
}

func printPokedex(config *config) func() error {
    return func() error {
    	fmt.Println("Your Pokedex:")
    	for _, pokemon := range config.pokedex {
    		fmt.Printf(" - %s\n", pokemon.Name)
    	}
    	return nil
    }
	
}

func main() {
	// Step 1: Declare an empty map
	m := make(map[string]cliCommand)
	pokedex := make(map[string]Pokemon)
	//declare configs
	exitconf := &config{}
	helpconf := &config{}
	locconf  := &config{
				prev: "",
				next: "https://pokeapi.co/api/v2/location-area",
			}
	visconf := &config{
				param: "",
				pokedex: pokedex,
			}

	cache := internal.NewCache(15 * time.Second)

	// Step 2: Now add commands to the map, using the completed map
	m["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		config:		 exitconf,
		callback:    commandExit,
	}
	m["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		config:		 helpconf,
		callback:    makeHelpCommand(m), // Pass the map here
	}
	m["map"]  = cliCommand{
		name:        "map",
		description: "Display the next page of locations from the world",
		config:		 locconf,
		callback:    makeGetLocationArea(locconf, cache),
	}
	m["mapb"]  = cliCommand{
		name:        "mapb",
		description: "Display previous page of locations from the world",
		config:		 locconf,
		callback:    GetPrevLocationArea(locconf, cache),
	}
	m["explore"] = cliCommand{
		name:		 "explore {location}",
		description: "Explore the given location",
		config:		 visconf,
		callback:    GetLocationData(visconf, cache),
	}
	m["catch"] = cliCommand{
		name:		 "catch {pokemon}",
		description: "attempt to catch a pokemon",
		config:		 visconf,
		callback:    GetCatchPokemon(visconf, cache),
	}
	m["inspect"] = cliCommand{
		name:		 "inspect {pokemon}",
		description: "inspect the details of a caught pokemon",
		config:		 visconf,
		callback:	 inspectPokemon(visconf),
	}
	m["pokedex"] = cliCommand{
		name:		 "pokedex",
		description: "print a list of the names of pokemon you have caught",
		callback:	 printPokedex(visconf),
	}
	
	input := bufio.NewScanner(os.Stdin)
	commands := []string{}
	text := ""
	command := ""

	for ;; {
		fmt.Print("Pokedex >")
		input.Scan()
		text = input.Text()
		commands = cleanInput(text)
		if len(commands) > 0 {
			command = commands[0]
			if len(commands) > 1 {
				visconf.param = commands[1]
			}
			_, exists := m[command]
			if exists {
				err := m[command].callback()
				if err != nil {
					fmt.Println(err)
				}

			} else {
				fmt.Println("Unknown command")
			}
		visconf.param = ""
		}
	}
}
