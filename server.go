package main

import (
	Config "crud-gin/config"
	"crud-gin/controller"
	"crud-gin/entity"
	"crud-gin/service"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

// Write log to files
func setupLogOutput() {
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
}

func main() {
	var err error

	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))

	if err != nil {
		fmt.Println("Status:", err)
	}

	defer Config.DB.Close()
	Config.DB.AutoMigrate(&entity.Video{})

	setupLogOutput()

	server := gin.Default()

	endpoint := server.Group("/videos")

	endpoint.GET("", func(ctx *gin.Context) {
		query, err := videoController.FindAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error})
		} else {
			ctx.JSON(http.StatusOK, query)
		}
	})

	endpoint.POST("", func(ctx *gin.Context) {
		err := videoController.Save(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Successfully inserted"})
		}

	})

	endpoint.GET(":id", func(ctx *gin.Context) {
		query, err := videoController.GetDetails(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error})
		} else {
			ctx.JSON(http.StatusOK, query)
		}
	})

	endpoint.PUT(":id", func(ctx *gin.Context) {
		err := videoController.Update(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Successfully updated"})
		}
	})

	endpoint.DELETE(":id", func(ctx *gin.Context) {
		err := videoController.Delete(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": error.Error})
		} else {
			ctx.JSON(http.StatusOK, gin.H{"message": "Successfully deleted"})
		}
	})

	server.Run(":8080")
}
