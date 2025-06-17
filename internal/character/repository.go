package character

import (
	"errors"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Repository interface {
	FindAll() ([]*Character, error)
	FindByName(name string) (*Character, error)
	Save(character *Character) error
}

type repository struct {
	db *gorm.DB
}

func NewStorage(db *gorm.DB) Repository {
	return &repository{db}
}

func (r *repository) FindAll() ([]*Character, error) {
	var characters []*Character

	// Retrieve all characters from the database
	err := r.db.Find(&characters).Error
	if err != nil {
		return nil, err
	}
	return characters, nil
}

func (r *repository) FindByName(name string) (*Character, error) {
	var character Character

	// using %% to allow for partial matches. Similar to the api
	// that if you send "Go", it will return "Goku", "Gohan", etc.
	// and we will retrieve only the first match.
	err := r.db.
		Where("LOWER(name) LIKE LOWER(?)", name+"%").
		First(&character).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &character, err
}

func (r *repository) Save(character *Character) error {
	if character == nil {
		return errors.New("character cannot be nil")
	}
	// If the character already exists, do nothing
	err := r.db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "id"}}, // conflict target
		DoNothing: true,                          // skip insert if exists
	}).Create(character).Error

	return err
}
