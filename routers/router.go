package routers

import (
	"url-shortener/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	err := r.SetTrustedProxies([]string{"127.0.0.1", "192.168.1.0/24"})
	if err != nil {
		panic("Failed to set trusted proxies: " + err.Error())
	}

	SetupRoutes(r)
	return r
}

func SetupRoutes(r *gin.Engine) {
	routes := r.Group("/v1/api/shortener")
	{
		routes.POST("/", handlers.ShortenURL)
		routes.GET("/:code", handlers.ResolveURL)
	}
}
