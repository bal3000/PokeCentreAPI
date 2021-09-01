package pokedex

type Pokemon struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Types       []string `json:"types"`
	EvolvesInto int      `json:"evolvesInto"`
}
