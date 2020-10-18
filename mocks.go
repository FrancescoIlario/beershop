package beershop

//go:generate mockgen -source storage.go -package=mocks -destination=internal/mocks/storage.go github.com/FrancescoIlario/beershop/storage.go Repository
//go:generate mockgen -source backend.go -package=mocks -destination=internal/mocks/backend.go github.com/FrancescoIlario/beershop/backend.go Backend
//go:generate mockgen -source validation.go -package=mocks -destination=internal/mocks/validation.go github.com/FrancescoIlario/beershop/validation.go ValidationResult
