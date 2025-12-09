package api

import (
	"go-http/internal/api/handlers"
	"go-http/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DefineRoutes(r *gin.Engine) {
	r.GET("/inventory", func(ctx *gin.Context) {
		pageStr := ctx.Query("page")
		var page int
		limit := 10
		var err error
		if pageStr != "" {
			page, err = strconv.Atoi(pageStr)
			if err != nil || page == 0 {
				page = 1
			}
		}
		limitStr := ctx.Query("limit")
		if limitStr != "" {
			limit, err = strconv.Atoi(limitStr)
		}
		items := handlers.FetchInventory(limit, page)
		ctx.Header("X-Total-Count", strconv.Itoa(len(items)))
		ctx.JSON(http.StatusOK, items)
	})
	r.POST("/inventory", func(ctx *gin.Context) {
		var data models.InventoryItem
		if err := ctx.ShouldBindBodyWithJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "The inventory item is not valid",
			})
			return
		}
		err := handlers.AddToInventory(&data)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Done",
		})

	})
}
