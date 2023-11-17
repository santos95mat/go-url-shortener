package url

import "time"

type Url struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Original  string    `json:"original"`
	Click     int       `json:"click"`
}

type CreateUrl struct {
	Original string `json:"original"`
}

type UrlRepository interface {
	IsID(id string) bool
	FindByID(id string) (*Url, bool)
	FindByUrl(url string) *Url
	Save(url Url) error
}

type UrlUsecase interface {
	SearchOrCreateNewUrl(originalUrl string) (u *Url, new bool, err error)
	Find(id string) (*Url, bool)
	Status(id string) (*Url, bool)
}
