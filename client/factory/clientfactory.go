package factory

import (
	"github.com/rrylee/go-tinyid/client/config"
	"github.com/rrylee/go-tinyid/client/service"
	"github.com/rrylee/go-tinyid/core"
)

var clientIdFactory *core.IdGeneratorFactory

func Init(cfg *config.Config) {
	clientIdFactory = core.NewIdGeneratorFactory(func(bizType string) (core.IdGenerator, error) {
		return core.NewCacheIdGenerator(bizType, &service.HttpSegmentIdService{Config: cfg})
	})
}

func GetIdGenerator(bizType string) (core.IdGenerator, error) {
	return clientIdFactory.GetGenerator(bizType)
}
