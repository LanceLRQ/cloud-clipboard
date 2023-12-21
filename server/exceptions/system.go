package exceptions

import "github.com/LanceLRQ/cloud-clipboard/server/server"

// Bad Request Errno Namespace: 4000xxx
// Server Errno Namespace: 5000xxx

var ParseJSONError = server.NewHTTPServerRequestError(4000000, "JSON数据解析错误")
var ParamsValidatorError = server.NewHTTPServerRequestError(4000001, "参数内容校验失败")
var ParseIdError = server.NewHTTPServerRequestError(4000002, "ID解析错误")

var InternalServerError = server.NewHTTPServerError(5000000, "内部服务器错误")
var MySQLError = server.NewHTTPServerError(5000001, "访问数据库服务失败")
