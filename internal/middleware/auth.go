package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CookieController() gin.HandlerFunc {
	return func(c *gin.Context) {
		owner, err := c.Cookie("owner_id")
		if err != nil {
			owner = uuid.New().String()
			c.SetCookie("owner_id", owner, 3600, "/", "localhost", false, true)
		}
		c.Set("owner_id", owner)

		c.Next()
	}
}
