package e

// 查询成功  20000开头
// 服务器出错 50000开头
// 无效请求  40000开头
// 重定向  30000开头

var (
	SUCCESS = newError(200, "Success")
)

var (
	INVALID_REQUEST = newError(40001, "无效的请求")
	NOTFOUND        = newError(40002, "目标不存在")
	HAS_EXIST       = newError(40003, "目标已存在")
	LOGIN_ERROR     = newError(40004, "d登陆失败")
	UNTHORIZATION   = newError(40005, "鉴权失败")
)

var (
	INTERNAL_SERVER_ERROR = newError(50001, "内部服务错误")
)
