package service

import (
	"zuoye/code"
	"zuoye/dao"
	"github.com/pkg/errors"
)

type User struct{
	ID int
}
func GetUserById() (*User,error){

	err:=dao.GetUser();
	if errors.Is(err, code.NotFound) {
		//请教一个问题 这里的错误应该是返回 还是直接返回空 我记得错误只处理一次
		return nil,code.NotFound
}

	//其他业务...
	return nil,nil
}