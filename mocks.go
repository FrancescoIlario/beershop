package beershop

//go:generate mockgen -source storage.go -package=mocks -destination=internal/mocks/storage.go github.com/FrancescoIlario/beershop/storage.go Repository
