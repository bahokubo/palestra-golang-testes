package main

import (
	"fmt"
	"user-crud/config"
	"user-crud/internal/http/gin"
)

func main() {
	envs := config.LoadEnvVars()
	r := gin.Handlers(envs)
	err := r.Run()
	if err != nil {
		fmt.Errorf("error fatal")
	}
}
