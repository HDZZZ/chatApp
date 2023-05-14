package main

const REQUEST_PARAMS_USERNAME = "username"
const REQUEST_PARAMS_PASSWORD = "password"
const REQUEST_PARAMS_UID = "uid"
const REQUEST_HEADER_TOKEN = "token"

const ERROR_CODE_1001 = 1001
const ERROR_CODE_1002 = 1002
const ERROR_CODE_1003 = 1003
const ERROR_CODE_1004 = 1004
const ERROR_CODE_101 = 101

var errCode = map[int]string{
	ERROR_CODE_101:  "系统ip异常",
	ERROR_CODE_1001: "用户名与密码不能为空",
	ERROR_CODE_1002: "该用户名已被占用",
	ERROR_CODE_1003: "无此用户",
	ERROR_CODE_1004: "密码错误",
}
