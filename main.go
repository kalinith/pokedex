package main
import ("strings"
	"fmt"
	"bufio"
	"os"
	)

func cleanInput(text string) []string {
	lowertext := strings.ToLower(text)
	output := strings.Fields(lowertext)
	return output
}

func main() {
	input := bufio.NewScanner(os.Stdin)
	commands := []string{}
	text := ""

	for ;; {
		fmt.Print("Pokedex >")
		input.Scan()
		text = input.Text()
		commands = cleanInput(text)
		if len(commands) > 0 {
			fmt.Printf("Your command was: %s\n", commands[0])
		}
	}
}
