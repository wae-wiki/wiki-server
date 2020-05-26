package router

import (
    "github.com/kataras/iris/v12"
)

type ContentStr struct {
    Content string `json:"content"`
}

func UpdateContent(ctx iris.Context) {
    c := &ContentStr{}
    println("here")
    if err := ctx.ReadJSON(c); err != nil {
        panic(err.Error())
    } else {
        ctx.Header("Access-Control-Allow-Methods", "*")
        ctx.Header("Access-Control-Allow-Origin", "*")
        ctx.JSON(c)
    }
}
