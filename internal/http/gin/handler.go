package gin

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"user-crud/config"
	"user-crud/internal/mongo"
	"user-crud/user"
	userRepository "user-crud/user/mongo"
)

func Handlers(envs *config.Environments) *gin.Engine {
	r := gin.Default()

	ctx := context.Background()
	dbcConn, _ := mongo.Open(envs.MongoAddress)
	db := dbcConn.Database(envs.DBName)
	userRepository := userRepository.New(db, ctx)

	us := user.NewService(userRepository)

	r.GET("/health", HealthHandler)
	fmt.Print(us)
	v1 := r.Group("/api/v1")

	MakeProductHandler(v1, us)

	return r
}

func HealthHandler(c *gin.Context) {
	c.String(http.StatusOK, "App is heath")
}
