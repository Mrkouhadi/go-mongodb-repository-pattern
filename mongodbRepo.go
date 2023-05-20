package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ************ Mongodb repository type
type MongodbRepo struct {
	MongodbClient *mongo.Client
}

// ************ create new mongodb repository
func CreateNewMongoDbRepo(url string) *MongodbRepo {
	return &MongodbRepo{
		MongodbClient: connectToMongoDb(url),
	}
}

// ************ CONNECT to mongodb
func connectToMongoDb(url string) *mongo.Client {
	clientOptions := options.Client().ApplyURI(url)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	fmt.Printf("%T", client)
	if err != nil {
		log.Fatal(err)
	}
	// CHECK CONNECTION
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" --------------- Connected to MongoDB!")
	return client
}

// //////////////////////////////////////////// METHODS
// *********** first method: insert one BOOK
func (mr *MongodbRepo) addBook(b Book) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")
	insertResult, err := booksCollection.InsertOne(context.TODO(), b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" --------------- Inserted a single Book: ", insertResult.InsertedID)
}

// *********** insert multiple books at a time
func (mr *MongodbRepo) addMultipleBooks(booksTobeInserted ...Book) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	books := []interface{}{}
	for _, book := range booksTobeInserted {
		books = append(books, book)
	}
	insertManyResult, err := booksCollection.InsertMany(context.TODO(), books)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" --------------- Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

// /////////////////////////////////////////  FIND documents  ///////////////////////////////////////////
// find a single dcument
func (mr *MongodbRepo) findeSingleBook(id int) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	filter := bson.D{{Key: "id", Value: id}}

	var result Book

	err := booksCollection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" --------------- Found a single document: %+v\n", result)
}

// find multiple documents
func (mr *MongodbRepo) findMultipleBooks(limit int64) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	var results []*Book
	// Finding multiple documents returns a cursor
	cur, err := booksCollection.Find(context.TODO(), bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		var elem Book
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	fmt.Printf(" --------------- Found multiple documents (array of pointers): %+v\n", results) //array of pointer
	fmt.Println("- - -Here's the second found book (value of pointer 2) : ", *results[2])       // get the value of the pointers
}

// ////////////////////////////////////////  DELETE DOCUMENTS  ///////////////////////////////////////////
// delete only single document
func (mr *MongodbRepo) deleteSingleBook(id int) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	deleteResult, err := booksCollection.DeleteMany(context.TODO(), bson.D{{Key: "id", Value: id}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" --------------- Deleted %v documents in the books collection\n", deleteResult.DeletedCount)
}

// / delete all documents
func (mr *MongodbRepo) deleteAllBooks() {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")
	deleteResult, err := booksCollection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" --------------- Deleted %v documents in the books collection\n", deleteResult.DeletedCount)
}

// /////////////////////////////////////////  DISCONNECTION  ///////////////////////////////////////////
func (mr *MongodbRepo) disconnectDB() {
	err := mr.MongodbClient.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" --------------- Connection to MongoDB closed.")
}
