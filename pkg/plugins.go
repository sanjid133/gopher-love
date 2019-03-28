package pkg

import (
	"context"
	"fmt"
)

type PlatformFactory func(ctx context.Context) (Love, error)
type ManagerFactory func(ctx context.Context) LoveBag

var (
	platform = make(map[string]PlatformFactory)
	manager  = make(map[string]ManagerFactory)
)

func RegistarPlatform(name string, p PlatformFactory) {
	if _, found := platform[name]; found {
		fmt.Errorf("Platform %v already exists", name)
	}
	platform[name] = p
}

func RegistarManager(name string, m ManagerFactory) {
	if _, found := manager[name]; found {
		fmt.Errorf("Manager %v already exists", name)
	}
	manager[name] = m
}

func GetAllRegistereredManager() []string {
	managers := []string{}
	for name := range manager {
		managers = append(managers, name)
	}
	return managers
}

func GetPlatform(name string, ctx context.Context) (Love, error) {
	p, found := platform[name]
	if !found {
		return nil, fmt.Errorf("platform %v not found", name)
	}
	return p(ctx)
}

func GetManager(name string, ctx context.Context) (LoveBag, error) {
	m, found := manager[name]
	if !found {
		return nil, fmt.Errorf("platform %v not found", name)
	}
	return m(ctx), nil

}
