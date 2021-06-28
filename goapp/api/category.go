package api

import (
	"context"

	"github.com/pramodshenkar/inventoryserver/connectionHelper"
	"github.com/pramodshenkar/inventoryserver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllItemCategories() ([]models.ItemCategory, error) {

	itemCategories := []models.ItemCategory{}

	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return []models.ItemCategory{}, err
	}

	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.ITEMCATEGORY)

	filter := bson.D{{}}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return []models.ItemCategory{}, err
	}

	for cur.Next(context.TODO()) {
		itemCategory := models.ItemCategory{}

		err := cur.Decode(&itemCategory)
		if err != nil {
			return []models.ItemCategory{}, err
		}

		itemCategories = append(itemCategories, itemCategory)
	}

	cur.Close(context.TODO())

	if len(itemCategories) == 0 {
		return []models.ItemCategory{}, mongo.ErrNoDocuments
	}

	return itemCategories, nil
}
