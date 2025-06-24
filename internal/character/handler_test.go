package character_test

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/gclamigueiro/dragon-ball-api/internal/character"
	"github.com/gclamigueiro/dragon-ball-api/internal/character/mocks"
)

func setupRouter(handler *character.Handler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	handler.RegisterRoutes(r)
	return r
}

func TestGetByName_Success(t *testing.T) {
	mockService := new(mocks.Service)
	handler := character.NewHandler(mockService)
	router := setupRouter(handler)

	expectedChar := &character.Character{Name: "Goku"}
	mockService.On("GetByName", "goku").Return(expectedChar, nil)

	req, _ := http.NewRequest(http.MethodGet, "/characters/goku", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp character.Character
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Equal(t, "Goku", resp.Name)
	mockService.AssertExpectations(t)
}

func TestGetByName_NotFound(t *testing.T) {
	mockService := new(mocks.Service)
	handler := character.NewHandler(mockService)
	router := setupRouter(handler)

	mockService.On("GetByName", "Vegeta").Return(nil, character.ErrCharacterNotFound)

	req, _ := http.NewRequest(http.MethodGet, "/characters/Vegeta", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
	var resp map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, character.ErrCharacterNotFound.Error(), resp["error"])
	mockService.AssertExpectations(t)
}

func TestGetByName_InternalError(t *testing.T) {
	mockService := new(mocks.Service)
	handler := character.NewHandler(mockService)
	router := setupRouter(handler)

	mockService.On("GetByName", "piccolo").Return(nil, errors.New("some internal error"))

	req, _ := http.NewRequest(http.MethodGet, "/characters/piccolo", nil)
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var resp map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "some internal error", resp["error"])
	mockService.AssertExpectations(t)
}

func TestGetAll_Success(t *testing.T) {
	mockService := new(mocks.Service)
	handler := character.NewHandler(mockService)
	router := setupRouter(handler)

	expectedChars := []*character.Character{
		{Name: "Goku"},
		{Name: "Vegeta"},
	}
	mockService.On("GetAll").Return(expectedChars, nil)

	req, _ := http.NewRequest(http.MethodGet, "/characters", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	var resp []character.Character
	err := json.Unmarshal(w.Body.Bytes(), &resp)
	assert.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "Goku", resp[0].Name)
	assert.Equal(t, "Vegeta", resp[1].Name)
	mockService.AssertExpectations(t)
}

func TestGetAll_InternalError(t *testing.T) {
	mockService := new(mocks.Service)
	handler := character.NewHandler(mockService)
	router := setupRouter(handler)

	mockService.On("GetAll").Return(nil, errors.New("db error"))

	req, _ := http.NewRequest(http.MethodGet, "/characters", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusInternalServerError, w.Code)
	var resp map[string]interface{}
	_ = json.Unmarshal(w.Body.Bytes(), &resp)
	assert.Equal(t, "Failed to retrieve characters", resp["error"])
	mockService.AssertExpectations(t)
}
