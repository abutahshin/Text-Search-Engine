package database

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
	"github.com/typesense/typesense-go/typesense"
	"github.com/typesense/typesense-go/typesense/api"
	"github.com/typesense/typesense-go/typesense/api/pointer"
)

var client *typesense.Client

// Initialize Typesense Client
func init() {

	err := gotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	apiKey := os.Getenv("TYPESENSE_API_KEY")
	host := os.Getenv("TYPESENSE_HOST")
	client = typesense.NewClient(
		typesense.WithServer(host),
		typesense.WithAPIKey(apiKey),
	)
}

type SearchResult struct {
	Title           string   `json:"title"`
	Authors         []string `json:"authors"`
	AverageRating   float64  `json:"average_rating"`
	PublicationYear int      `json:"publication_year"`
	ImageURL        string   `json:"image_url"`
	RatingsCount    int      `json:"ratings_count"`
}

func SearchBooks(c *gin.Context) {
	query := c.DefaultQuery("q", "")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Missing search query"})
		return
	}
	searchParams := &api.SearchCollectionParams{
		Q:       query,
		QueryBy: "title,authors",
		SortBy:  pointer.String("ratings_count:desc"),
		Limit:   pointer.Int(10),
	}

	// Perform the search
	res, err := client.Collection("books").Documents().Search(context.TODO(), searchParams)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Parse JSON response properly
	var result []*SearchResult
	// Iterate through search hits and convert to book struct
	for _, hit := range *res.Hits {
		var book SearchResult
		documentBytes, err := json.Marshal(hit.Document)
		if err != nil {
			log.Printf("Error marshalling document: %v\n", err)
			continue
		}
		err = json.Unmarshal(documentBytes, &book)
		if err != nil {
			log.Printf("Error unmarshalling document: %v\n", err)
			continue
		}
		result = append(result, &book)
	}
	c.JSON(http.StatusOK, result)
}
