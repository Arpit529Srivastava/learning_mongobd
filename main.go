package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"time"
)

type movies struct {
	ID           string `json:"id" bson:"_id,omitempty"`
	Name         string `json:"name" bson:"name"`
	ReleaseDate  string `json:"date" bson:"release_date"`
	Director     string `json:"director" bson:"director"`
	Stars        int    `json:"stars" bson:"stars"`
}

var movieCollection *mongo.Collection

// Connect to MongoDB
func connectToMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to MongoDB!")
	return client, nil
}

// Fetch all movies
func GetAllMovies(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := movieCollection.Find(ctx, bson.M{})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not retrieve movies"})
		return
	}

	var movies []movies
	if err = cursor.All(ctx, &movies); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not decode movies"})
		return
	}

	c.IndentedJSON(http.StatusOK, movies)
}

// Fetch movie by ID
func moviesByID(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var movie movies
	err := movieCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&movie)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, movie)
}

// Fetch movie by name
func moviesByName(c *gin.Context) {
	name := c.Param("name")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var movie movies
	err := movieCollection.FindOne(ctx, bson.M{"name": name}).Decode(&movie)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, movie)
}

// Delete movie by ID
func deleteMovieByID(c *gin.Context) {
	id := c.Param("id")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := movieCollection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil || result.DeletedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "movie deleted"})
}

// Create a movie
func createMovie(c *gin.Context) {
	var newMovie movies
	if err := c.BindJSON(&newMovie); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "incorrect Json format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := movieCollection.InsertOne(ctx, newMovie)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "could not create movie"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "movie added successfully"})
}

// Update a movie
func UpdateMovie(c *gin.Context) {
	id := c.Param("id")
	var update movies
	if err := c.BindJSON(&update); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "can't be done in json/invaild json format"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := movieCollection.UpdateOne(ctx, bson.M{"_id": id}, bson.M{"$set": update})
	if err != nil || result.ModifiedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "movie not found or could not be updated"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "movie updated successfully"})
}

// Define the routes for the handler
func main() {
	client, err := connectToMongoDB()
	if err != nil {
		panic(err)
	}

	// Update the database and collection names
	movieCollection = client.Database("Movies").Collection("film")

	router := gin.Default()
	router.GET("/movies", GetAllMovies)
	router.GET("/movies/id/:id", moviesByID)
	router.GET("/movies/name/:name", moviesByName)
	router.DELETE("/movies/id/:id", deleteMovieByID)
	router.POST("/movie/add", createMovie)
	router.PUT("/movie/update/:id", UpdateMovie)
	router.Run("localhost:9090")
}
