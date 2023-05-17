package common

type ResultData struct {
	Code    int
	Data    interface{}
	Message string
}

func CreateResultDataSuccess(data interface{}) ResultData {
	return ResultData{
		1,
		data,
		"success"}
}

func CreateResultDataError(code int, message string) ResultData {
	return ResultData{
		code,
		"",
		message}
}
