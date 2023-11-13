package url

import "time"

type Url struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Original  string    `json:"original"`
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

type Headers map[string]string
