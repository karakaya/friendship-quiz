package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repository interface {
	Create(ctx context.Context, obj interface{}) (interface{}, error)
	Get(ctx context.Context, id uuid.UUID) (interface{}, error)
}

type repository struct {
	client     *mongo.Client
	collection *mongo.Collection
}

func NewRepository(mongoClient *mongo.Client, collection *mongo.Collection) Repository {
	return &repository{
		client:     mongoClient,
		collection: collection,
	}
}

func (r *repository) Create(ctx context.Context, obj interface{}) (interface{}, error) {
	result, err := r.collection.InsertOne(ctx, obj)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return result.InsertedID, nil
}

func (r *repository) Get(ctx context.Context, id uuid.UUID) (interface{}, error) {
	var obj interface{}
	r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&obj)

	return obj, nil
}
