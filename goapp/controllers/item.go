package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/inventoryserver/api"
	"github.com/pramodshenkar/inventoryserver/models"
)

func AddItem(c *gin.Context) {
	// var itemRequest struct {
	// 	ItemName        string              `json:"itemname" bson:"itemname"`
	// 	ItemCode        string              `json:"itemcode" bson:"itemcode"`
	// 	ItemDescription string              `json:"itemdescription" bson:"itemdescription"`
	// 	ItemCategory    models.ItemCategory `json:"itemcategory" bson:"itemcategory"`
	// 	Quantity        int                 `json:"quantity" bson:"quantity"`
	// 	Vender          models.Vender       `json:"vender" bson:"vender"`
	// 	BasePrice       int                 `json:"basprice" bson:"baseprice"`
	// 	InternalProduct string              `json:"internalproduct" bson:"internalproduct"`
	// }

	// var itemRequest struct {
	// 	Item models.Item `json:item`
	// }

	item := models.Item{}

	err := c.BindJSON(&item)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	result, err := api.AddItem(item)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": err})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"item": result})

}
