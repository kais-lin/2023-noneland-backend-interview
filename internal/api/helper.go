package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type okResponse struct {
	OK bool `json:"ok"`
}

func errResponse(c *gin.Context) {
	c.AbortWithStatusJSON(http.StatusInternalServerError, okResponse{
		OK: false,
	})
}
