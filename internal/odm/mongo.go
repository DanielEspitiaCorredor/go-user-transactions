package odm

import (
	"context"
	"fmt"
	"os"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	doOnce sync.Once
	client *mongo.Client
	dbErr  error
	dbName = os.Getenv("DB_NAME")
)

func GetConnection() (*mongo.Client, error) {

	doOnce.Do(func() {

		config := options.Client().ApplyURI(os.Getenv("DB_MONGO_URI"))

		client, dbErr = mongo.Connect(context.Background(), config)

		dbErr = client.Ping(context.Background(), nil)

		if dbErr != nil {
			fmt.Println("[GetConnection] Error connecting to DB - ping fail", dbErr)
		}

		initCollections(client)

	})

	return client, dbErr

}

func initCollections(cli *mongo.Client) {
	tx := &Transaction{}

	cli.Database(dbName).Collection(tx.GetMongoCollection())
}
