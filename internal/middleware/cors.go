package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupCors() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}

	return cors.New(config)
}
