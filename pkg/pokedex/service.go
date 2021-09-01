package pokedex

import (
	"context"

	"github.com/bal3000/PokeCentreAPI/pkg/database"
)

type Service interface {
	GetAllPokemon(ctx context.Context, size int, page int) ([]Pokemon, error)
	FindPokemon(ctx context.Context, name string, pokemonType string) ([]Pokemon, error)
	GetPokemon(ctx context.Context, id int) (Pokemon, error)
	ListMoves(ctx context.Context, id int) ([]string, error)
}

type service struct {
	db database.PokemonDB
}

func NewService(db database.PokemonDB) Service {
	return &service{db}
}

func (s *service) GetAllPokemon(ctx context.Context, size int, page int) ([]Pokemon, error) {
	pokemon, err := s.db.GetAllPokemon(ctx, 10, 1)
	if err != nil {
		return nil, err
	}

	var pmons []Pokemon
	for _, p := range pokemon {
		pmons = append(pmons, Pokemon{
			Name: p.Name,
		})
	}

	return pmons, nil
}

func (s *service) FindPokemon(ctx context.Context, name string, pokemonType string) ([]Pokemon, error) {
	return []Pokemon{}, nil
}

func (s *service) GetPokemon(ctx context.Context, id int) (Pokemon, error) {
	p, err := s.db.GetPokemon(ctx, id)
	if err != nil {
		return Pokemon{}, err
	}

	return Pokemon{
		Name: p.Name,
	}, nil
}

func (s *service) ListMoves(ctx context.Context, id int) ([]string, error) {
	return []string{}, nil
}
