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

