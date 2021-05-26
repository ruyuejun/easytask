package common

type Response struct {
	Code 	int						`json:"code"`
	Msg 	string					`json:"msg"`
	Err 	error					`json:"err"`
	Data 	interface{}				`json:"data"`
}

var OK 					*Response
var ServerError 		*Response
var NotFound 			*Response
var ParamWrong 			*Response
var AccountNotFound 	*Response
var AccountFound 		*Response

func init() {

	// 正确请求
	OK = &Response{1, "成功", nil, nil}

	// 服务器状态 5
	ServerError = &Response{5001, "服务器错误", nil, nil}

	// 资源状态 4
	NotFound = &Response{4001, "资源不存在", nil, nil}

	// 数据校验 3
	ParamWrong = &Response{ 3001, "参数不合法", nil, nil}

	// 账户状态 2
	AccountNotFound = &Response{2001, "未查找到该账户", nil, nil}
	AccountFound = &Response{2002, "信息已存在", nil, nil}
}

func NewResponse(res *Response, data interface{}) *Response{
	res.Data = data
	return res
}

func BuildResponse(resp *Response) {
	if resp.Err != nil {
		resp.Code = 5001
		resp.Msg = "服务端错误"
	}

	resp.Code = 1
	resp.Msg = "成功"
}