package gin

import (
	"context"
	"fmt"
	"net/http"
	"user-crud/user"
	userRepository "user-crud/user/mongo"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func Handlers(db *mongo.Database) *gin.Engine {
	r := gin.Default()

	ctx := context.Background()
	userRepository := userRepository.NewUserStorage(db, ctx)

	us := user.NewService(userRepository)

	r.GET("/health", HealthHandler)
	fmt.Print(us)
	v1 := r.Group("/api/v1")

	MakeUserHandler(v1, us)

	return r
}

func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "App is heath")
}
