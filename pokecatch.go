package main
import (
	"math/rand"
	"fmt"
	"time"
)



func catchPokemon(pokemon Pokemon, pokedex map[string]Pokemon) {
	t := rand.NewSource(time.Now().UTC().UnixNano())
	random := rand.New(t)

	rolled := random.Intn(100)

    catchDifficulty := pokemon.BaseExperience / 3
    
    // Ensure difficulty is between 10-90 to always give some chance
    if catchDifficulty < 10 {
        catchDifficulty = 10 // Minimum 10% failure rate
    } else if catchDifficulty > 90 {
        catchDifficulty = 90 // Minimum 10% success rate
    }

	if rolled < catchDifficulty {
		fmt.Printf("you failed to catch a %s, your values was %d and the target was above %d\n", pokemon.Name, rolled, catchDifficulty)
		return
	}
	pokedex[pokemon.Name] = pokemon
	fmt.Printf("you caught a %s by rolling a %d when you needed to get above %d\n", pokemon.Name, rolled, catchDifficulty)
	return
}

func inspectPokemon(config *config) func() error {
	return func () error {
		pokemon := config.param
		if pokemon == "" {
			fmt.Println("No pokemon entered")
			return nil
		}
	
		data, exists := config.pokedex[pokemon]
		if !exists {
			//fmt.Printf("You haven't caught a %s yet\n", pokemon)
			fmt.Println("you have not caught that pokemon")
			return nil
		}
		fmt.Printf("Name: %s\nHeight: %d\nWeight: %d\nStats:\n",data.Name, data.Height, data.Weight)
		//, stats and type(s)
		for _, stat := range data.Stats {
			fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, element := range data.Types {
			fmt.Printf("  - %s\n", element.Type.Name)
		}
		return nil
	}
}