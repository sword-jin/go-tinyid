package mysql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rrylee/go-tinyid/core/util"
)

var dbs []*sql.DB

func GetConn() (db *sql.DB) {
	n := len(dbs)
	if n == 0 {
		panic("init mysql first.")
	} else if n == 1 {
		db = dbs[0]
	} else {
		randomIndex := util.RandomInt(n)
		db = dbs[randomIndex]
	}
	return
}

func Init(dsnes []string) error {
	dbs = make([]*sql.DB, len(dsnes))
	for index, dsn := range dsnes {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			return nil
		}
		dbs[index] = db
	}
	return nil
}
