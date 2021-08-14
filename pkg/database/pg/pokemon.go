package pg

import (
	"context"

	"github.com/bal3000/PokeCentreAPI/pkg/database"
	"github.com/jackc/pgx"
)

type PokemonPGDB struct {
	conn *pgx.Conn
}

func NewPokemonPGDB(host, database, username, password string, port int) (database.PokemonDB, func(), error) {
	config := pgx.ConnConfig{
		Host:     host,
		Port:     uint16(port),
		Password: password,
		User:     username,
		Database: database,
	}
	conn, err := pgx.Connect(config)
	if err != nil {
		return nil, nil, err
	}

	return PokemonPGDB{conn}, func() {
		conn.Close()
	}, nil
}

func (p PokemonPGDB) SavePokemon(ctx context.Context, pokemon database.PokemonModel) error {
	return nil
}

func (p PokemonPGDB) GetPokemon(ctx context.Context, id int) (database.PokemonModel, error) {
	return database.PokemonModel{}, nil
}
