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
	location	string
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

func main() {
	// Step 1: Declare an empty map
	m := make(map[string]cliCommand)
	//declare configs
	exitconf := &config{}
	helpconf := &config{}
	locconf  := &config{
				prev: "",
				next: "https://pokeapi.co/api/v2/location-area"}
	visconf := &config{
				prev: "https://pokeapi.co/api/v2/location-area",
				next: ""}
	
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
		config:		 locconf,
		callback:    GetLocationData(visconf, cache),
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
				visconf.next = commands[1]
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
		}
	}
}
