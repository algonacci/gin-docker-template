package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Root struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

var root = Root{StatusCode: 200, Message: "Success fetching the API!"}

func getRoot(c *gin.Context) {
	c.JSON(http.StatusOK, root)
}

func main() {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.GET("/", getRoot)

	fmt.Println("🚀 [SERVER] is running on port http://localhost:8080")

	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
