module chatApp/channel

go 1.20

require (
	github.com/HDDDZ/test/chatApp/db v0.0.0
	github.com/HDDDZ/test/chatApp/common v0.0.0
	github.com/gorilla/websocket v1.5.0
)

require (
	github.com/HDDDZ/test/chatApp/util v0.0.0 // indirect
	github.com/go-sql-driver/mysql v1.7.1 // indirect
)

replace github.com/HDDDZ/test/chatApp/db => /Users/z./project/goProject/test/chatApp/DB

replace github.com/HDDDZ/test/chatApp/util => /Users/z./project/goProject/test/chatApp/util
replace github.com/HDDDZ/test/chatApp/common => /Users/z./project/goProject/test/chatApp/common
