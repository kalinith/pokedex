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

