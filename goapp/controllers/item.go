package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/inventoryserver/api"
	"github.com/pramodshenkar/inventoryserver/models"
)

func AddItem(c *gin.Context) {

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

func GetAllItems(c *gin.Context) {
	courses, err := api.GetAllItems()

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": err})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"items": courses})
}

func GetItem(c *gin.Context) {

	var itemRequest struct {
		ItemCode string `json:"itemcode" binding:"required"`
	}

	if c.BindJSON(&itemRequest) != nil {
		fmt.Println("Provide required details for GetCoursesByID")
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	item, err := api.GetItem(itemRequest.ItemCode)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.JSON(200, gin.H{"item": item})

}

func DeleteItem(c *gin.Context) {
	var item struct {
		ItemID string `json:"_id" bson:"_id"`
	}

	if c.BindJSON(&item) != nil {
		fmt.Println("Provide required details for GetCoursesByID")
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	result, err := api.DeleteItem(item.ItemID)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": err})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"item": result})
}

func UpdateItem(c *gin.Context) {

	item := models.Item{}

	err := c.BindJSON(&item)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": "Provide required details"})
		c.Abort()
		return
	}

	result, err := api.UpdateItem(item)

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": err})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"item": result})

}
