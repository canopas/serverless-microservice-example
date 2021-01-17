package main

import (
	"fmt"
	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func inLambda() bool {
	if lambdaTaskRoot := os.Getenv("LAMBDA_TASK_ROOT"); lambdaTaskRoot != "" {
		return true
	}
	return false
}

func setupRouter() *gin.Engine {

	r := gin.Default()

	r.GET("/hello-world", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "hello-world called successfully"})
	})

	return r
}

func main() {
	if inLambda() {
		fmt.Println("running aws lambda in aws")
		log.Fatal(gateway.ListenAndServe(":8880", setupRouter()))
	} else {
		fmt.Println("running aws lambda in local")
		log.Fatal(http.ListenAndServe(":8880", setupRouter()))
	}
}
