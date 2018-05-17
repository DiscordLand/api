package main

import (
	"math/rand"

	"github.com/kataras/iris"
)

const port = ":8008"

func handle(err error) {
	if err != nil {
		panic(err)
	}
}

func cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Next()
}

func send(ctx iris.Context, file string) {
	if ctx.GetHeader("accept") == "text/plain" {
		ctx.ContentType("text/plain")
		ctx.WriteString(file)
	} else {
		ctx.JSON(iris.Map{"file": file})
	}
}

func main() {
	cats, err := cats()
	handle(err)
	app := iris.New()

	app.Use(cors)

	random := app.Party("/random")

	random.Get("/cat", func(ctx iris.Context) {
		cat := cats[rand.Intn(len(cats)-1)]
		send(ctx, cat)
	})

	app.Run(iris.Addr(port))
}
