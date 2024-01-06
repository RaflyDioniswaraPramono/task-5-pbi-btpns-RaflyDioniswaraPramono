package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func Auth(c *gin.Context) {
	prefix := "Bearer "
	authHeader := c.Request.Header.Get("Authorization")
	reqToken := strings.TrimPrefix(authHeader, prefix)

	if authHeader == "" || reqToken == authHeader {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})

		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User authorized!"})
}
