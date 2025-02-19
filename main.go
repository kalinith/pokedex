package main
import ("strings"
	"fmt"
	"bufio"
	"os"
	)

type cliCommand struct {
	name        string
	description string
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

	// Step 2: Now add commands to the map, using the completed map
	m["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	m["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    makeHelpCommand(m), // Pass the map here
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
			_, exists := m[command]
			if exists {
				m[command].callback()

			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
