package main
import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
	"github.com/kalinith/pokedex/internal"
)

type LocationPage struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	Abilities []struct {
		Ability struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"ability"`
		IsHidden bool `json:"is_hidden"`
		Slot     int  `json:"slot"`
	} `json:"abilities"`
	BaseExperience int `json:"base_experience"`
	Cries          struct {
		Latest string `json:"latest"`
		Legacy string `json:"legacy"`
	} `json:"cries"`
	Forms []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"forms"`
	GameIndices []struct {
		GameIndex int `json:"game_index"`
		Version   struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"version"`
	} `json:"game_indices"`
	Height    int `json:"height"`
	HeldItems []struct {
		Item struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"item"`
		VersionDetails []struct {
			Rarity  int `json:"rarity"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"held_items"`
	ID                     int    `json:"id"`
	IsDefault              bool   `json:"is_default"`
	LocationAreaEncounters string `json:"location_area_encounters"`
	Moves                  []struct {
		Move struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"move"`
		VersionGroupDetails []struct {
			LevelLearnedAt  int `json:"level_learned_at"`
			MoveLearnMethod struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"move_learn_method"`
			VersionGroup struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version_group"`
		} `json:"version_group_details"`
	} `json:"moves"`
	Name          string        `json:"name"`
	Order         int           `json:"order"`
	PastAbilities []interface{} `json:"past_abilities"`
	PastTypes     []interface{} `json:"past_types"`
	Species       struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"species"`
	Sprites struct {
		BackDefault      string `json:"back_default"`
		BackFemale       string `json:"back_female"`
		BackShiny        string `json:"back_shiny"`
		BackShinyFemale  string `json:"back_shiny_female"`
		FrontDefault     string `json:"front_default"`
		FrontFemale      string `json:"front_female"`
		FrontShiny       string `json:"front_shiny"`
		FrontShinyFemale string `json:"front_shiny_female"`
		Other            struct {
			DreamWorld struct {
				FrontDefault string      `json:"front_default"`
				FrontFemale  interface{} `json:"front_female"`
			} `json:"dream_world"`
			Home struct {
				FrontDefault     string `json:"front_default"`
				FrontFemale      string `json:"front_female"`
				FrontShiny       string `json:"front_shiny"`
				FrontShinyFemale string `json:"front_shiny_female"`
			} `json:"home"`
			OfficialArtwork struct {
				FrontDefault string `json:"front_default"`
				FrontShiny   string `json:"front_shiny"`
			} `json:"official-artwork"`
			Showdown struct {
				BackDefault      string      `json:"back_default"`
				BackFemale       string      `json:"back_female"`
				BackShiny        string      `json:"back_shiny"`
				BackShinyFemale  interface{} `json:"back_shiny_female"`
				FrontDefault     string      `json:"front_default"`
				FrontFemale      string      `json:"front_female"`
				FrontShiny       string      `json:"front_shiny"`
				FrontShinyFemale string      `json:"front_shiny_female"`
			} `json:"showdown"`
		} `json:"other"`
		Versions struct {
			GenerationI struct {
				RedBlue struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"red-blue"`
				Yellow struct {
					BackDefault      string `json:"back_default"`
					BackGray         string `json:"back_gray"`
					BackTransparent  string `json:"back_transparent"`
					FrontDefault     string `json:"front_default"`
					FrontGray        string `json:"front_gray"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"yellow"`
			} `json:"generation-i"`
			GenerationIi struct {
				Crystal struct {
					BackDefault           string `json:"back_default"`
					BackShiny             string `json:"back_shiny"`
					BackShinyTransparent  string `json:"back_shiny_transparent"`
					BackTransparent       string `json:"back_transparent"`
					FrontDefault          string `json:"front_default"`
					FrontShiny            string `json:"front_shiny"`
					FrontShinyTransparent string `json:"front_shiny_transparent"`
					FrontTransparent      string `json:"front_transparent"`
				} `json:"crystal"`
				Gold struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"gold"`
				Silver struct {
					BackDefault      string `json:"back_default"`
					BackShiny        string `json:"back_shiny"`
					FrontDefault     string `json:"front_default"`
					FrontShiny       string `json:"front_shiny"`
					FrontTransparent string `json:"front_transparent"`
				} `json:"silver"`
			} `json:"generation-ii"`
			GenerationIii struct {
				Emerald struct {
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"emerald"`
				FireredLeafgreen struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"firered-leafgreen"`
				RubySapphire struct {
					BackDefault  string `json:"back_default"`
					BackShiny    string `json:"back_shiny"`
					FrontDefault string `json:"front_default"`
					FrontShiny   string `json:"front_shiny"`
				} `json:"ruby-sapphire"`
			} `json:"generation-iii"`
			GenerationIv struct {
				DiamondPearl struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"diamond-pearl"`
				HeartgoldSoulsilver struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"heartgold-soulsilver"`
				Platinum struct {
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"platinum"`
			} `json:"generation-iv"`
			GenerationV struct {
				BlackWhite struct {
					Animated struct {
						BackDefault      string `json:"back_default"`
						BackFemale       string `json:"back_female"`
						BackShiny        string `json:"back_shiny"`
						BackShinyFemale  string `json:"back_shiny_female"`
						FrontDefault     string `json:"front_default"`
						FrontFemale      string `json:"front_female"`
						FrontShiny       string `json:"front_shiny"`
						FrontShinyFemale string `json:"front_shiny_female"`
					} `json:"animated"`
					BackDefault      string `json:"back_default"`
					BackFemale       string `json:"back_female"`
					BackShiny        string `json:"back_shiny"`
					BackShinyFemale  string `json:"back_shiny_female"`
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"black-white"`
			} `json:"generation-v"`
			GenerationVi struct {
				OmegarubyAlphasapphire struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"omegaruby-alphasapphire"`
				XY struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"x-y"`
			} `json:"generation-vi"`
			GenerationVii struct {
				Icons struct {
					FrontDefault string      `json:"front_default"`
					FrontFemale  interface{} `json:"front_female"`
				} `json:"icons"`
				UltraSunUltraMoon struct {
					FrontDefault     string `json:"front_default"`
					FrontFemale      string `json:"front_female"`
					FrontShiny       string `json:"front_shiny"`
					FrontShinyFemale string `json:"front_shiny_female"`
				} `json:"ultra-sun-ultra-moon"`
			} `json:"generation-vii"`
			GenerationViii struct {
				Icons struct {
					FrontDefault string `json:"front_default"`
					FrontFemale  string `json:"front_female"`
				} `json:"icons"`
			} `json:"generation-viii"`
		} `json:"versions"`
	} `json:"sprites"`
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Weight int `json:"weight"`
}

func makeGetLocationArea(conf *config, cache *internal.Cache) func() error {
	return func() error {
		if conf.next == "" {
			return fmt.Errorf("you're on the last page")
		}
		body, err := apiCall(conf.next, cache)
		if err != nil {
			return err
		}

		locationdata := LocationPage{}
		err = json.Unmarshal(body, &locationdata)
		if err != nil {
			return err
		}

		conf.prev = (locationdata.Previous)
		conf.next = locationdata.Next

		for _, location := range locationdata.Results {
			fmt.Println(location.Name)
		}
		return nil
	}
}

func GetPrevLocationArea(conf *config, cache *internal.Cache) func() error {
	return func() error {
		if conf.prev == "" {
			return fmt.Errorf("you're on the first page")
		}
		body, err := apiCall(conf.prev, cache)
		if err != nil {
			return err
		}

		locationdata := LocationPage{}
		err = json.Unmarshal(body, &locationdata)
		if err != nil {
			return err
		}

		conf.prev = (locationdata.Previous)
		conf.next = locationdata.Next

		for _, location := range locationdata.Results {
			fmt.Println(location.Name)
		}
		return nil
	}

}

func GetLocationData(conf *config, cache *internal.Cache) func() error {
	return func() error {
		if conf.param == "" {
			return fmt.Errorf("no location given")
		}
		url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", conf.param)
		body, err := apiCall(url, cache)
		if err != nil {
			return err
		}

		location := Location{}
		err = json.Unmarshal(body, &location)
		if err != nil {
			return err
		}

		for _, encounter := range location.PokemonEncounters {
			fmt.Println(encounter.Pokemon.Name)
		}
		return nil
	}
}

func GetCatchPokemon(conf *config, cache *internal.Cache) func() error {
	return func() error {
		if conf.param == "" {
			return fmt.Errorf("No pokemon targeted")
		}
		fmt.Printf("Throwing a Pokeball at %s...\n", conf.param)

		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s",conf.param)
		body, err := apiCall(url, cache)
		if err != nil {
			return err
		}

		pokemon := Pokemon{}
		err = json.Unmarshal(body, &pokemon)
		if err != nil {
			return err
		}

		catchPokemon(pokemon, conf.pokedex)

		return nil
	}
}

func apiCall(url string, cache *internal.Cache) ([]byte, error) {
	cachebody, iscached := cache.Get(url)
	if iscached {
		//fmt.Println("received data from Cache")
		return cachebody, nil
	}
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	cache.Add(url, body)
	//fmt.Println("received data from API")
	return body, nil
}

