package controller

import (
	"crud-gin/entity"
	"crud-gin/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(ctx *gin.Context) error
	FindAll() ([]entity.Video, error)
	Update(ctx *gin.Context) error
	GetDetails(ctx *gin.Context) (entity.Video, error)
	Delete(ctx *gin.Context) error
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() ([]entity.Video, error) {
	return c.service.FindAll()
	// if err != nil {
	// 	return entity.Video{}
	// }
	// return query				/
}
func (c *controller) Save(ctx *gin.Context) error {
	var video entity.Video

	if err := ctx.ShouldBindJSON(&video); err != nil {
		return err
	}
	c.service.Save(video)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var video entity.Video

	if err := ctx.ShouldBindJSON(&video); err != nil {
		return err
	}
	c.service.Update(video)
	return nil
}
func (c *controller) GetDetails(ctx *gin.Context) (entity.Video, error) {
	result, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	var video entity.Video
	if err != nil {
		return video, err
	}
	return c.service.GetDetails(result)
}

func (c *controller) Delete(ctx *gin.Context) error {
	result, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	return c.service.Delete(result)
}
