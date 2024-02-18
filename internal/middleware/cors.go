package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

type Cors struct {
	log *zap.Logger
}

func NewCorsMiddleware(log *zap.Logger) *Cors {
	return &Cors{
		log: log,
	}
}

func (cors *Cors) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.log.Info("request", zap.String("path", c.Request.URL.Path), zap.String("method", c.Request.Method), zap.String("origin", c.Request.Header.Get("Origin")))
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")
		c.Header("Access-Control-Allow-Origin", origin)
		c.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type, new-token, New-Expires-At")
		c.Header("Access-Control-Allow-Credentials", "true")
		if method == "OPTIONS" {
			c.Status(http.StatusNoContent)
		}
		c.Next()
	}
}
