package main

import (
    "context"
    "log"

    "github.com/go-redis/redis"
    "github.com/gofiber/fiber/v2"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
    ctx := context.Background()

    // Connect to Redis
    redisClient, redisErr := connectRedis("redis://:@redis_cache:6379")
    if redisErr != nil {
        panic(redisErr)
    }
    defer redisClient.Close()
    if _, err := redisClient.Ping().Result(); err != nil {
        panic(err)
    }
    log.Println("Connected to Redis")

    // Connect to MongoDB
    mongoClient, mongoErr := connectMongo(ctx, "mongodb://root:password@mongo_db:27017")
    if mongoErr != nil {
        panic(mongoErr)
    }
    defer mongoClient.Disconnect(ctx)
    log.Println("Connected to MongoDB")

    // Create Fiber app
    app := fiber.New()

    // Add routes
    app.Get("/hello", func(c *fiber.Ctx) error {
        return c.JSON(fiber.Map{
            "response": "Hi",
        })
    })

    // Start app server
    log.Fatal(app.Listen(":9000"))
}

func connectRedis(redisAddress string) (*redis.Client, error) {
    redisOptions, err := redis.ParseURL(redisAddress)
    if err != nil {
        return nil, err
    }
    client := redis.NewClient(redisOptions)
    return client, nil
}

func connectMongo(ctx context.Context, mongoAddress string) (*mongo.Client, error) {
    client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoAddress))
    return client, err
}
