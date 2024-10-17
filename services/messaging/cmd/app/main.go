package main

import (
	"fmt"
	"log"

	"go-gin-boilerplate/config"
	"go-gin-boilerplate/pkg/mongodb"
)

func main() {
	mongoURI := config.GetMongoDBURI()
	mongoDB := config.GetMongoDBDatabase()

	fmt.Println("MongDB url:" + mongoURI)
	mongodb.InitMongoDB(mongoURI)
	log.Printf("Connected to MongoDB database: %s", mongoDB)
}
