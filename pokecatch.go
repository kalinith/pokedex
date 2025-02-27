package main
import (
	"math/rand"
	"fmt"
	"time"
)



func catchPokemon(pokemon Pokemon) {
	random := rand.New(time.Now())
	perCatch := random.intn(100)
	perToCach := pokemon.BaseExperience - 50
	if perCatch >= perToCach {
		fmt.Printf("you caught a %s", pokemon.name)
	}
	return
}

