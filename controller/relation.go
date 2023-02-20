package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")
	author_id := c.Query("to_user_id")
	authorId,err := strconv.ParseInt(author_id, 0, 64)
	if err!=nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Num error"})
	}
	action_type := c.Query("action_type")
	action,err := strconv.ParseInt(action_type, 0, 64)
	if err!=nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Num error"})
	}

	if user, exist := usersLoginInfo[token]; exist {
		if err := addFollow(user, authorId, action); err == 1 {
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		}else {
			if err == 3{
				c.JSON(http.StatusOK, Response{StatusCode: 1,  StatusMsg: "你已关注该用户"})
			}else if err == 2{
				c.JSON(http.StatusOK, Response{StatusCode: 1,  StatusMsg: "你已取关该用户"})
			}
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	s_id := c.Query("user_id")
	id,err := strconv.ParseInt(s_id, 0, 64)
	if err!=nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Num error"})
	}
	if err := findFollow(id); err!=nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}else{
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: DemoFollows,
		})
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	s_id := c.Query("user_id")
	id,err := strconv.ParseInt(s_id, 0, 64)
	if err!=nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Num error"})
	}
	if err := findFollower(id); err!=nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}else{
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{
				StatusCode: 0,
			},
			UserList: DemoFollowers,
		})
	}
}

// FriendList all users have same friend list
func FriendList(c *gin.Context) {
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: DemoFollows,
	})
}
