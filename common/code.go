package common

type res struct {
	Code int
	Msg string
	Data interface{}
}

var OK *res
var ServerErr *res
var DBErr *res
var NotFound *res
var ParamWrong *res
var AccountNotFound *res
var AccountFound *res

func init() {

	// 正确请求
	OK = &res{1, "成功", nil}

	// 服务器状态 5
	ServerErr = &res{5001, "服务器错误", nil}
	DBErr = &res{5002, "数据库错误", nil}

	// 资源状态 4
	NotFound = &res{4001, "资源不存在", nil}

	// 数据校验 3
	ParamWrong = &res{ 3001, "参数不合法", nil}

	// 账户状态 2
	AccountNotFound = &res{2001, "未查找到该账户", nil}
	AccountFound = &res{2002, "信息已存在", nil}
}

func NewCode(res *res, data interface{}) *res{
	res.Data = data
	return res
}
