package url

import (
	"math/rand"
	"time"
)

type UrlUsecase struct {
	urlRepository UrlRepository
}

func NewUrlUsecase(repo UrlRepository) *UrlUsecase {
	return &UrlUsecase{urlRepository: repo}
}

func (r *UrlUsecase) SearchOrCreateNewUrl(originalUrl string) (u *Url, new bool, err error) {

	if u = r.urlRepository.FindByUrl(originalUrl); u != nil {
		return u, false, nil
	}

	id := r.generateID()

	url := Url{id, time.Now(), originalUrl}
	r.urlRepository.Save(url)

	return &url, true, nil
}

func (r *UrlUsecase) Find(id string) (*Url, bool) {
	url, b := r.urlRepository.FindByID(id)

	return url, b
}

func (r *UrlUsecase) generateID() string {
	const (
		length  = 5
		simbols = "abcdefghijklmnopqrdtuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	)

	novoID := func() string {
		id := make([]byte, length, length)

		for i := range id {
			id[i] = simbols[rand.Intn(len(simbols))]
		}
		return string(id)
	}

	for {
		if id := novoID(); !r.urlRepository.IsID(id) {
			return id
		}
	}
}
