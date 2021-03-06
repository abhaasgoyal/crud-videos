package service

import (
	Config "crud-gin/config"
	"crud-gin/entity"
)

type VideoService interface {
	Save(entity.Video) error
	FindAll() ([]entity.Video, error)
	Update(entity.Video) error
	GetDetails(uint64) (entity.Video, error)
	Delete(uint64) error
}

type videoService struct {
	videos []entity.Video
}

func New() VideoService {
	return &videoService{}
}

func (service *videoService) Save(video entity.Video) error {
	return Config.DB.Create(video).Error
}

func (service *videoService) FindAll() ([]entity.Video, error) {
	var data []entity.Video
	err := Config.DB.Find(&data).Error
	return data, err
}

func (service *videoService) Update(video entity.Video) error {
	return Config.DB.Save(&video).Error
}
func (service *videoService) GetDetails(id uint64) (entity.Video, error) {
	var data entity.Video
	err := Config.DB.Take(&data, id).Error
	return data, err
}
func (service *videoService) Delete(id uint64) error {
	return Config.DB.Delete(&entity.Video{}, id).Error
}
