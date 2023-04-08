package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"user-crud/user"
)

const userCollection = "users"

type UserStorage struct {
	collection *mongo.Collection
	ctx        context.Context
}

func New(client *mongo.Database, ctx context.Context) *UserStorage {
	return &UserStorage{
		collection: client.Collection(userCollection),
		ctx:        ctx,
	}
}

func (us *UserStorage) Create(users []*user.User) ([]*user.User, error) {

	for i, u := range users {
		result, err := us.collection.InsertOne(us.ctx, u)
		if err != nil {
			fmt.Sprintf("[Repository] Create user error: %v", err)
			return nil, err
		}
		users[i].ID = result.InsertedID.(primitive.ObjectID).Hex()
	}

	return users, nil
}
