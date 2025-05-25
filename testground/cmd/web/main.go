package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/iamhi/hi-zone-go/dbclient"
	"github.com/iamhi/hi-zone-go/service"
)

func Router(gin_engine *gin.Engine) {
	fmt.Println("Hello we are setting up controller")
}

func main() {
	service.StartService(Router)
	database.ConnectDB("root", "password", "http://localhost:5432", "hi-zone-go-testground")
}
