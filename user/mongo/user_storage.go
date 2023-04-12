package mongo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
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

	for _, u := range users {
		_, err := us.collection.InsertOne(us.ctx, u)
		if err != nil {
			fmt.Sprintf("[Repository] Create user error: %v", err)
			return nil, err
		}
	}

	return users, nil
}

func (us *UserStorage) List() (users []*user.User, err error) {
	ctx := context.Background()
	//opts := options.FindOptions{}
	filter := bson.M{}

	cursor := us.collection.FindOne(ctx, filter)

	if err != nil {
		return users, nil
	}
	//
	//for cursor.Next(ctx) {
	var u user.User
	if err = cursor.Decode(&u); err != nil {
		return nil, err
	}
	users = append(users, &u)
	//}

	return users, nil
}
