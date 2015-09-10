package persistence

func NewInMemoryRepository() InMemoryRepository {
	return InMemoryRepository{
		store: make(map[string]interface{}),
	}
}

type InMemoryRepository struct {
	store map[string]interface{}
}

func (i InMemoryRepository) Load(key string) (interface{}, error) {
	res, ok := i.store[key]
	if ok {
		return res, nil
	}
	return nil, NotFoundErr
}

func (i InMemoryRepository) Store(key string, data interface{}) error {
	i.store[key] = data
	return nil
}
