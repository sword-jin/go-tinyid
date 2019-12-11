package entity

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type SegmentId struct {
	mu        sync.Mutex
	haveInit  bool
	maxId     int64
	loadingId int64 // start pre loading id
	currentId int64
	delta     int64
	remainder int64
}

func (s *SegmentId) CurrentId() int64 {
	return s.currentId
}

func (s *SegmentId) Delta() int64 {
	return s.delta
}

func (s *SegmentId) Remainder() int64 {
	return s.remainder
}

func (s *SegmentId) LoadingId() int64 {
	return s.loadingId
}

func (s *SegmentId) MaxId() int64 {
	return s.maxId
}

func NewSegmentId(maxId, currentId, delta, remainder, loadingId int64) *SegmentId {
	return &SegmentId{
		maxId:     maxId,
		currentId: currentId,
		delta:     delta,
		remainder: remainder,
		loadingId: loadingId,
		haveInit:  false,
	}
}

func (s *SegmentId) NextId() *IdResult {
	s.init()
	id := atomic.AddInt64(&s.currentId, s.delta)
	id -= 1
	if id > s.maxId {
		return &IdResult{
			Id:   id,
			Code: Over,
		}
	}
	if id >= s.loadingId {
		return &IdResult{
			Id:   id,
			Code: NeedLoad,
		}
	}
	return &IdResult{Id: id, Code: Normal}
}

func (s *SegmentId) init() {
	if s.haveInit {
		return
	}
	s.mu.Lock()
	id := atomic.LoadInt64(&s.currentId)
	if id%s.delta == s.remainder {
		s.haveInit = true
		s.mu.Unlock()
		return
	}
	var i int64
	for i = 0; i <= s.delta; i++ {
		id = atomic.AddInt64(&s.currentId, 1)
		if id%s.delta == s.remainder {
			atomic.AddInt64(&s.currentId, -s.delta)
			s.haveInit = true
			s.mu.Unlock()
			return
		}
	}
}

func (s *SegmentId) Useful() bool {
	return atomic.LoadInt64(&s.currentId) <= s.maxId
}

func (s *SegmentId) String() string {
	return fmt.Sprintf("maxId=%d, loadingId=%d, currentId=%d, delta=%d, remainder=%d", s.maxId, s.loadingId, s.currentId, s.delta, s.remainder)
}

