package message

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IRotationRepository interface {
	RotationUpserter
	RotationFinder
}

type RotationUpserter interface {
	Upsert(Rotation) error
}

type RotationFinder interface {
	FindOne(string) (*Rotation, error)
}

type RotationRepository struct {
	collection *mongo.Collection
	context    context.Context
}

func NewRotationRepository(collection *mongo.Collection) *RotationRepository {
	return &RotationRepository{collection, context.TODO()}
}

func (repository *RotationRepository) Upsert(rotation Rotation) error {
	opts := options.Update().SetUpsert(true)
	filter := createFilterForServerID(rotation.ServerID)
	update := bson.D{
		primitive.E{
			Key: "$set",
			Value: bson.D{
				primitive.E{Key: "queue", Value: rotation.Queue},
				primitive.E{Key: "updated_at", Value: time.Now()},
			},
		},
	}
	_, err := repository.collection.UpdateOne(repository.context, filter, update, opts)

	return err
}

func (repository *RotationRepository) FindOne(serverID string) (*Rotation, error) {
	var rotation Rotation
	filter := createFilterForServerID(serverID)
	err := repository.collection.FindOne(repository.context, filter).Decode(&rotation)
	if err != nil {
		return nil, err
	}

	return &rotation, nil
}

func createFilterForServerID(serverID string) bson.D {
	return bson.D{primitive.E{Key: "server_id", Value: serverID}}
}
