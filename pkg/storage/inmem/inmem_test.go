package inmem_test

import (
	"testing"

	"github.com/FrancescoIlario/beershop/pkg/domain"
	"github.com/FrancescoIlario/beershop/pkg/storage"
	"github.com/FrancescoIlario/beershop/pkg/storage/inmem"
	"github.com/google/uuid"
	"github.com/matryer/is"
)

func Test_create(t *testing.T) {
	// Arrange
	is := is.New(t)
	repo := inmem.New()
	b := domain.Beer{
		Name: "beer",
		Abv:  4.5,
	}

	// Act
	id, err := repo.Create(b)

	// Assert
	is.NoErr(err)
	is.True(id.Variant() != uuid.Invalid)
}

func Test_read_not_existing_beer(t *testing.T) {
	// Arrange
	is := is.New(t)
	repo := inmem.New()

	// Act
	_, err := repo.Read(uuid.New())

	// Assert
	is.Equal(err, storage.ErrNotFound)
}

func Test_read_existing_beer(t *testing.T) {
	// Arrange
	is := is.New(t)
	repo := inmem.New()
	b := domain.Beer{
		Name: "beer",
		Abv:  4.5,
	}

	id, err := repo.Create(b)
	is.NoErr(err)
	b.Id = id

	// Act
	br, err := repo.Read(id)

	// Assert
	is.Equal(b, br)
}

func Test_delete_not_existing_beer(t *testing.T) {
	// Arrange
	is := is.New(t)
	repo := inmem.New()

	// Act
	err := repo.Delete(uuid.New())

	// Assert
	is.Equal(err, storage.ErrNotFound)
}

func Test_delete_existing_beer(t *testing.T) {
	// Arrange
	is := is.New(t)
	repo := inmem.New()
	b := domain.Beer{
		Name: "beer",
		Abv:  4.5,
	}

	id, err := repo.Create(b)
	is.NoErr(err)

	// Act
	err = repo.Delete(id)

	// Assert
	is.NoErr(err)
	_, err = repo.Read(id)
	is.Equal(err, storage.ErrNotFound)
}

func Test_list_empty_beershop(t *testing.T) {
	// Arrange
	is := is.New(t)
	repo := inmem.New()

	// Act
	l, err := repo.List()

	// Assert
	is.NoErr(err)
	is.True(len(l) == 0)
}

func Test_list_beershop(t *testing.T) {
	// Arrange
	is := is.New(t)
	repo := inmem.New()
	bs := []domain.Beer{
		{
			Name: "beer 1",
			Abv:  4.5,
		},
		{
			Name: "beer 2",
			Abv:  5.0,
		},
	}
	bsm := make(map[uuid.UUID]domain.Beer, len(bs))

	for _, b := range bs {
		id, err := repo.Create(b)
		is.NoErr(err)
		b.Id = id
		bsm[id] = b
	}

	// Act
	l, err := repo.List()

	// Assert
	is.NoErr(err)
	is.True(len(l) == len(bs))
	for _, bl := range l {
		bm, ok := bsm[bl.Id]
		is.True(ok)
		is.Equal(bm, bl)
	}
}
