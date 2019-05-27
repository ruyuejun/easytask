package code

type code struct {
	Code int
	Msg string
}

var OK code
var SerErr code
var DBErr code
var NotFound code
var ParamWrong code
var AccountNotFound code
var AccountFound code

func init() {

	// 正确请求
	OK = code{1, "成功",}

	// 服务器状态 5
	SerErr = code{5001, "服务器错误"}
	DBErr = code{5002, "数据库错误"}

	// 资源状态 4
	NotFound = code{4001, "资源不存在"}

	// 数据校验 3
	ParamWrong = code{ 3001, "参数不合法"}

	// 账户状态 2
	AccountNotFound = code{2001, "未查找到该账户"}
	AccountFound = code{2002, "信息已存在"}
}

