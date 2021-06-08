package entity

type Video struct {
	ID          uint64 `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"version"`
	URL         string `json:"url"`
}
