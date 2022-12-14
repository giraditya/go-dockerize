package httpx

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	var app = gin.Default()
	app.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"*"},
		AllowHeaders:  []string{"*"},
		AllowWildcard: true,
	}))

	initRoutesBassic(app)
	return app
}

func initRoutesBassic(route *gin.Engine) {
	groupRoute := route.Group("/v1")
	groupRoute.GET("/check/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "oks",
		})
	})
}
