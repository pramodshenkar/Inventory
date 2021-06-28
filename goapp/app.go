package main

import (
	"github.com/gin-gonic/gin"
	"github.com/pramodshenkar/inventoryserver/controllers"
)

func main() {

	router := gin.Default()
	router.POST("/additem", controllers.AddItem)
	router.Run(":5000")

}
