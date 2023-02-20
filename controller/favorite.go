package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// FavoriteAction no practical effect, just check if token is valid
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")
	video_id := c.Query("video_id")
	videoId,err := strconv.ParseInt(video_id, 0, 64)
	if err!=nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Num error"})
	}
	action_type := c.Query("action_type")
	action,err := strconv.ParseInt(action_type, 0, 64)
	if err!=nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Num error"})
	}

	if user, exist := usersLoginInfo[token]; exist {
		addFavorite(user, videoId, action)
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FavoriteList all users have same favorite video list
func FavoriteList(c *gin.Context) {
	s_id := c.Query("user_id")
	id,err := strconv.ParseInt(s_id, 0, 64)
	if err!=nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Num error"})
	}

	if err := favoriteFind(id); err == nil {

		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: DemoLoveVideos,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
