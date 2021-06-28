package api

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"

	"github.com/pramodshenkar/inventoryserver/connectionHelper"
	"github.com/pramodshenkar/inventoryserver/models"
)

func AddItem(item models.Item) (*mongo.InsertOneResult, error) {
	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return nil, err
	}

	item.DateCreated = time.Now()

	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.ITEMS)

	result, err := collection.InsertOne(context.TODO(), item)
	if err != nil {
		return nil, err
	}

	return result, nil

}
