package dao

import (
	"database/sql"
	"fmt"
	"github.com/rrylee/go-tinyid/internal"
	"github.com/rrylee/go-tinyid/server/dao/entity"
)

const (
	SelectSql = "select id, biz_type, max_id, step, delta, remainder, create_time, update_time, version from tiny_id_info where biz_type = '%s'"
	UpdateSql = "update tiny_id_info set max_id= ?," +
		" update_time=now(), version=version+1" +
		" where id=? and max_id=? and version=?"
)

func QueryByBizType(db *sql.DB, bizType string) (*entity.TinyIdInfo, error) {
	tinyIdInfo := &entity.TinyIdInfo{}
	selectSql := fmt.Sprintf(SelectSql, bizType)
	err := db.QueryRow(selectSql, &tinyIdInfo.Id, &tinyIdInfo.BizType, &tinyIdInfo.MaxId, &tinyIdInfo.Step, &tinyIdInfo.Delta, &tinyIdInfo.Remainder, &tinyIdInfo.CreateTime, &tinyIdInfo.UpdateTime, &tinyIdInfo.Version)
	if err != nil {
		return nil, fmt.Errorf("QueryByBizType query error. err=%v||bizType=%s", err, bizType)
	}
	return nil, fmt.Errorf("QueryByBizType bizType not in db. bizType=%s", bizType)
}

func UpdateMaxId(db *sql.DB, id, newMaxId, oldMaxId, version int64) int64 {
	stmt, err := db.Prepare(UpdateSql)
	if err != nil {
		internal.Warnf("UpdateMaxId prepare error. err=%v", err)
		return 0
	}

	ret, err := stmt.Exec(newMaxId, id, oldMaxId, version)
	if err != nil {
		internal.Warnf("UpdateMaxId exec error. err=%v", err)
		return 0
	}

	rowsAffected, err := ret.RowsAffected()
	if err != nil {
		internal.Warnf("UpdateMaxId RowsAffected error. err=%v", err)
	}
	return rowsAffected
}
