package data

import "time"

type Url struct {
	ID        int64     `json:"id" gorm:"primary_key;auto_increment"`
	ShortLink string    `json:"short_link"`
	FullLink  string    `json:"full_link"`
	Cliks     int64     `json:"cliks"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
