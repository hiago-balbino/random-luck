package handler

import (
	"github.com/gin-gonic/gin"
)

// APIHandler is a struct that implements the Handler interface.
type APIHandler struct{}

// NewAPIHandler is a constructor for creating a new APIHandler instance.
func NewAPIHandler() APIHandler {
	return APIHandler{}
}

// Process is a function implementation to execute calls to create random luck numbers.
func (h APIHandler) Process(_ *gin.Context) {}
