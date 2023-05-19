package main

import (
	Channel "github.com/HDDDZ/test/chatApp/channel"
	Common "github.com/HDDDZ/test/chatApp/common"
	DB "github.com/HDDDZ/test/chatApp/db"
	Group "github.com/HDDDZ/test/chatApp/group"
	User "github.com/HDDDZ/test/chatApp/user"
)

func main() {
	Common.InitHTTPService()
	//run.如果不显示调用,无法加入其他包内部内容
	User.UserMain()
	Channel.ChannelMain()
	Group.GroupMain()
	defer DB.AppClosed()
}
