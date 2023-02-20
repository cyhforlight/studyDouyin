package controller

import (
	"fmt"
	"net/http"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type VideoListResponse struct {
	Response
	VideoList []Video `json:"video_list"`
}

// Publish check token then save upload file to public directory
func Publish(c *gin.Context) {
	token := c.PostForm("token")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./public/play", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	name := strings.Split(finalName, ".")
	coverName := name[0] + ".jpg"
	var video = Video{Id: int64(len(DemoVideos) + 1), AuthorId: user.Id, Author: user, PlayUrl: finalName,
		CoverUrl: coverName, FavoriteCount: 0, CommentCount: 0,
		IsFavorite: false}

	//SaveCover(finalName, coverName)
	dbCreate(video)

	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}

func SaveCover(finalName string, coverName string) {
	cmd := exec.Command("ffmpeg", "-i", "./public/play/"+finalName, "-ss", "1", "-f", "image2", "./public/cover/"+coverName)
	if err := cmd.Run(); err != nil {
		fmt.Print(err)
	}
}

// PublishList all users have same publish video list
func PublishList(c *gin.Context) {
	s_id := c.Query("user_id")
	id,err := strconv.ParseInt(s_id, 0, 64)
	if err!=nil {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "Num error"})
	}

	if err := publishFind(id); err == nil {

		c.JSON(http.StatusOK, VideoListResponse{
			Response: Response{
				StatusCode: 0,
			},
			VideoList: DemoPublishs,
		})
	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}
