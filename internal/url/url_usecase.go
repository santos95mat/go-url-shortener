package url

import (
	"math/rand"
	"time"
)

type urlUsecase struct {
	urlRepository UrlRepository
}

func NewUrlUsecase(repo UrlRepository) *urlUsecase {
	return &urlUsecase{urlRepository: repo}
}

func (r *urlUsecase) SearchOrCreateNewUrl(originalUrl string) (u *Url, new bool, err error) {

	if u = r.urlRepository.FindByUrl(originalUrl); u != nil {
		return u, false, nil
	}

	id := r.generateID()

	url := Url{id, time.Now(), originalUrl, 0}
	r.urlRepository.Save(url)

	return &url, true, nil
}

func (r *urlUsecase) Find(id string) (*Url, bool) {
	url, exist := r.urlRepository.FindByID(id)

	if !exist {
		return url, exist
	}

	url.Click++
	r.urlRepository.Save(*url)

	return url, exist
}

func (r *urlUsecase) Status(id string) (*Url, bool) {
	url, exist := r.urlRepository.FindByID(id)

	return url, exist
}

func (r *urlUsecase) generateID() string {
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
