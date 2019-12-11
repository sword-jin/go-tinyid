package core

import (
	"github.com/rrylee/go-tinyid/core/entity"
	"github.com/rrylee/go-tinyid/internal"
	"sync"
)

type SegmentIdService interface {
	GetNextSegmentId(bizType string) (*entity.SegmentId, error)
}

type CacheIdGenerator struct {
	bizType          string
	segmentIdService SegmentIdService
	current          *entity.SegmentId
	next             *entity.SegmentId
	isLoadingNext    bool
	mu               sync.Mutex
	nextMu           sync.Mutex
}

func (generator *CacheIdGenerator) loadCurrent() error {
	internal.Logf("loadCurrent start")
	generator.mu.Lock()
	if generator.current == nil || !generator.current.Useful() {
		if generator.next != nil {
			generator.current = generator.next
			generator.next = nil
		} else {
			segmentId, err := generator.querySegmentId()
			if err != nil {
				generator.mu.Unlock()
				internal.Warnf("loadCurrent err. err=%v", err)
				return err
			} else {
				generator.current = segmentId
			}
		}
	}
	internal.Logf("loadCurrent end")
	generator.mu.Unlock()
	return nil
}

func (generator *CacheIdGenerator) NextId() (int64, error) {
	for {
		if generator.current == nil {
			err := generator.loadCurrent()
			if err != nil {
				return 0, err
			}
			continue
		}
		result := generator.current.NextId()
		if result.Code == entity.Over {
			err := generator.loadCurrent()
			if err != nil {
				return 0, err
			}
		} else {
			if result.Code == entity.NeedLoad {
				generator.loadNext()
			}
			return result.Id, nil
		}
	}
}

func (generator *CacheIdGenerator) GetCurrentSegmentId(bizType string) *entity.SegmentId {
	return generator.current
}

func (generator *CacheIdGenerator) GetNextSegmentId(bizType string) *entity.SegmentId {
	return generator.next
}

func (generator *CacheIdGenerator) NextBatchIds(size int64) ([]int64, error) {
	ids := make([]int64, size)
	var i int64
	var err error
	for i = 0; i < size; i++ {
		ids[i], err = generator.NextId()
		if err != nil {
			return ids, err
		}
	}
	return ids, nil
}

func (generator *CacheIdGenerator) querySegmentId() (*entity.SegmentId, error) {
	return generator.segmentIdService.GetNextSegmentId(generator.bizType)
}

func (generator *CacheIdGenerator) loadNext() {
	if generator.next == nil && !generator.isLoadingNext {
		generator.nextMu.Lock()
		if generator.next == nil && !generator.isLoadingNext {
			internal.Logf("loadNext start")
			generator.isLoadingNext = true

			go func() {
				segmentId, err := generator.querySegmentId()
				if err != nil {
					internal.Warnf("loadNext err. err=%v", err)
				} else {
					generator.next = segmentId
					generator.isLoadingNext = false
					internal.Logf("loadNext success")
				}
			}()
		}
		generator.nextMu.Unlock()
	}
}

func NewCacheIdGenerator(bizType string, segmentIdService SegmentIdService) (*CacheIdGenerator, error) {
	generator := &CacheIdGenerator{bizType: bizType, segmentIdService: segmentIdService}
	err := generator.loadCurrent()
	if err != nil {
		return generator, err
	}
	return generator, nil
}
