package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/gorm"
	"pkg/src/routers/middleware"
)

func Routes(db *gorm.DB) *gin.Engine {

	environment := viper.Get("server.environment")
	if environment == "prod" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(middleware.CORSMiddleware())

	RegisterRoutes(router, db) //routes register

	return router
}
