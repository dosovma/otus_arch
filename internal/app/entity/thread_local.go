package entity

import (
	"sync"
)

const defaultScopeID = "default"

//go:generate mockgen -destination=./mocks/thread_local.go -package=mocks -source=thread_local.go
type IThreadLocal interface {
	GetCurrentScope() map[string]func(obj UObject) Executable
	SetCurrentScope(scope map[string]func(obj UObject) Executable)
	GetValue(key string) (map[string]func(obj UObject) Executable, bool)
	SetValue(key string, value map[string]func(obj UObject) Executable)
}

type ThreadLocal struct {
	scope  map[string]func(obj UObject) Executable
	scopes map[string]map[string]func(obj UObject) Executable
	mu     *sync.Mutex
}

func NewThreadLocal() *ThreadLocal {
	scope := make(map[string]func(obj UObject) Executable)
	scopes := make(map[string]map[string]func(obj UObject) Executable)
	scopes[defaultScopeID] = scope

	return &ThreadLocal{
		scope:  scope,
		scopes: scopes,
		mu:     &sync.Mutex{},
	}
}

func (tl *ThreadLocal) GetValue(key string) (map[string]func(obj UObject) Executable, bool) {
	tl.mu.Lock()
	scope, ok := tl.scopes[key]
	tl.mu.Unlock()

	return scope, ok
}

func (tl *ThreadLocal) SetValue(key string, value map[string]func(obj UObject) Executable) {
	tl.mu.Lock()
	tl.scopes[key] = value
	tl.mu.Unlock()
}

func (tl *ThreadLocal) GetCurrentScope() map[string]func(obj UObject) Executable {
	return tl.scope
}

func (tl *ThreadLocal) SetCurrentScope(scope map[string]func(obj UObject) Executable) {
	tl.scope = scope
}
