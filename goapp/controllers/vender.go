package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/inventoryserver/api"
)

func GetAllVenders(c *gin.Context) {
	categories, err := api.GetAllVenders()

	if err != nil {
		fmt.Println(err)
		c.JSON(400, gin.H{"message": err})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"items": categories})
}
