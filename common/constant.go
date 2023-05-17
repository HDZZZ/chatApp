package common

const REQUEST_PARAMS_USERNAME = "username"
const REQUEST_PARAMS_PASSWORD = "password"
const REQUEST_PARAMS_UID = "uid"
const REQUEST_PARAMS_RECEIVERUID = "receiverUid"
const REQUEST_PARAMS_REQUEST_MSG = "requestMsg"
const REQUEST_PARAMS_REQUEST_ID = "requestId"
const REQUEST_PARAMS_FRIEND_UID = "friendId"

const REQUEST_HEADER_TOKEN = "token"

const ERROR_CODE_1001 = 1001
const ERROR_CODE_1002 = 1002
const ERROR_CODE_1003 = 1003
const ERROR_CODE_1004 = 1004
const ERROR_CODE_101 = 101
const ERROR_CODE_102 = 102
const ERROR_CODE_103 = 103
const ERROR_CODE_104 = 104

const ERROR_CODE_2001 = 2001
const ERROR_CODE_2002 = 2002
const ERROR_CODE_2003 = 2003
const ERROR_CODE_2004 = 2004
const ERROR_CODE_2005 = 2005
const ERROR_CODE_2006 = 2006
const ERROR_CODE_2007 = 2007
const ERROR_CODE_2008 = 2008

var ErrCode = map[int]string{
	ERROR_CODE_101:  "系统ip异常",
	ERROR_CODE_102:  "没有收到预期的参数",
	ERROR_CODE_103:  "token无效",
	ERROR_CODE_104:  "你没有相关权限",
	ERROR_CODE_1001: "用户名与密码不能为空",
	ERROR_CODE_1002: "该用户名已被占用",
	ERROR_CODE_1003: "无此用户",
	ERROR_CODE_1004: "密码错误",
	ERROR_CODE_2001: "请求发送失败,",
	ERROR_CODE_2002: "同意无效,",
	ERROR_CODE_2003: "拒绝失败,",
	ERROR_CODE_2004: "没有该条请求",
	ERROR_CODE_2005: "无法对该条请求实行该操作",
	ERROR_CODE_2006: "你们不是好友,",
	ERROR_CODE_2007: "请求不能重复发送",
	ERROR_CODE_2008: "添加异常",
}

const RedisKeyUnSendMsgIds = "UnSendMsgIds"
