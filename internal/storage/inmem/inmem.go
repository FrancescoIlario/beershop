package inmem

import (
	"sync"

	"github.com/FrancescoIlario/beershop/internal/domain"
	"github.com/FrancescoIlario/beershop/internal/storage"
	"github.com/google/uuid"
)

type repo struct {
	store map[uuid.UUID]domain.Beer
	mutex sync.RWMutex
}

func New() storage.Repository {
	return &repo{
		store: map[uuid.UUID]domain.Beer{},
		mutex: sync.RWMutex{},
	}
}

func (r *repo) Create(b domain.Beer) (uuid.UUID, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, err
	}

	b.Id = id
	r.store[id] = b
	return id, nil
}

func (r *repo) Delete(id uuid.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, b := range r.store {
		if b.Id == id {
			delete(r.store, id)
			return nil
		}
	}
	return storage.ErrNotFound
}

func (r *repo) List() ([]domain.Beer, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	v := make([]domain.Beer, len(r.store))
	count := 0
	for _, b := range r.store {
		v[count] = b
		count++
	}
	return v, nil
}

func (r *repo) Read(id uuid.UUID) (domain.Beer, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	b, ok := r.store[id]
	if !ok {
		return domain.Beer{}, storage.ErrNotFound
	}
	return b, nil
}
