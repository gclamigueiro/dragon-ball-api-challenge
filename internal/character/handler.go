package character

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes(r *gin.Engine) {
	group := r.Group("/characters")
	group.POST("", h.GetByName) // POST /characters
	group.GET("", h.GetAll)     // GET /characters
}

type getByNameRequest struct {
	Name string `json:"name" binding:"required"`
}

// GetByName handles POST /characters
func (h *Handler) GetByName(c *gin.Context) {
	var req getByNameRequest

	// Bind and validate JSON body
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing or invalid 'name'"})
		return
	}

	// Use the service to get character by name
	char, err := h.service.GetByName(req.Name)

	if err != nil {
		switch err {
		case ErrCharacterNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		case ErrInvalidCharacter:
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Return the character
	c.JSON(http.StatusOK, char)
}

// List all characters saved in the database
func (h *Handler) GetAll(c *gin.Context) {
	// Use the service to get all characters
	characters, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve characters"})
		return
	}

	// Return the list of characters
	c.JSON(http.StatusOK, characters)
}
