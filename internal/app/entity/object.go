package entity

//go:generate mockgen -destination=./mocks/object.go -package=mocks -source=object.go
type UObject interface {
	GetProperty(name string) (any, bool)
	SetProperty(name string, value any) bool
}

type Object struct{}

func (o Object) GetProperty(string) (any, bool) {
	return false, false
}

func (o Object) SetProperty(string, any) bool {
	return false
}
