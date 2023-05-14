package main

import (
	DB "github.com/HDDDZ/test/chatApp/db"
)

func main() {
	// http.ListenAndServe(":9002", nil)
	defer DB.AppClosed()
}
