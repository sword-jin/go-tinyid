package dao

import (
	"fmt"
	"github.com/rrylee/go-tinyid/internal"
	"github.com/rrylee/go-tinyid/server/dao/entity"
	"github.com/rrylee/go-tinyid/server/dbconnection/mysql"
	"log"
	"os"
	"testing"
)

func TestQueryByBizType(t *testing.T) {
	mysql.Init([]string{"root:root@tcp(127.0.0.1:3306)/test"})
	got, err := QueryByBizType(mysql.GetConn(), "test")
	if err != nil {
		t.Errorf("出现错误, err=%v", err)
	}

	actural := &entity.TinyIdInfo{
		Id:        1,
		BizType:   "test",
		MaxId:     1,
		Step:      100000,
		Delta:     1,
		Remainder: 0,
		Version:   1,
	}

	if got.Version != actural.Version || got.Id != actural.Id || got.BizType != actural.BizType || got.MaxId != actural.MaxId || got.Step != actural.Step || got.Delta != actural.Step || got.Remainder != actural.Remainder {
		t.Errorf("查询数据不匹配")
	}
}

func TestUpdateMaxId(t *testing.T) {
	internal.Logger = log.New(os.Stdout, "[test]", 0)
	mysql.Init([]string{"root:root@tcp(127.0.0.1:3306)/test"})
	updatedNum := UpdateMaxId(mysql.GetConn(), 1, 100000, 1, 1)
	fmt.Println(updatedNum)
}
