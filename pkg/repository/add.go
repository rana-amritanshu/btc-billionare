package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

type Add struct {
	Amount   float32 `bson:"amount"`
	Datetime string  `bson:"datetime"`
}

type AddRepository struct {
	db *mongo.Database
}

func (a *AddRepository) Save(add *Add) error {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	_, err := a.db.Collection("wallet").InsertOne(ctx, add)
	if err != nil {
		log.Println("Error: DB Error |", time.Now().UTC().Format(time.RFC3339), err)
		return errors.New("Server Error")
	}

	return nil
}

func NewAddRepository(db *mongo.Database) *AddRepository {
	return &AddRepository{db}
}
