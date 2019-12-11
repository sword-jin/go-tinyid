package core

import "sync"

type IdGeneratorFactory struct {
	mu              sync.Mutex
	generators      map[string]IdGenerator
	CreateGenerator func(bizType string) (IdGenerator, error)
}

func NewIdGeneratorFactory(createGenerator func(bizType string) (IdGenerator, error)) *IdGeneratorFactory {
	return &IdGeneratorFactory{
		CreateGenerator: createGenerator,
		generators:      make(map[string]IdGenerator),
	}
}

func (factory *IdGeneratorFactory) GetGenerator(bizType string) (IdGenerator, error) {
	if g, ok := factory.generators[bizType]; ok {
		return g, nil
	}
	factory.mu.Lock()
	defer factory.mu.Unlock()
	idGenerator, err := factory.CreateGenerator(bizType)
	if err != nil {
		return nil, err
	}
	if g, ok := factory.generators[bizType]; ok {
		return g, nil
	}
	factory.generators[bizType] = idGenerator
	return idGenerator, nil
}
