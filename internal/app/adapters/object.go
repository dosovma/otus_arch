package adapters

//go:generate mockgen -destination=./mocks/object.go -package=mocks -source=object.go
type UObject interface {
	GetProperty(name string) (any, bool)
	SetProperty(name string, value any) bool
}
