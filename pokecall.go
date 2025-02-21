package main
import (
	"net/http"
	"encoding/json"
	"fmt"
	"io"
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


func makeGetLocationArea(conf *config) func() error {
	return func() error {
		if conf.next == "" {
			return fmt.Errorf("you're on the last page")
		}
		body, err := apiCall(conf.next)
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

func GetPrevLocationArea(conf *config) func() error {
	return func() error {
		if conf.prev == "" {
			return fmt.Errorf("you're on the first page")
		}
		body, err := apiCall(conf.prev)
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

func apiCall(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
	
}

