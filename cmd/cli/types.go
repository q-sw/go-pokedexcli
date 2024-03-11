package cli

import (
	pokeApi "github.com/q-sw/go-pokedexcli/internal/pokeApi"
	pokecache "github.com/q-sw/go-pokedexcli/internal/pokeCache"
)

type command struct {
	name        string
	description string
	callback    func(*state, ...string) error
}

type state struct {
	LocationNextUrl  *string
	LocationPrevtUrl *string
	PokeCache        pokecache.Cache
	PokemonCatch     map[string]pokeApi.Pokemon
}
