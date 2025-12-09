package main

import (
	"go-http/internal/api"
	"go-http/internal/constants"

	"github.com/gin-gonic/gin"
)

func main() {
	constants.InitConstants()
	r := gin.Default()
	api.DefineRoutes(r)
}
