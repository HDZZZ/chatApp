package db

import (
	"testing"

	"github.com/HDDDZ/test/chatApp/data/common"
)

func TestFriendSendRequest(t *testing.T) {
	// sendRequest(100000, 100011, "的买鞋呀")
}

func TestFriendAgreeRequest(t *testing.T) {
	agreeRequest(41)
}

func TestFriendRefuseRequest(t *testing.T) {
	refuseRequest(42)
}

func TestFriendMakeRequestState(t *testing.T) {
	makeRequestState(100000, 100001, common.AlreadyAgree)
}

func TestFriendGetAllRequestOfSomebody(t *testing.T) {
	requests := getAllRequestOfSomebody(100000)
	log("TestFriendGetAllRequestOfSomebody=", requests)
}

func TestFriendGetAllFriends(t *testing.T) {
	requests := getAllFriends(100000)
	log("TestFriendGetAllFriends=", requests)
}

func TestFriendGetAllFriendsUid(t *testing.T) {
	requests := getAllFriendsUid(100000)
	log("TestFriendGetAllFriendsUid=", requests)
}

func TestFriendDeleteFriend(t *testing.T) {
	deleteFriend(100018, 100015)
}

func TestFriendQueryRequestById(t *testing.T) {
	requests := queryRequestById(42)
	log("TestFriendQueryRequestById=", requests)
}

func TestFriendQueryRequestBySuidAndRuid(t *testing.T) {
	requests := queryRequestBySuidAndRuid(100000, 100010)
	log("TestFriendQueryRequestBySuidAndRuid=", requests)
}

func TestFriendQueryRequestByUids(t *testing.T) {
	requests := queryRequestByUids(100000, 100010)
	log("TestFriendQueryRequestByUids=", requests)
}

func TestFriendFriendWithSomeone(t *testing.T) {
	friendWithSomeone(100000, 100010)
}
