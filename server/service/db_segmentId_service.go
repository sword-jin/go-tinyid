package service

import (
	"fmt"
	"github.com/rrylee/go-tinyid/constant"
	coreEntity "github.com/rrylee/go-tinyid/core/entity"
	"github.com/rrylee/go-tinyid/internal"
	"github.com/rrylee/go-tinyid/server/dao"
	daoEntity "github.com/rrylee/go-tinyid/server/dao/entity"
	"github.com/rrylee/go-tinyid/server/dbconnection/mysql"
)

type DbSegmentIdService struct {
}

func (d DbSegmentIdService) GetNextSegmentId(bizType string) (*coreEntity.SegmentId, error) {
	for i := 0; i < constant.MaxRetryFromDB; i++ {
		db := mysql.GetConn()
		tinyIdInfo, err := dao.QueryByBizType(db, bizType)
		if err != nil {
			internal.Warnf("%s", err.Error())
			return nil, err
		}
		newMaxId := tinyIdInfo.MaxId + tinyIdInfo.Step
		oldMaxId := tinyIdInfo.MaxId
		updated := dao.UpdateMaxId(db, tinyIdInfo.Id, newMaxId, oldMaxId, tinyIdInfo.Version)
		if updated == 1 {
			tinyIdInfo.MaxId = newMaxId
			segmentId := convertSegmentId(tinyIdInfo)
			internal.Logf("GetNextSegmentId success. tinyIdInfo: %v, segmentId: %v", tinyIdInfo, segmentId)
			return segmentId, nil
		} else {
			internal.Warnf("GetNextSegmentId conflict. tinyIdInfo: %v", tinyIdInfo)
		}
	}
	return nil, fmt.Errorf("GetNextSegmentId confict, ret max time. retry=%d", constant.MaxRetryFromDB)
}

func convertSegmentId(info *daoEntity.TinyIdInfo) *coreEntity.SegmentId {
	currentId := info.MaxId - info.Step
	return coreEntity.NewSegmentId(info.MaxId, currentId, info.Delta, info.Remainder, currentId+info.Step*constant.LoadingPercent/100)
}
