package ecode

// 设置错误码和响应码

const(
     SUCCESS = 10000
     ERROR = 10001
     Request_Params_ERROR = 10002
     PARSE_AUTH_CHECK_TOKEN_FAIl = 10003
     PARSE_AUTH_CHECK_TOKEN_TIMEOUT = 10004
     AUTH_TOKEN_ERROR = 10005
     AUTH_ERROR = 10006
     Email_PASSWORD_NULL_ERROR = 10007

     CreateNoteFailed = 10008
     EditNoteFailed = 10009
     DelNoteFailed = 10010
     QueryNoteFailed = 10011


)

var MsgStatus = map[int]string{
     SUCCESS: "success",
     ERROR: "failed",
     AUTH_ERROR: "授权出错",
     AUTH_TOKEN_ERROR: "授权Token出错",
     Request_Params_ERROR: "请求参数错误",
     PARSE_AUTH_CHECK_TOKEN_FAIl : "解析Token失败",
     PARSE_AUTH_CHECK_TOKEN_TIMEOUT :"解析Token超时",
     Email_PASSWORD_NULL_ERROR :"邮箱或密码或验证码不能为空",
     CreateNoteFailed:"创建帖子失败",
     EditNoteFailed : "修改帖子失败",
     DelNoteFailed : "删除帖子失败",
     QueryNoteFailed : "查询帖子失败",
}


func GetMsg(code int) string{
     msg, ok := MsgStatus[code]
     if ok {
          return msg
     }
     return MsgStatus[ERROR]
}