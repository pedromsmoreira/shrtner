package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *RestHandler) Update(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "Update"})
}
