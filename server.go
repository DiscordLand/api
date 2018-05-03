package main

import (
	"math/rand"

	"github.com/kataras/iris"
)

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	cats, err := cats()
	handle(err)
	app := iris.New()

	app.Use(func(ctx iris.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
	})

	app.Get("/random-cat", func(ctx iris.Context) {
		cat := cats[rand.Intn(len(cats)-1)]
		ctx.JSON(iris.Map{"file": cat})
	})

	app.Run(iris.Addr(":8008"))
}
