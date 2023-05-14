package main

type ResultData struct {
	Code    int
	Data    interface{}
	Message string
}

func createResultDataSuccess(data interface{}) ResultData {
	return ResultData{
		1,
		data,
		"success"}
}

func createResultDataError(code int, message string) ResultData {
	return ResultData{
		code,
		"",
		message}
}
