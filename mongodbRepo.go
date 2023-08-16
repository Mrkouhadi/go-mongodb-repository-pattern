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
func CreateNewMongoDbRepo(url string, ctx context.Context) *MongodbRepo {
	return &MongodbRepo{
		MongodbClient: connectToMongoDb(url, ctx),
	}
}

// ************ CONNECT to mongodb
func connectToMongoDb(url string, ctx context.Context) *mongo.Client {
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

// //////////////////////////////////////////// INSERTION ///////////////////////////////////////////
// *********** insert one BOOK
func (mr *MongodbRepo) addBook(ctx context.Context, b Book) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")
	insertResult, err := booksCollection.InsertOne(ctx, b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" --------------- Inserted a single Book: ", insertResult.InsertedID)
}

// *********** insert multiple books at a time
func (mr *MongodbRepo) addMultipleBooks(ctx context.Context, booksTobeInserted ...Book) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	books := []interface{}{}
	for _, book := range booksTobeInserted {
		books = append(books, book)
	}
	insertManyResult, err := booksCollection.InsertMany(ctx, books)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" --------------- Inserted multiple documents: ", insertManyResult.InsertedIDs)
}

// ///////////////////////////////////////////////// UPDATE documents /////////////////////////////////////////
func (mr *MongodbRepo) UpdateBookTitle(ctx context.Context, id int, title string) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: title},
		}},
	}
	updateResult, err := booksCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" ------------ Edit Title: Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}
func (mr *MongodbRepo) increaseBookPrice(ctx context.Context, id int, inc int) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "price", Value: inc},
		}},
	}
	updateResult, err := booksCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" ------------ Increase price: Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
}

// /////////////////////////////////////////  FIND documents  ///////////////////////////////////////////
// find a single dcument
func (mr *MongodbRepo) findeSingleBook(ctx context.Context, id int) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	filter := bson.D{{Key: "id", Value: id}}

	var result Book

	err := booksCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(" --------------- Found a single document: %+v\n", result)
}

// find multiple documents
func (mr *MongodbRepo) findMultipleBooks(ctx context.Context, limit int64) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	var results []*Book
	// Finding multiple documents returns a cursor
	cur, err := booksCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Iterate through the cursor
	for cur.Next(ctx) {
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
	cur.Close(ctx)

	fmt.Printf(" --------------- Found multiple documents (array of pointers): %+v\n", results) //array of pointer
	fmt.Println("- - -Here's the second found book (value of pointer 2) : ", *results[2])       // get the value of the pointers
}

// ////////////////////////////////////////  DELETE DOCUMENTS  ///////////////////////////////////////////
// delete only single document
func (mr *MongodbRepo) deleteSingleBook(ctx context.Context, id int) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	deleteResult, err := booksCollection.DeleteMany(ctx, bson.D{{Key: "id", Value: id}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" --------------- Deleted %v documents in the books collection\n", deleteResult.DeletedCount)
}

// / delete all documents
func (mr *MongodbRepo) deleteAllBooks(ctx context.Context) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")
	deleteResult, err := booksCollection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(" --------------- Deleted %v documents in the books collection\n", deleteResult.DeletedCount)
}

// /////////////////////////////////////////  DISCONNECTION  ///////////////////////////////////////////
func (mr *MongodbRepo) disconnectDB(ctx context.Context) {
	err := mr.MongodbClient.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" --------------- Connection to MongoDB has been closed.")
}
