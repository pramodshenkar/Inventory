package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Item struct {
	ID              primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	ItemName        string             `json:"itemname" bson:"itemname" binding:"required"`
	ItemCode        string             `json:"itemcode" bson:"itemcode" binding:"required"`
	ItemDescription string             `json:"itemdescription" bson:"itemdescription"`
	ItemCategory    ItemCategory       `json:"itemcategory" bson:"itemcategory" binding:"required"`
	Quantity        int                `json:"quantity" bson:"quantity" binding:"required"`
	Vender          Vender             `json:"vender" bson:"vender" binding:"required"`
	BasePrice       int                `json:"baseprice" bson:"baseprice" binding:"required"`
	InternalProduct bool               `json:"internalproduct" bson:"internalproduct" binding:"required"`
	DateCreated     time.Time          `json:"datecreated" bson:"datecreated"`
	DateModified    time.Time          `json:"datemodified" bson:"datemodified"`
}

type ItemCategory struct {
	ID               primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" binding:"required"`
	ItemCategoryName string             `json:"itemcategoryname" bson:"itemcategoryname" binding:"required"`
}
type Vender struct {
	ID         primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty" binding:"required"`
	VenderName string             `json:"vendername" bson:"vendername" binding:"required"`
}
