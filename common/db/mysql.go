package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"time"
)

type MysqlOpt struct {
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

func NewMysqlWithOpt(conn string, opt MysqlOpt) (eng *xorm.Engine, err error) {
	if conn == "" {
		panic("conn must be not empty")
	}

	return newMysql(conn, opt)
}

func NewMysql(conn string) (eng *xorm.Engine, err error) {
	if conn == "" {
		panic("conn must be not empty")
	}

	opt := MysqlOpt{
		MaxOpenConns:    50,
		MaxIdleConns:    10,
		ConnMaxLifetime: 10 * time.Second,
	}
	return newMysql(conn, opt)
}

func newMysql(conn string, opt MysqlOpt) (eng *xorm.Engine, err error) {
	eng, err = xorm.NewEngine("mysql", conn)
	if err != nil {
		return
	}

	if opt.MaxOpenConns > 0 {
		eng.SetMaxOpenConns(opt.MaxOpenConns)
	}

	if opt.MaxIdleConns > 0 {
		eng.SetMaxOpenConns(opt.MaxOpenConns)
	}

	if opt.ConnMaxLifetime > 0 {
		eng.SetConnMaxLifetime(opt.ConnMaxLifetime)
	}

	return eng, nil
}
