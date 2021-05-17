package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// initial setup before running main
func init() {

	var err error

	// Load in environment variables
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MONGO_URI := os.Getenv("MONGO_CREDENTIALS")

	fmt.Println(MONGO_URI)

	// Setup mgm default config
	err = mgm.SetDefaultConfig(nil, "Cluster0", options.Client().ApplyURI(MONGO_URI))

	if err != nil {
		log.Fatal(err)
	}

}

func main() {

	var err error

	//book := NewBook("Wowza", 420)

	// // Make sure to pass the model by reference.
	// err = mgm.Coll(book).Create(book)

	// if err != nil {
	// 	fmt.Println("OVER HERE", err)
	// }

	result := []Book{}

	err = mgm.Coll(&Book{}).SimpleFind(&result,
		bson.M{
			"name": primitive.Regex{Pattern: "wowza", Options: "i"},
		},
	)

	if err != nil {
		fmt.Println("OVER HERE", err)
	}

	fmt.Println("RESULT:", result)

	fmt.Println("Hello world from go!")
}

type Book struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
	Pages            int    `json:"pages" bson:"pages"`
}

func (model *Book) CollectionName() string {
	return "my_books"
}

func NewBook(name string, pages int) *Book {
	return &Book{
		Name:  name,
		Pages: pages,
	}
}
