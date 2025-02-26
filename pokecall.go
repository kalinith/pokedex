package main
import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
	"github.com/kalinith/pokedex/internal"
)

type locationpage struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type location struct {
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

func makeGetLocationArea(conf *config, cache *internal.Cache) func() error {
	return func() error {
		if conf.next == "" {
			return fmt.Errorf("you're on the last page")
		}
		body, err := apiCall(conf.next, cache)
		if err != nil {
			return err
		}
		locationdata := locationpage{}
		json.Unmarshal(body, &locationdata)
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
		locationdata := locationpage{}
		json.Unmarshal(body, &locationdata)
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
		if conf.next == "" {
			return fmt.Errorf("no location given")
		}
		url := fmt.Sprintf("%s/%s",conf.prev, conf.next)
		body, err := apiCall(url, cache)
		if err != nil {
			return err
		}
		loc := location{}
		json.Unmarshal(body, &loc)

		for _, encounter := range loc.PokemonEncounters {
			fmt.Println(encounter.Pokemon.Name)
		}
		return nil
	}
}

func apiCall(url string, cache *internal.Cache) ([]byte, error) {
	cachebody, iscached := cache.Get(url)
	if iscached {
		fmt.Println("received data from Cache")
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
	fmt.Println("received data from API")
	return body, nil
}

