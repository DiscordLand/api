package main

import (
	"math/rand"
	"net/http"

	"github.com/gin-gonic/gin"
)

const port = ":8001"

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func cors(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Next()
}

func send(ctx *gin.Context, file string) {
	if ctx.GetHeader("Accept") == "text/plain" {
		ctx.Header("Content-Type", "text/plain; charset=UTF-8")
		ctx.String(http.StatusOK, file)
	} else {
		ctx.JSON(http.StatusOK, gin.H{"file": file})
	}
}

func main() {
	cats, err := cats()
	handle(err)
	app := gin.New()

	app.Use(cors)
	app.StaticFile("/favicon.ico", "./assets/images/favicon.ico")

	random := app.Group("/random")

	random.GET("/cat", func(ctx *gin.Context) {
		send(ctx, cats[rand.Intn(len(cats)-1)])
	})

	handle(app.Run(port))
}
