package db

import (
    "context"
    "log"
    "time"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

type Handler struct {
    DB *mongo.Client
}

func Connect() *mongo.Client  {
    mongoDbUri := os.Getenv("MONGO_DB_URI")

    if mongoDbUri == "" {
        log.Fatalf("mongoDbUri needs to be defined")
    }

    client, err := mongo.NewClient(options.Client().ApplyURI(mongoDbUri))
    if err != nil {
        log.Fatal(err)
    }
  
    ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
    err = client.Connect(ctx)
    if err != nil {
        log.Fatal(err)
    }

    //ping the database
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Connected to MongoDB")
    
    return Handler{client}
}

//getting database collections
func GetCollection(client *mongo.Client) *mongo.Collection {
    collection := client.Database("product_service").Collection("product")
    return collection
}