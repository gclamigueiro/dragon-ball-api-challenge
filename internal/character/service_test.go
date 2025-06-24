package character_test

import (
	"errors"
	"testing"

	"github.com/gclamigueiro/dragon-ball-api/internal/character"
	"github.com/gclamigueiro/dragon-ball-api/internal/character/mocks"
	"github.com/gclamigueiro/dragon-ball-api/internal/client/dragonball"
	mock_dragonball "github.com/gclamigueiro/dragon-ball-api/internal/client/dragonball/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestService_GetByName_FoundInRepository(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockClient := new(mock_dragonball.Client)
	svc := character.NewService(mockClient, mockRepo)

	expected := &character.Character{Name: "Goku"}
	mockRepo.On("FindByName", "Goku").Return(expected, nil)

	result, err := svc.GetByName("Goku")
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
	mockClient.AssertExpectations(t)
}

func TestService_GetByName_FoundInAPI(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockClient := new(mock_dragonball.Client)
	svc := character.NewService(mockClient, mockRepo)

	mockRepo.On("FindByName", "Vegeta").Return(nil, nil)
	apiChar := &dragonball.Character{ID: 2, Name: "Vegeta"}
	mockClient.On("GetCharacterByName", "Vegeta").Return(apiChar, nil)
	mockRepo.On("Save", mock.AnythingOfType("*character.Character")).Return(nil)

	result, err := svc.GetByName("Vegeta")
	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, "Vegeta", result.Name)
	mockRepo.AssertExpectations(t)
	mockClient.AssertExpectations(t)
}

func TestService_GetByName_EmptyName(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockClient := new(mock_dragonball.Client)
	svc := character.NewService(mockClient, mockRepo)

	result, err := svc.GetByName("")
	assert.ErrorIs(t, err, character.ErrNameEmpty)
	assert.Nil(t, result)
}

func TestService_GetByName_NotFound(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockClient := new(mock_dragonball.Client)
	svc := character.NewService(mockClient, mockRepo)

	mockRepo.On("FindByName", "Piccolo").Return(nil, nil)
	mockClient.On("GetCharacterByName", "Piccolo").Return(nil, nil)

	result, err := svc.GetByName("Piccolo")
	assert.ErrorIs(t, err, character.ErrCharacterNotFound)
	assert.Nil(t, result)
}

func TestService_GetAll_Success(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockClient := new(mock_dragonball.Client)
	svc := character.NewService(mockClient, mockRepo)

	expected := []*character.Character{{Name: "Goku"}, {Name: "Vegeta"}}
	mockRepo.On("FindAll").Return(expected, nil)

	result, err := svc.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, expected, result)
	mockRepo.AssertExpectations(t)
}

func TestService_GetAll_Error(t *testing.T) {
	mockRepo := new(mocks.Repository)
	mockClient := new(mock_dragonball.Client)
	svc := character.NewService(mockClient, mockRepo)

	mockRepo.On("FindAll").Return(nil, errors.New("db error"))

	result, err := svc.GetAll()
	assert.Error(t, err)
	assert.Nil(t, result)
	mockRepo.AssertExpectations(t)
}
