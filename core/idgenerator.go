package core

import "github.com/rrylee/go-tinyid/core/entity"

type IdGenerator interface {
	NextId() (int64, error)
	NextBatchIds(size int64) ([]int64, error)
	GetCurrentSegmentId(bizType string) *entity.SegmentId
	GetNextSegmentId(bizType string) *entity.SegmentId
}
