package repository

import (
	"context"

	"github.com/samuelrey/spot-the-bot/message"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type IProvider interface {
	GetRotationRepository() message.IRotationRepository
}

type Provider struct {
	database           *mongo.Database
	rotationRepository message.IRotationRepository
}

func NewProvider(mongoURI string) (IProvider, error) {
	dbOpt := options.Client().ApplyURI(mongoURI)
	mongoClient, err := mongo.Connect(context.TODO(), dbOpt)
	if err != nil {
		return nil, err
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	database := mongoClient.Database("spot-the-bot")
	rotationCollection := database.Collection("rotations")
	provider := Provider{
		database:           database,
		rotationRepository: message.NewRotationRepository(rotationCollection),
	}
	return &provider, nil
}

func (rp *Provider) GetRotationRepository() message.IRotationRepository {
	return rp.rotationRepository
}
