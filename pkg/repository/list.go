package repository

import (
	"context"
	"errors"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Wallet struct {
	Amount   float64 `bson:"amount"`
	Datetime string  `bson:"datetime"`
}

type WalletSearchParams struct {
	StartDatetime string
	EndDatetime   string
}

type ListRepository struct {
	db *mongo.Database
}

func (l *ListRepository) List(params *WalletSearchParams) ([]*Wallet, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	opts := options.Find().SetSort(bson.D{{Key: "datetime", Value: -1}})
	filter := bson.D{}

	if params.StartDatetime != "" {
		filter = append(filter, bson.E{Key: "datetime", Value: bson.M{"$gte": params.StartDatetime}})
	}

	if params.EndDatetime != "" {
		filter = append(filter, bson.E{Key: "datetime", Value: bson.M{"$lte": params.EndDatetime}})
	}

	cursor, err := l.db.Collection("wallet").Find(ctx, filter, opts)
	var wallets []*Wallet

	if err != nil {
		log.Println("Error: DB Error |", time.Now().UTC().Format(time.RFC3339), err)
		return nil, errors.New("Server Error")
	}

	if err = cursor.All(context.TODO(), &wallets); err != nil {
		panic(err)
	}

	return wallets, nil
}

func NewListRepository(db *mongo.Database) *ListRepository {
	return &ListRepository{db}
}
