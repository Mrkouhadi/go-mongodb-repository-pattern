package internal

import (
	"github.com/alexedwards/scs/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Config struct {
	Mongo        *mongo.Client
	Session      *scs.SessionManager
	InProduction bool
}
