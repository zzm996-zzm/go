package dao

import (
	"github.com/pkg/errors"
	"zuoye/code"
)

func GetUser() error {
	sql := "select * from user where id = 1"

 	//执行 sql
	//模拟异常
	err := errors.New("record not found")

	return errors.Wrapf(code.NotFound,"sql:%s - %v",sql,err)
}
