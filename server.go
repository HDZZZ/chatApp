package main

import (
	Common "github.com/HDDDZ/test/chatApp/common"
	DB "github.com/HDDDZ/test/chatApp/db"
)

func main() {
	Common.InitHTTPService()
	defer DB.AppClosed()
}
