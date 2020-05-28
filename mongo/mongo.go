package mongo

import (
    "context"
    "encoding/json"
    "fmt"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/bson"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Article struct {
    title string
    content string
}

func openConnectMongo() *mongo.Client {
    clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/wiki")
    client, err := mongo.Connect(context.TODO(), clientOptions)
    if err != nil {
        log.Fatal(err)
        return nil
    }

    err = client.Ping(context.TODO(), nil)
    if err != nil {
        log.Fatal(err)
        return nil
    }
    fmt.Println("Connected to MongoDB")
    return client
}

func FindOne(collectionName string, filter bson.M) (bson.M, error) {
    mongoClient := openConnectMongo()
    defer mongoClient.Disconnect(context.Background())
    if mongoClient == nil {
        return nil, fmt.Errorf("没有连接到数据库")
    }
    ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
    collection := mongoClient.Database("wiki").Collection(collectionName)
    var result bson.M
    err := collection.FindOne(ctx, filter).Decode(&result)
    if err != nil {
        return nil, err
    }
    return result, nil
}

func FindAll(collectionName string, filter bson.M) ([]bson.M, error) {
    mongoClient := openConnectMongo()
    if mongoClient == nil {
        return nil, fmt.Errorf("没有连接到数据库")
    }
    ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)
    collection := mongoClient.Database("wiki").Collection(collectionName)
    cur, err := collection.Find(ctx, filter)
    if err != nil {
        return nil, err
    }
    defer cur.Close(ctx)
    var resultArr []bson.M
    for cur.Next(ctx) {
        var result bson.M
        err := cur.Decode(&result)
        if err != nil {
            return nil, err
        }
        resultArr = append(resultArr, result)
    }
    return resultArr, nil
}

func InsertOne(title string, content string) string {
    mongoClient := openConnectMongo()
    if mongoClient == nil {
        return "没有链接到数据库"
    }
    collection := mongoClient.Database("wiki").Collection("content")
    _, err := collection.InsertOne(context.Background(), bson.M{
        "title": title,
        "content": content,
    })
    if err != nil {
        return "插入失败"
    }
    defer mongoClient.Disconnect(context.Background())
    result := Article{title, content }
    b, _ := json.Marshal(result)
    res := string(b)
    return res
}


