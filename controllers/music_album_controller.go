package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shreyansh90/jukebox-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var musicAlbumCollection *mongo.Collection

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}
	ctx := context.TODO()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}
	musicAlbumCollection = client.Database("jukebox").Collection("musicAlbums")
}

func CreateOrUpdateAlbum(c *gin.Context) {
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
	c.JSON(http.StatusCreated, gin.H{"message": "Music album created/updated successfully", "data": musicAlbum})
}

func GetAlbumsSortedByReleaseDate(c *gin.Context) {
	var albums []models.MusicAlbum

	options := options.Find().SetSort(bson.D{{"release_date", 1}})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := musicAlbumCollection.Find(ctx, bson.D{}, options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var album models.MusicAlbum
		if err := cursor.Decode(&album); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		albums = append(albums, album)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	// Return the sorted list of albums
	c.JSON(http.StatusOK, gin.H{"data": albums})
}

func GetAlbumsForMusicianSortedByPrice(c *gin.Context) {
	// Get musician ID from path parameters
	musicianID := c.Param("musician_id")

	var albums []models.MusicAlbum

	options := options.Find().SetSort(bson.D{{"price", 1}})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{"musicians", musicianID}}
	cursor, err := musicAlbumCollection.Find(ctx, filter, options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var album models.MusicAlbum
		if err := cursor.Decode(&album); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		albums = append(albums, album)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": albums})
}
