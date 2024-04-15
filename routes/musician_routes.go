package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shreyansh90/jukebox-app/controllers"
)

func SetupMusicianRoutes(router *gin.Engine) {
	router.POST("/api/musicians", controllers.CreateMusician)
}
