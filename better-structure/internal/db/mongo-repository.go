package db

import (
	"context"
	"fmt"

	"github.com/mrkouhadi/go-graphql-mongo/internal"
	"github.com/mrkouhadi/go-graphql-mongo/internal/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ************ Mongodb repository type
type MongodbRepo struct {
	MongodbClient *mongo.Client
}

// ************ create new mongodb repository
func CreateNewMongoDbRepo(a *internal.Config) *MongodbRepo {
	return &MongodbRepo{
		MongodbClient: a.Mongo,
	}
}

// //////////////////////////////////////////// INSERTION ///////////////////////////////////////////
// *********** insert one BOOK
func (mr *MongodbRepo) AddBook(ctx context.Context, b models.Book) error {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")
	insertResult, err := booksCollection.InsertOne(ctx, b)
	if err != nil {
		return err
	}
	fmt.Println(" --------------- Inserted a single Book: ", insertResult.InsertedID)
	return nil
}

// *********** insert multiple books at a time
func (mr *MongodbRepo) AddMultipleBooks(ctx context.Context, booksTobeInserted ...models.Book) error {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	books := []interface{}{}
	for _, book := range booksTobeInserted {
		books = append(books, book)
	}
	insertManyResult, err := booksCollection.InsertMany(ctx, books)
	if err != nil {
		return err
	}
	fmt.Println(" --------------- Inserted multiple documents: ", insertManyResult.InsertedIDs)
	return nil
}

// ///////////////////////////////////////////////// UPDATE documents /////////////////////////////////////////
func (mr *MongodbRepo) UpdateBookTitle(ctx context.Context, id int, title string) error {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "name", Value: title},
		}},
	}
	updateResult, err := booksCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	fmt.Printf(" ------------ Edit Title: Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}
func (mr *MongodbRepo) IncreaseBookPrice(ctx context.Context, id int, inc int) error {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	filter := bson.D{{Key: "id", Value: id}}
	update := bson.D{
		{Key: "$inc", Value: bson.D{
			{Key: "price", Value: inc},
		}},
	}
	updateResult, err := booksCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	fmt.Printf(" ------------ Increase price: Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)
	return nil
}

// /////////////////////////////////////////  FIND documents  ///////////////////////////////////////////
// find a single dcument
func (mr *MongodbRepo) FindSingleBook(ctx context.Context, id int) (models.Book, error) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	filter := bson.D{{Key: "id", Value: id}}

	var result models.Book

	err := booksCollection.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		return result, err
	}

	fmt.Printf(" --------------- Found a single document: %+v\n", result)
	return result, nil
}

// find multiple documents
func (mr *MongodbRepo) FindMultipleBooks(ctx context.Context, limit int64) ([]*models.Book, error) {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	findOptions := options.Find()
	findOptions.SetLimit(limit)

	var results []*models.Book
	// Finding multiple documents returns a cursor
	cur, err := booksCollection.Find(ctx, bson.D{{}}, findOptions)
	if err != nil {
		return results, err
	}

	// Iterate through the cursor
	for cur.Next(ctx) {
		var elem models.Book
		err := cur.Decode(&elem)
		if err != nil {
			return results, err
		}
		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		return results, err
	}

	// Close the cursor once finished
	cur.Close(ctx)

	fmt.Printf(" --------------- Found multiple documents (array of pointers): %+v\n", results) //array of pointer
	fmt.Println("- - -Here's the second found book (value of pointer 2) : ", *results[2])       // get the value of the pointers
	return results, nil
}

// ////////////////////////////////////////  DELETE DOCUMENTS  ///////////////////////////////////////////
// delete only single document
func (mr *MongodbRepo) DeleteSingleBook(ctx context.Context, id int) error {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")

	deleteResult, err := booksCollection.DeleteMany(ctx, bson.D{{Key: "id", Value: id}})
	if err != nil {
		return err
	}
	fmt.Printf(" --------------- Deleted %v documents in the books collection\n", deleteResult.DeletedCount)
	return nil
}

// / delete all documents
func (mr *MongodbRepo) DeleteAllBooks(ctx context.Context) error {
	booksCollection := mr.MongodbClient.Database("library-db").Collection("books")
	deleteResult, err := booksCollection.DeleteMany(ctx, bson.D{{}})
	if err != nil {
		return err
	}
	fmt.Printf(" --------------- Deleted %v documents in the books collection\n", deleteResult.DeletedCount)
	return nil
}

// /////////////////////////////////////////  DISCONNECTION  ///////////////////////////////////////////
func (mr *MongodbRepo) DisconnectDB(ctx context.Context) error {
	err := mr.MongodbClient.Disconnect(ctx)
	if err != nil {
		return err
	}
	fmt.Println(" --------------- Connection to MongoDB has been closed.")
	return nil
}
