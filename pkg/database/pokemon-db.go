package database

import "context"

type PokemonDB interface {
	SavePokemon(ctx context.Context, pokemon PokemonModel) error
	GetPokemon(ctx context.Context, id int) (PokemonModel, error)
}
