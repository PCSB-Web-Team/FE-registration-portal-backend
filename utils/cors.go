package utils

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors(acceptedOrigins string) gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{acceptedOrigins}
	return cors.New(config)
}
