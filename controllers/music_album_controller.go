package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shreyansh90/jukebox-app/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var musicAlbumCollection *mongo.Collection

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	musicAlbumCollection = client.Database("jukebox").Collection("musicAlbums")
}

func CreateMusicAlbum(c *gin.Context) {
	var musicAlbum models.MusicAlbum
	if err := c.ShouldBindJSON(&musicAlbum); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := musicAlbumCollection.InsertOne(context.Background(), musicAlbum)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Music album created successfully", "data": musicAlbum})
}
