package main

import (
	"github.com/gin-gonic/gin"
	"github.com/shreyansh90/jukebox-app/routes"
)

func main() {
	r := gin.Default()

	routes.SetupMusicAlbumRoutes(r)
	routes.SetupMusicianRoutes(r)

	r.Run(":8080")
}
