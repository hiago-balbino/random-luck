package handler

import (
	"github.com/gin-gonic/gin"
)

// WebHandler is a struct that implements the Handler interface.
type WebHandler struct{}

// NewWebHandler is a constructor for creating a new WebHandler instance.
func NewWebHandler() WebHandler {
	return WebHandler{}
}

// Process is a function implementation to execute calls to create random luck numbers.
func (h WebHandler) Process(_ *gin.Context) {}
