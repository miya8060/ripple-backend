package api

import (
	"github.com/gin-gonic/gin"
	"github.com/miya8060/ripple-backend/internal/middleware"
)

// SetupRouter はAPIルーターの設定を行う
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// CORSミドルウェアの設定
	r.Use(middleware.SetupCors())

	// APIエンドポイントの設定
	r.GET("/api/health", HealthCheck)

	// TODO: 以下のようなグループを追加していく
	// auth := r.Group("/api/auth")
	// users := r.Group("/api/users")
	// rooms := r.Group("/api/rooms")
	// messages := r.Group("/api/messages")
	// media := r.Group("/api/media")

	return r
}
