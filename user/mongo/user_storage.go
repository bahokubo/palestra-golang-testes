package mongo

import (
	"context"
	"fmt"
	"log"
	"user-crud/user"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		users[i].ID = uuid.New().String()
		_, err := us.collection.InsertOne(us.ctx, u)
		if err != nil {
			log.Println(fmt.Errorf("[Repository] Create user error: %v", err))
			return nil, err
		}

	}
	log.Println("[Repository] Create() succeeded")
	return users, nil
}

func (us *UserStorage) Update(user *user.User) (*user.User, error) {
	if err := us.collection.FindOneAndUpdate(
		us.ctx,
		bson.D{{"_id", user.ID}},
		bson.D{{"$set", user}},
		options.FindOneAndUpdate().SetReturnDocument(1)).
		Decode(&user); err != nil {
		log.Println(fmt.Errorf("[Repository] Update user Decode error: %v for UserId %s", err, user.ID))
		return nil, err
	}

	log.Println("[Repository] Update() succeeded")
	return user, nil
}

func (us *UserStorage) Delete(users []*user.User) ([]*user.User, error) {
	for _, u := range users {
		_, err := us.collection.InsertOne(us.ctx, u)
		if err != nil {
			log.Println(fmt.Errorf("[Repository] Create user error: %v", err))
			return nil, err
		}
	}
	log.Println("[Repository] Create() succeeded")
	return users, nil
}

func (us *UserStorage) List() ([]*user.User, error) {
	logData := map[string]string{}

	cursor, err := us.collection.Find(us.ctx, bson.M{
		"deletedAt": bson.M{"$exists": false},
	})

	if err != nil {
		log.Println(fmt.Errorf("[Repository] Create user error: %v", err))
		return nil, err
	}

	return us.listUsers(cursor, logData)
}

func (us *UserStorage) listUsers(c *mongo.Cursor, logData map[string]string) (users []*user.User, err error) {
	var user *user.User
	if err := c.All(us.ctx, &user); err != nil {
		log.Println(fmt.Errorf("[Repository] Create user error: %v", err))
		return nil, err
	}

	for _, u := range users {
		users = append(users, u)
	}

	return users, nil
}
