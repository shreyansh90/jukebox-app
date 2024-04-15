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

var musicianCollection *mongo.Collection

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
	musicianCollection = client.Database("jukebox").Collection("musicians")
}

func CreateOrUpdateMusician(c *gin.Context) {
	var musician models.Musician
	if err := c.ShouldBindJSON(&musician); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := musicianCollection.InsertOne(context.Background(), musician)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Musician created/updated successfully", "data": musician})
}

func GetMusiciansForAlbumSortedByName(c *gin.Context) {
	// Get album ID from path parameters
	albumID := c.Param("album_id")

	var musicians []models.Musician

	options := options.Find().SetSort(bson.D{{"name", 1}})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.D{{"albums", albumID}}
	cursor, err := musicianCollection.Find(ctx, filter, options)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var musician models.Musician
		if err := cursor.Decode(&musician); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			return
		}
		musicians = append(musicians, musician)
	}

	if err := cursor.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": musicians})
}
