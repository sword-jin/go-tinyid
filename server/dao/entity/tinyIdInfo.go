package entity

type TinyIdInfo struct {
	Id         int64
	BizType    string
	BeginId    int64
	MaxId      int64
	Step       int64
	Delta      int64
	Remainder  int64
	Version    int64
	CreateTime string
	UpdateTime string
}
