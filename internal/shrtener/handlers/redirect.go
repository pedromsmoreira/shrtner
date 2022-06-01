package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *RestHandler) Redirect(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"message": "Redirect"})
}
