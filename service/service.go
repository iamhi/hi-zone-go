package service

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type RouterSetup func(*gin.Engine)

func StartService(setup_routers RouterSetup) bool {
	gin_engine := gin.Default()
	gin_engine.Use(gin.Recovery())
	gin_engine.Use(gin.Logger())

	gin_engine.Use(cors.New(cors.Config{
		AllowCredentials: true,
		AllowHeaders:     []string{"Content-Type", "Accept", "User-Agent"},
		AllowMethods:     []string{"POST", "GET", "DELETE"},
		AllowOrigins:     []string{"*"},
		AllowWildcard:    true,
	}))

	setup_routers(gin_engine)

	gin_engine.Run("localhost:8083")

	fmt.Println("Server started")

	return true
}
