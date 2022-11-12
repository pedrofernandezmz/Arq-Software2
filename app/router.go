package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func MapUrls(router *gin.Engine, dependencies *Dependencies) {
	// Products Mapping

	router.GET("/items/:id", dependencies.ItemController.Get)
	//router.GET("/items", dependencies.ItemController.Get)
	router.POST("/items", dependencies.ItemController.Insert)

	fmt.Println("Finishing mappings configurations")
}
