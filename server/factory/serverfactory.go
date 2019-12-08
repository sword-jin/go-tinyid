package factory

import (
	"github.com/rrylee/go-tinyid/core"
	"github.com/rrylee/go-tinyid/server/service"
)

var serverIdFactory *core.IdGeneratorFactory

func init() {
	serverIdFactory = core.NewIdGeneratorFactory(func(bizType string) (core.IdGenerator, error) {
		return core.NewCacheIdGenerator(bizType, &service.DbSegmentIdService{})
	})
}

func GetIdGenerator(bizType string) (core.IdGenerator, error) {
	return serverIdFactory.GetGenerator(bizType)
}
