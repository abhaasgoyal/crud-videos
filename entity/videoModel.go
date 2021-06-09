package entity

type Video struct {
	ID          uint64 `json:"id" gorm:"primary_key;auto_increment"`
	Title       string `json:"title" binding:"required" gorm:"type:varchar(100)"`
	Description string `json:"description" binding:"required" gorm:"type:varchar(200)"`
	URL         string `json:"url" binding:"required" gorm:"type:varchar(100)"`
}
