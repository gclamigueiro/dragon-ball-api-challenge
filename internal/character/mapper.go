package character

import "github.com/gclamigueiro/dragon-ball-api/internal/client/dragonball"

func FromAPIResponse(apiChar *dragonball.Character) *Character {
	return &Character{
		ID:   apiChar.ID,
		Name: apiChar.Name,
		Ki:   apiChar.Ki,
		Race: apiChar.Race,
	}
}
