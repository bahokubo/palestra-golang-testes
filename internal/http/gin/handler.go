package gin

import (
	"context"
	"fmt"
	"net/http"
	"user-crud/config"
	"user-crud/internal/mongo"
	"user-crud/user"
	userRepository "user-crud/user/mongo"

	"github.com/gin-gonic/gin"
)

func Handlers(envs *config.Environments) *gin.Engine {
	r := gin.Default()

	ctx := context.Background()
	dbConn, _ := mongo.Open(envs.MongoAddress)
	db := dbConn.Database(envs.DBName)
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
