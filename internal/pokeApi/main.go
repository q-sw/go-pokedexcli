package pokeApi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/q-sw/go-pokedexcli/internal/pokeCache"
)

type PokeLocation struct {
	Count    int      `json:"count"`
	Next     *string  `json:"next"`
	Previous *string  `json:"previous"`
	Results  []Result `json:"results"`
}

type Result struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

const apiUrl = "https://pokeapi.co/api/v2"

func getRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	body, err := io.ReadAll(resp.Body)

	return body, nil

}

func GetLocation(url *string, cache pokecache.Cache) (PokeLocation, error) {
	endpoint := "/location-area"
	fullUrl := apiUrl + endpoint
	if url != nil {
		fullUrl = *url
	}

	r, ok := cache.Get(fullUrl)
	if !ok {
		fmt.Println("Not in Cache")
		var err error
		r, err = getRequest(fullUrl)
		if err != nil {
			return PokeLocation{}, err
		}
		cache.Add(fullUrl, r)
	}
	var locations PokeLocation

	json.Unmarshal(r, &locations)

	return locations, nil
}
