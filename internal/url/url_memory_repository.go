package url

type urlRepositoryMemory struct {
	urls map[string]*Url
}

func NewRepositoryMemory() *urlRepositoryMemory {
	return &urlRepositoryMemory{make(map[string]*Url)}
}

func (r *urlRepositoryMemory) IsID(id string) bool {
	_, exist := r.urls[id]

	return exist
}

func (r *urlRepositoryMemory) FindByID(id string) (*Url, bool) {
	url, exist := r.urls[id]

	return url, exist
}

func (r *urlRepositoryMemory) FindByUrl(url string) *Url {
	for _, u := range r.urls {
		if u.Original == url {
			return u
		}
	}

	return nil
}

func (r *urlRepositoryMemory) Save(url Url) error {

	r.urls[url.ID] = &url

	return nil
}
