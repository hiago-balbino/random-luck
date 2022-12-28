package handler

import (
	"github.com/gin-gonic/gin"
)

// Handler is an interface to support web functions.
type Handler interface {
	// Process is a function to execute calls to create random luck numbers.
	Process(*gin.Context)
}
