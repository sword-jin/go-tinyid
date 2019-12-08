package entity

import (
	"fmt"
	"testing"
)

func TestSegmentId_NextId(t *testing.T) {
	s := NewSegmentId(1000, 1, 3, 0, 800)
	fmt.Println(s.NextId())
	fmt.Println(s.NextId())
	fmt.Println(s.NextId())
	fmt.Println(s.NextId())
	fmt.Println(s.NextId())
}
