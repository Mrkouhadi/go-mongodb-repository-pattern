package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/mrkouhadi/go-graphql-mongo/internal"
	"github.com/mrkouhadi/go-graphql-mongo/internal/db"
	"github.com/mrkouhadi/go-graphql-mongo/internal/handlers"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var app internal.Config

func main() {

	// set up mongodb database
	app.Mongo = ConnectToMongoDb("mongodb://localhost:27017")
	Mongodb_Repo := db.CreateNewMongoDbRepo(&app)

	// set up handlers
	repo := handlers.NewRepo(&app, Mongodb_Repo)
	handlers.NewHandlers(repo)

	// run the server
	server := &http.Server{
		Addr:    ":8080",
		Handler: routes(&app),
	}
	log.Printf("Listening on Port %s", ":8080")
	log.Fatal(server.ListenAndServe())
}

// ************ CONNECT to mongodb
func ConnectToMongoDb(url string) *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// CHECK CONNECTION
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" --------------- App has been Connected to MongoDB!")
	return client
}
