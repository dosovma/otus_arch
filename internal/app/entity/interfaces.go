package entity

//go:generate mockgen -destination=./mocks/executable.go -package=mocks -source=interfaces.go
type Executable interface {
	Execute() error
}
