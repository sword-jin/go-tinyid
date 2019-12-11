package core

import (
	"github.com/rrylee/go-tinyid/core/entity"
	"testing"
)

func TestCacheIdGenerator_NextId(t *testing.T) {
	generator := &CacheIdGenerator{
		bizType:          "test",
		segmentIdService: newTestService(),
	}
	t.Run("测试获取下一个 id", func(t *testing.T) {
		if id, _ := generator.NextId(); id != 5 {
			t.Errorf("下一个id是5，实际为%d", id)
		}
	})
	t.Run("测试获取下一组 id", func(t *testing.T) {
		excepted := []int64{10, 15}
		if ids, _ := generator.NextBatchIds(2); ids[0] != excepted[0] || ids[1] != excepted[1] {
			t.Errorf("下一组id是[10, 15]，实际为%v", ids)
		}
	})
}

type TestService struct{}

func (t TestService) GetNextSegmentId(bizType string) (*entity.SegmentId, error) {
	return entity.NewSegmentId(1000, 1, 5, 0, 800), nil
}

func newTestService() SegmentIdService {
	return &TestService{}
}
