package entity

type ResultCode int8

const (
	Over ResultCode = iota + 1
	NeedLoad
	Normal
)

type IdResult struct {
	Id   int64
	Code ResultCode
}
