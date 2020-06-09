package router

import (
    "IRIS_WEB/mongo"
    "github.com/kataras/iris/v12"
    "go.mongodb.org/mongo-driver/bson"
)

type ContentStr struct {
    Title string `json:"title"`
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
        res := mongo.InsertOne(c.Title, c.Content)
        ctx.WriteString(res)
    }
}

func GetArticlesList(ctx iris.Context) {
    list, err := mongo.FindAll("content", bson.M{})
    if err != nil {
        ctx.JSON(err)
    }
    ctx.Header("Access-Control-Allow-Methods", "*")
    ctx.Header("Access-Control-Allow-Origin", "*")
    ctx.JSON(list)
}
