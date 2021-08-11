package pokedex

import "context"

type Service interface {
	GetAllPokemon(ctx context.Context, size int, page int) ([]Pokemon, error)
	FindPokemon(ctx context.Context, name string, pokemonType string) ([]Pokemon, error)
	GetPokemon(ctx context.Context, id int) (Pokemon, error)
	ListMoves(ctx context.Context, id int) ([]string, error)
}

type service struct {
	// db goes here
}

func NewService() Service {
	return service{}
}

func (s service) GetAllPokemon(ctx context.Context, size int, page int) ([]Pokemon, error) {
	return []Pokemon{}, nil
}

func (s service) FindPokemon(ctx context.Context, name string, pokemonType string) ([]Pokemon, error) {
	return []Pokemon{}, nil
}

func (s service) GetPokemon(ctx context.Context, id int) (Pokemon, error) {
	return Pokemon{}, nil
}

func (s service) ListMoves(ctx context.Context, id int) ([]string, error) {
	return []string{}, nil
}
