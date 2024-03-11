package pokeApi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	pokecache "github.com/q-sw/go-pokedexcli/internal/pokeCache"
)

const apiUrl = "https://pokeapi.co/api/v2"

func getRequest(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return []byte{}, err
	}

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

func GetLocationDetails(location string, cache pokecache.Cache) (PokeLocationDetails, error) {
	endpoint := apiUrl + "/location-area/" + location

	r, ok := cache.Get(endpoint)
	if !ok {
		fmt.Println("Not in Cache")
		var err error
		r, err = getRequest(endpoint)
		if err != nil {
			fmt.Println("Error in GetLocationDetails request")
			return PokeLocationDetails{}, err
		}
		cache.Add(endpoint, r)
	}
	var locations PokeLocationDetails

	json.Unmarshal(r, &locations)

	return locations, nil
}
