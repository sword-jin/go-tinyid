package client

import (
	"github.com/rrylee/go-tinyid/client/config"
	"github.com/rrylee/go-tinyid/client/factory"
)

type TinyIdClient struct {
}

func Init(cfg *config.Config) {
	factory.Init(cfg)
}

func(client *TinyIdClient) NextId(bizType string) (int64, error) {
	clientGenerator, err := factory.GetIdGenerator(bizType)
	if err != nil {
		return 0, err
	}
	return clientGenerator.NextId()
}

func(client *TinyIdClient) NextBatchIds(bizType string, batchSize int64) ([]int64, error) {
	clientGenerator, err := factory.GetIdGenerator(bizType)
	if err != nil {
		return nil, err
	}
	return clientGenerator.NextBatchIds(batchSize)
}
