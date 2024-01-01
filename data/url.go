package data

import "time"

type Url struct {
	ID        int64     `json:"id" gorm:"primary_key;auto_increment"`
	ShortLink string    `json:"short_link"`
	FullLink  string    `json:"full_link"`
	UniqueID  string    `json:"unique_id"`
	Clicks    int64     `json:"clicks"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *Url) CreateShortUrl() *Url {
	DB.Create(u)
	return u
}

func FindByShortLink(shortLink string) *Url {
	var url Url
	DB.First(&url, "short_link = ?", shortLink)
	return &url
}

func FindByUniqueId(uniqueId string) *Url {
	var url Url
	DB.First(&url, "unique_id = ?", uniqueId)
	return &url
}

func (u *Url) UpdateClicks() *Url {
	DB.Model(u).Update("clicks", u.Clicks)
	return u
}

func (u *Url) FindByFullLink() *Url {
	DB.First(u, "full_link = ?", u.FullLink)
	return u
}

func (u *Url) FindById() *Url {
	DB.First(u, "id = ?", u.ID)
	return u
}

func (u *Url) Delete() {
	DB.Delete(u)
}

func FindAll() []Url {
	var urls []Url
	DB.Find(&urls)
	return urls
}

func (u *Url) Update() *Url {
	DB.Save(u)
	return u
}
