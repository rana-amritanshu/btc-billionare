package internal

import (
	"os"

	"context"
	"fmt"
	"log"
	"time"

	httpAdapter "btc/pkg/http"
	"btc/pkg/repository"
	"btc/pkg/service"

	"github.com/labstack/echo/v4"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func Run(e *echo.Echo) {

	mongoDB := mongoInit()
	addRepo := repository.NewAddRepository(mongoDB)
	addService := service.NewAddService(addRepo)
	addHandler := httpAdapter.NewAddHandler(addService)

	listRepo := repository.NewListRepository(mongoDB)
	listService := service.NewListService(listRepo)
	listHandler := httpAdapter.NewListHandler(listService)

	e.Static("/docs", "docs")
	e.Static("/swagger-ui", "swagger-ui")
	e.POST("/wallet", addHandler.Save)
	e.GET("/wallet", listHandler.List)
}

func mongoInit() *mongo.Database {
	mongoString := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")

	fmt.Println("mongo string:", mongoString)
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoString))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(dbName)
}
