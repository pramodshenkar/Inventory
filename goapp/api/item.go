package api

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func GetAllItems() ([]models.Item, error) {

	items := []models.Item{}

	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return []models.Item{}, err
	}

	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.ITEMS)

	filter := bson.D{{}}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return []models.Item{}, err
	}

	for cur.Next(context.TODO()) {
		item := models.Item{}

		err := cur.Decode(&item)
		if err != nil {
			return []models.Item{}, err
		}

		items = append(items, item)
	}

	cur.Close(context.TODO())

	if len(items) == 0 {
		return []models.Item{}, mongo.ErrNoDocuments
	}

	return items, nil
}

func GetItem(itemcode string) (models.Item, error) {
	item := models.Item{}

	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return models.Item{}, err
	}

	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.ITEMS)

	// objectID, _ := primitive.ObjectIDFromHex(itemid)
	filter := bson.D{primitive.E{Key: "itemcode", Value: itemcode}}

	err = collection.FindOne(context.TODO(), filter).Decode(&item)
	if err != nil {
		return models.Item{}, err
	}

	return item, nil
}

func DeleteItem(itemid string) (*mongo.DeleteResult, error) {
	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return nil, err
	}

	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.ITEMS)

	objectID, _ := primitive.ObjectIDFromHex(itemid)

	filter := bson.D{primitive.E{Key: "_id", Value: objectID}}

	result, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	return result, nil

}
