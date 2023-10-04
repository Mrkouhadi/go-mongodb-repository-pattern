package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mrkouhadi/go-graphql-mongo/internal"
	"github.com/mrkouhadi/go-graphql-mongo/internal/db"
	"github.com/mrkouhadi/go-graphql-mongo/internal/handlers"
	"github.com/mrkouhadi/go-graphql-mongo/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var app internal.Config
var session *scs.SessionManager

func main() {
	// Register what's gonna be stored in the sessions
	gob.Register(models.Book{})

	// set up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

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
