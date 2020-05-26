package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	"github.com/kataras/iris/v12/middleware/recover"
	//"github.com/iris-contrib/middleware/cors"

	"IRIS_WEB/controllers"
)

func main() {
	app := iris.New()
	app.Logger().SetLevel("debug")

	app.Use(recover.New())
	app.Use(logger.New())

	controllers.InnerRouter(app)

	app.Run(iris.Addr(":8088"), iris.WithOptimizations)
}

