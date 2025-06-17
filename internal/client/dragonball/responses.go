package dragonball

type CharacterResponse []*Character

type Character struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Ki   string `json:"ki"`
	Race string `json:"race"`
}
