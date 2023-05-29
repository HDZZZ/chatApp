package db

import (
	"testing"
)

func TestUserQueryUserByToken(t *testing.T) {
	users := queryUserByToken("f37f076b5c45a38304b38cbec296cf56987db688a38d986e0b01316a955283bf")
	log("TestUserQueryUserByToken,users=", users)
}

func TestUserQueryUserByUserName(t *testing.T) {
	users := queryUserByUserName("2023052507")
	log("TestUserQueryUserByUserName,users=", users)
}

func TestUserQueryUserByUserNameAndPwd(t *testing.T) {
	users, _ := queryUserByUserNameAndPwd("2023052507", "1111111")
	log("TestUserQueryUserByUserNameAndPwd,users=", users)
}

func TestUserAddUser(t *testing.T) {
	users, _ := addUser("2023052529", "chili", "asfjsafjasoijfaioj")
	log("TestUserQueryUserByUserNameAndPwd,users=", users)
}
