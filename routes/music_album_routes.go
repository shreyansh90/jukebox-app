package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shreyansh90/jukebox-app/controllers"
)

func SetupMusicAlbumRoutes(router *gin.Engine) {
	router.POST("/api/albums", controllers.CreateMusicAlbum)
}
