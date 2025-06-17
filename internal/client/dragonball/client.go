package dragonball

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Client interface {
	GetCharacterByName(name string) (*Character, error)
}

type apiClient struct {
	httpClient *http.Client
	baseURL    string
}

func NewClient(baseUrl string) Client {
	return &apiClient{
		httpClient: &http.Client{},
		baseURL:    baseUrl,
	}
}

func (c *apiClient) GetCharacterByName(name string) (*Character, error) {
	// Encode query param
	endpoint, err := url.Parse(c.baseURL + "/characters")
	if err != nil {
		return nil, fmt.Errorf("failed to parse base URL: %w", err)
	}

	query := endpoint.Query()
	query.Set("name", name)
	endpoint.RawQuery = query.Encode()

	fmt.Println("Requesting character by name:", name, "at", endpoint.String())

	resp, err := c.httpClient.Get(endpoint.String())
	if err != nil {
		return nil, fmt.Errorf("failed to make request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d for name %q", resp.StatusCode, name)
	}

	var characters CharacterResponse
	if err := json.NewDecoder(resp.Body).Decode(&characters); err != nil {
		return nil, fmt.Errorf("failed to decode character response: %w", err)
	}

	if len(characters) == 0 {
		return nil, nil
	}

	return characters[0], nil
}
