package main

import "github.com/gin-gonic/gin"

func SayHello(ctx *gin.Context) {
	ctx.JSON(200, "Hello World!")
}

func SayHelloWithName(ctx *gin.Context) {
	ctx.JSON(200, "Hello "+ctx.Param("name"))
}

func SayHelloWithQuery(ctx *gin.Context) {
	ctx.JSON(200, "Hello, this query is "+ctx.Query("age"))
}

func main() {

	server := gin.Default()

	// Modulariza el código
	helloPaths := server.Group("/hello")

	helloPaths.GET("", SayHello)
	helloPaths.GET("/:name", SayHelloWithName)
	helloPaths.GET("/queryParams", SayHelloWithQuery)

	//Captura la solicitud GET -> solo con server. para no repetir código se hace arriba
	// server.GET("/hello", SayHello)
	// server.GET("/hello/:name", SayHelloWithName)
	// server.GET("/hello/queryParams", SayHelloWithQuery)

	//localhost:8080
	server.Run(":8080")
}

// Crear un router con gin
// router := gin.Default()

// HANDLER
//Captura la solicitud GET "hello-world"
// router.GET("/hello-world", func(c *gin.Context)) {
// 	c.JSON(200, gin.H{
//		"message":"Hello World!"
// })
// })

// Correr servidor
// router.Run(":8081")
// go run main.go

// import (
// 	"encoding/json"
// 	"fmt"
// 	"log"
// )

// type product struct {
// 	Name      string
// 	Price     int
// 	Published bool
// }

// Marshal pasa a formato json
// p := product{
// 	Name:      "Mac",
// 	Price:     1,
// 	Published: true,
// }

// jsonData, err := json.Marshal(p)
// if err != nil {
// 	log.Fatal(err)
// }

// fmt.Println(string(jsonData))
