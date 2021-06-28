package api

import (
	"context"

	"github.com/pramodshenkar/inventoryserver/connectionHelper"
	"github.com/pramodshenkar/inventoryserver/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAllVenders() ([]models.Vender, error) {

	venders := []models.Vender{}

	client, err := connectionHelper.GetMongoClient()
	if err != nil {
		return []models.Vender{}, err
	}

	collection := client.Database(connectionHelper.DB).Collection(connectionHelper.VENDER)

	filter := bson.D{{}}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return []models.Vender{}, err
	}

	for cur.Next(context.TODO()) {
		vender := models.Vender{}

		err := cur.Decode(&vender)
		if err != nil {
			return []models.Vender{}, err
		}

		venders = append(venders, vender)
	}

	cur.Close(context.TODO())

	if len(venders) == 0 {
		return []models.Vender{}, mongo.ErrNoDocuments
	}

	return venders, nil
}
