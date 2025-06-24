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
	group.GET("/:name", h.GetByName) // GET /characters/:name
	group.GET("", h.GetAll)          // GET /characters
}

type getByNameRequest struct {
	Name string `uri:"name" binding:"required"`
}

// GetByName handles GET /characters/:name
func (h *Handler) GetByName(c *gin.Context) {
	var req getByNameRequest

	// Bind the request parameters
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request parameters"})
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
