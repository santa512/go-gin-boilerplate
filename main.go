package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"go-gin-boilerplate/config"
	"go-gin-boilerplate/pkg/mongodb"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)

	mongoURI := config.GetMongoDBURI()
	mongoDB := config.GetMongoDBDatabase()

	fmt.Println("MongDB url:" + mongoURI)
	mongodb.InitMongoDB(mongoURI)
	log.Printf("Connected to MongoDB database: %s", mongoDB)
}
