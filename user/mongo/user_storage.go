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

func NewUserStorage(client *mongo.Database, ctx context.Context) *UserStorage {
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
			fmt.Printf("[Repository] Create user error: %v", err)
			return nil, err
		}

	}
	log.Println("[Repository] Create() succeeded")
	return users, nil
}

func (us *UserStorage) Update(user *user.User) (*user.User, error) {
	if err := us.collection.FindOneAndUpdate(
		us.ctx,
		bson.D{{Key: "id", Value: user.ID}},
		bson.D{{Key: "$set", Value: user}},
		options.FindOneAndUpdate().SetReturnDocument(1)).
		Decode(&user); err != nil {
		log.Println(fmt.Errorf("[Repository] Update user Decode error: %v for UserId %s", err, user.ID))
		return nil, err
	}

	log.Println("[Repository] Update() succeeded")
	return user, nil
}

func (us *UserStorage) List() ([]*user.User, error) {
	logData := map[string]string{}

	cursor, err := us.collection.Find(us.ctx, bson.M{
		"deletedAt": bson.M{"$exists": false},
	})

	if err != nil {
		log.Println(fmt.Errorf("[Repository] List user error: %v", err))
		return nil, err
	}

	return us.listUsers(cursor, logData)
}

func (us *UserStorage) listUsers(c *mongo.Cursor, logData map[string]string) ([]*user.User, error) {
	var foundUsers []*user.User
	if err := c.All(us.ctx, &foundUsers); err != nil {
		log.Println(fmt.Errorf("[Repository] Create user error: %v", err))
		return nil, err
	}

	return foundUsers, nil
}

func (us *UserStorage) Delete(id string) (int, error) {
	fmt.Printf("[Repository] Delete user repository starting for id: %s", id)

	result, err := us.collection.DeleteOne(us.ctx, bson.M{"id": id})

	if err != nil {
		fmt.Printf("[Repository] Delete user repository error when trying to destroy: %v for id: %s", err, id)
		return 0, err
	}

	fmt.Printf("[Repository] Delete user succeeded for id %s", id)

	return int(result.DeletedCount), nil
}
