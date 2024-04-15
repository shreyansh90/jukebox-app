package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/shreyansh90/jukebox-app/controllers"
)

func SetupMusicAlbumRoutes(router *gin.Engine) {
	router.POST("/api/albums", controllers.CreateOrUpdateAlbum)
	router.GET("/api/albums", controllers.GetAlbumsSortedByReleaseDate)
	router.GET("/api/musicians/:musician_id/albums", controllers.GetAlbumsForMusicianSortedByPrice)
}
