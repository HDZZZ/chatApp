package user

import "github.com/gin-gonic/gin"

type FriendService interface {
	sendRequestOfAddingFriend(c *gin.Context)
	agreeRequestOfAddingFriend(c *gin.Context)
	refuseRequestOfAddingFriend(c *gin.Context)
	getRequestOfAddingFriendList(c *gin.Context)
	getAllFriendsInfo(c *gin.Context)
	getAllFriendsUid(c *gin.Context)
}

type FriendServiceInstance struct {
}
