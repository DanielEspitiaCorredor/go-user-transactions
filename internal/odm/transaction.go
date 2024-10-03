package odm

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	Id        int                `json:"id" bson:"id"`
	Date      primitive.DateTime `json:"date" bson:"date"`
	Name      string             `json:"name" bson:"name"`
	Value     float64            `json:"value" bson:"value"`
	CreatedAt primitive.DateTime `json:"created_at" bson:"created_at"`
}

func (t *Transaction) GetMongoCollection() string {

	return "transactions"
}

func (t *Transaction) Insert() error {

	cli, err := GetConnection()
	if err != nil {
		fmt.Println("[Transaction][Insert] GetConnection err", err)
		return err
	}

	collection := cli.Database(dbName).Collection(t.GetMongoCollection())

	filter := bson.D{
		primitive.E{Key: "id", Value: t.Id},
	}

	count, err := collection.CountDocuments(context.Background(), filter)
	if err != nil {
		fmt.Println("[Transaction][Insert] CountDocuments err", err)
		return err
	}

	if count > 0 {
		fmt.Println("[Transaction][Insert] Omit document, already exist", t.Id)
	}

	t.CreatedAt = primitive.NewDateTimeFromTime(time.Now().UTC())

	_, err = collection.InsertOne(context.Background(), t)
	if err != nil {
		fmt.Println("[Transaction][Insert] InsertOne err", err)
		return err
	}

	return nil
}
