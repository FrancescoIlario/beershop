package inmem

import (
	"sync"

	"github.com/FrancescoIlario/beershop"
	"github.com/FrancescoIlario/beershop/internal/storage"
	"github.com/google/uuid"
)

type repo struct {
	store map[uuid.UUID]beershop.Beer
	mutex sync.RWMutex
}

func New() storage.Repository {
	return &repo{
		store: map[uuid.UUID]beershop.Beer{},
		mutex: sync.RWMutex{},
	}
}

func (r *repo) Create(b beershop.Beer) (uuid.UUID, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	id, err := uuid.NewUUID()
	if err != nil {
		return uuid.UUID{}, err
	}

	b.ID = id
	r.store[id] = b
	return id, nil
}

func (r *repo) Delete(id uuid.UUID) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()

	for _, b := range r.store {
		if b.ID == id {
			delete(r.store, id)
			return nil
		}
	}
	return storage.ErrNotFound
}

func (r *repo) List() ([]beershop.Beer, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	v := make([]beershop.Beer, len(r.store))
	count := 0
	for _, b := range r.store {
		v[count] = b
		count++
	}
	return v, nil
}

func (r *repo) Read(id uuid.UUID) (beershop.Beer, error) {
	r.mutex.RLock()
	defer r.mutex.RUnlock()

	b, ok := r.store[id]
	if !ok {
		return beershop.Beer{}, storage.ErrNotFound
	}
	return b, nil
}
