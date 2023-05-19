package common

const REQUEST_PARAMS_USERNAME = "username"
const REQUEST_PARAMS_PASSWORD = "password"
const REQUEST_PARAMS_UID = "uid"
const REQUEST_PARAMS_RECEIVERUID = "receiverUid"
const REQUEST_PARAMS_REQUEST_MSG = "requestMsg"
const REQUEST_PARAMS_REQUEST_ID = "requestId"
const REQUEST_PARAMS_FRIEND_UID = "friendId"

const REQUEST_PARAMS_MEMBER_UID = "membersUid"
const REQUEST_PARAMS_NEW_OWNER_UID = "newOwnerUid"
const REQUEST_PARAMS_GROUP_GID = "gid"
const REQUEST_PARAMS_GROUP_NAME = "groupName"
const REQUEST_PARAMS_GROUP_DESCIPTION = "description"
const REQUEST_PARAMS_GROUP_MEMBER_ID = "memberId"
const REQUEST_PARAMS_GROUP_MEMBER_ALIAS = "alias"

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

const ERROR_CODE_3001 = 3001
const ERROR_CODE_3002 = 3002
const ERROR_CODE_3003 = 3003
const ERROR_CODE_3004 = 3004
const ERROR_CODE_3005 = 3005
const ERROR_CODE_3006 = 3006
const ERROR_CODE_3007 = 3007
const ERROR_CODE_3008 = 3008

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
	ERROR_CODE_3001: "群聊创建失败",
	ERROR_CODE_3002: "退出群聊失败,",
	ERROR_CODE_3003: "您是群主,无法退出群聊,需要先转让群主身份,",
	ERROR_CODE_3004: "更改群资料必须传递相关资料",
	ERROR_CODE_3005: "更改群资料失败",
	ERROR_CODE_3006: "没有该群聊",
	ERROR_CODE_3007: "添加好友失败",
	ERROR_CODE_3008: "更改个人群成员资料失败",
}

const RedisKeyUnSendMsgIds = "UnSendMsgIds"
