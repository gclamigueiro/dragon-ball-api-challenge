package character

import (
	"errors"
	"fmt"

	"github.com/gclamigueiro/dragon-ball-api/internal/client/dragonball"
)

var (
	ErrCharacterNotFound = errors.New("character not found")
	ErrNameEmpty         = errors.New("character name cannot be empty")
	ErrInvalidCharacter  = errors.New("invalid character data")
	ErrDatabase          = errors.New("database error")
)

type Service interface {
	GetByName(name string) (*Character, error)
	GetAll() ([]*Character, error)
}

type service struct {
	dgzClient  dragonball.Client
	repository Repository
}

func NewService(dgzClient dragonball.Client, storage Repository) Service {
	return &service{
		dgzClient:  dgzClient,
		repository: storage,
	}
}

// GetByName retrieves a character by name (case-insensitive)
func (s *service) GetByName(name string) (*Character, error) {

	if name == "" {
		return nil, ErrNameEmpty
	}

	// Try to find in local DB
	character, err := s.repository.FindByName(name)
	if err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabase, err)
	}

	if character != nil {
		return character, nil
	}

	// Fetch from external API
	apiCharacter, err := s.dgzClient.GetCharacterByName(name)
	if err != nil {
		return nil, fmt.Errorf("external API error: %w", err)
	}
	if apiCharacter == nil {
		return nil, ErrCharacterNotFound
	}

	character = FromAPIResponse(apiCharacter)

	// Validate the API response
	if !character.IsValid() {
		return nil, fmt.Errorf("%w: character data is invalid received from the api", ErrInvalidCharacter)
	}

	if err := s.repository.Save(character); err != nil {
		return nil, fmt.Errorf("%w: %v", ErrDatabase, err)
	}

	return character, nil
}

// GetAll retrieves all characters from the local database
func (s *service) GetAll() ([]*Character, error) {
	characters, err := s.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("failed to get all characters: %w", err)
	}
	return characters, nil
}
