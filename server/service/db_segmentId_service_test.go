package service

import (
	"fmt"
	"github.com/rrylee/go-tinyid/server/dbconnection/mysql"
	"testing"
)

func TestDbSegmentIdService_GetNextSegmentId(t *testing.T) {
	mysql.Init([]string{"root:root@tcp(127.0.0.1:3306)/test"})
	dbSegmentIdService := DbSegmentIdService{}
	for i := 0; i < 1; i++ {
		fmt.Println(dbSegmentIdService.GetNextSegmentId("test"))
	}
}
