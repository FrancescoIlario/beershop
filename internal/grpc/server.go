package grpc

import (
	context "context"

	"github.com/FrancescoIlario/beershop"
	"github.com/google/uuid"
)

type server struct {
	repo beershop.Repository
}

// Creates a new Beer
func (s *server) Create(ctx context.Context, r *CreateBeerRequest) (*CreateBeerReply, error) {
	b := beershop.Beer{
		Name: r.Name,
		Abv:  r.Abv,
	}

	id, err := s.repo.Create(ctx, b)
	if err != nil {
		return nil, err
	}

	return &CreateBeerReply{
		Id: id.String(),
	}, nil
}

// Reads an Beer
func (s *server) Read(ctx context.Context, r *ReadBeerRequest) (*ReadBeerReply, error) {
	id, err := uuid.Parse(r.Id)
	if err != nil {
		return nil, err
	}

	b, err := s.repo.Read(ctx, id)
	if err != nil {
		return nil, err
	}

	return &ReadBeerReply{
		Beer: convert(b),
	}, nil
}

// Delete an Beer
func (s *server) Delete(ctx context.Context, r *DeleteBeerRequest) (*DeleteBeerReply, error) {
	id, err := uuid.Parse(r.Id)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Delete(ctx, id); err != nil {
		return nil, err
	}

	return &DeleteBeerReply{}, nil
}

// List an Beer
func (s *server) List(ctx context.Context, r *ListActivitiesRequest) (*ListActivitiesReply, error) {
	bb, err := s.repo.List(ctx)
	if err != nil {
		return nil, err
	}

	bbl := make([]*Beer, len(bb))
	for i, b := range bb {
		bbl[i] = convert(b)
	}
	return &ListActivitiesReply{Beers: bbl}, nil
}

func convert(b beershop.Beer) *Beer {
	return &Beer{
		Id:   b.ID.String(),
		Name: b.Name,
		Abv:  b.Abv,
	}
}
