package main

import (
	"log"
	"user-crud/config"
	http "user-crud/internal/http/gin"
	"user-crud/internal/mongo"
)

func main() {
	envs := config.LoadEnvVars()
	dbConn, _ := mongo.Open(envs.MongoAddress)
	db := dbConn.Database(envs.DBName)

	r := http.Handlers(db)
	err := r.Run(envs.APIPort)
	if err != nil {
		log.Fatalf("error fatal")
	}
}
