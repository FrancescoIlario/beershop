package beershop

import "context"

type Backend interface {
	Create(context.Context, CreateBeerCmd) (*CreateBeerCmdResult, error)
	Delete(context.Context, DeleteBeerCmd) (*DeleteBeerCmdResult, error)
	Read(context.Context, ReadBeerQry) (*ReadBeerQryResult, error)
	List(context.Context, ListBeerQry) (*ListBeerQryResult, error)
}

type backend struct {
	repo   Repository
	create CreateBeerHandlerFunc
	delete DeleteBeerHandlerFunc
	read   ReadBeerHandlerFunc
	list   ListBeerHandlerFunc
}

func NewBackend(repo Repository) Backend {
	return &backend{
		repo:   repo,
		create: NewCreateBeerHandler(repo),
		delete: NewDeleteBeerHandler(repo),
		read:   NewReadBeerHandler(repo),
		list:   NewListBeerHandler(repo),
	}
}

func (b *backend) Create(ctx context.Context, cmd CreateBeerCmd) (*CreateBeerCmdResult, error) {
	return b.create(ctx, cmd)
}

func (b *backend) Delete(ctx context.Context, cmd DeleteBeerCmd) (*DeleteBeerCmdResult, error) {
	return b.delete(ctx, cmd)
}

func (b *backend) Read(ctx context.Context, qry ReadBeerQry) (*ReadBeerQryResult, error) {
	return b.read(ctx, qry)
}

func (b *backend) List(ctx context.Context, qry ListBeerQry) (*ListBeerQryResult, error) {
	return b.List(ctx, qry)
}
