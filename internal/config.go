package internal

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Mongo        *mongo.Client
	InProduction bool
}
