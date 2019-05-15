package config

/*
	这里没有使用 CODE["OK"] = code_obj{1, "ok"}  这样更简单直观的map结构
	因为map结构定义的服务状态在获取时，无法通过CODE["OK"] 直接获取，开发不方便
	采用繁琐的结构体，却能直接使用点语法快速获取
 */

type code struct {
	OK
	SerErr
	DBErr
	NotFound
	ParamWrong
	AccountNotFound
	AccountFound
}

var CODE code

func init() {

	// 正确请求
	CODE.OK = OK{1, "成功",}

	// 服务器状态 5
	CODE.SerErr = SerErr{5001, "服务器错误"}
	CODE.DBErr = DBErr{5002, "数据库错误"}

	// 资源状态 4
	CODE.NotFound = NotFound{4001, "资源不存在"}

	// 数据校验 3
	CODE.ParamWrong = ParamWrong{ 3001, "参数不合法"}

	// 账户状态 2
	CODE.AccountNotFound = AccountNotFound{2001, "未查找到该账户"}
	CODE.AccountFound = AccountFound{2002, "信息已存在"}
}

type AccountFound struct {
	Code int
	Msg string
}

type AccountNotFound struct {
	Code int
	Msg string
}

type NotFound struct {
	Code int
	Msg string
}

type ParamWrong struct {
	Code int
	Msg string
}

type DBErr struct {
	Code int
	Msg string
}

type SerErr struct {
	Code int
	Msg string
}

type OK struct {
	Code int
	Msg string
}


