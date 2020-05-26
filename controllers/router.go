package controllers

import (
    "IRIS_WEB/router"
    "github.com/kataras/iris/v12"
)

func err500(ctx iris.Context) {
    ctx.WriteString("500")
}

func err400(ctx iris.Context) {
    ctx.WriteString("400")
}

func InnerRouter(app *iris.Application) {
    app.OnErrorCode(iris.StatusInternalServerError, err500)
    app.OnErrorCode(iris.StatusBadRequest, err400)
    app.Post("/editor", router.UpdateContent)
}