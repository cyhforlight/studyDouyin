package controller

import (
	"fmt"
	
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)



func (v Video) TableName() string {
	return "videos"
}

func (u User) TableName() string {
	return "users"
}

func InitVideo() error {
	dsn := dbUserName+":"+dbUserPass+"@tcp(" + IpAddress + ":3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed to connect database")
		return err
	}

	db.Find(&DemoLogins)
	db.Find(&DemoUsers)
	for i := 1; i <= int(len(DemoUsers)); i++ {
		toke := DemoLogins[i-1].Name + DemoLogins[i-1].Password
		usersLoginInfo[toke] = DemoUsers[i-1]
	}

	var demoVideoL = Video{}
	db.Last(&demoVideoL)
	for i := 1; i <= int(demoVideoL.Id); i++ {
		var demoVideo = Video{}
		db.Preload("Author").Find(&demoVideo, i)
		demoVideo.PlayUrl = "http://" + IpAddress + ":8080/static/play/" + demoVideo.PlayUrl
		demoVideo.CoverUrl = "http://" + IpAddress + ":8080/static/cover/" + demoVideo.CoverUrl
		DemoVideos = append(DemoVideos, demoVideo)
	}

	var DemoCommentL = Comment{}
	db.Last(&DemoCommentL)
	for i := 1; i <= int(DemoCommentL.Id); i++ {
		var demoComment = Comment{}
		db.Preload("User").Find(&demoComment, i)
		DemoComments = append(DemoComments, demoComment)
	}

	return nil
}

func dbCreate(v Video) error {
	dsn := dbUserName+":"+dbUserPass+"@tcp(" + IpAddress + ":3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed to connect database")
		return err
	}

	db.Model(&Video{}).Omit("Author").Create(&v)
	v.PlayUrl = "http://" + IpAddress + ":8080/static/play/" + v.PlayUrl
	v.CoverUrl = "http://" + IpAddress + ":8080/static/cover/" + v.CoverUrl
	DemoVideos = append(DemoVideos, v)
	return nil
}

func userCreate(newUser User, newLogin login) error {
	dsn := dbUserName+":"+dbUserPass+"@tcp(" + IpAddress + ":3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed to connect database")
		return err
	}

	db.Model(&User{}).Create(&newUser)
	db.Model(&login{}).Create(&newLogin)

	return nil
}

func publishFind(id int64) error {
	dsn := dbUserName+":"+dbUserPass+"@tcp(" + IpAddress + ":3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed to connect database")
		return err
	}

	DemoPublishs = DemoPublishs[:0]

	db.Where("author_id = ?", id).Find(&DemoPublishs)
	for i := 0; i < len(DemoPublishs); i++ {
		DemoPublishs[i].PlayUrl = "http://" + IpAddress + ":8080/static/play/" + DemoPublishs[i].PlayUrl
		DemoPublishs[i].CoverUrl = "http://" + IpAddress + ":8080/static/cover/" + DemoPublishs[i].CoverUrl
	}

	return nil
}

func favoriteFind(id int64) error {
	dsn := dbUserName+":"+dbUserPass+"@tcp(" + IpAddress + ":3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed to connect database")
		return err
	}

	DemoLoveAdmins = DemoLoveAdmins[:0]
	DemoLoveVideos = DemoLoveVideos[:0]

	db.Where("user_id = ?", id).Find(&DemoLoveAdmins)
	for _, item := range DemoLoveAdmins {
		var demoVideo = Video{}
		db.Find(&demoVideo, item.VideoId)
		demoVideo.PlayUrl = "http://" + IpAddress + ":8080/static/play/" + demoVideo.PlayUrl
		demoVideo.CoverUrl = "http://" + IpAddress + ":8080/static/cover/" + demoVideo.CoverUrl
		DemoLoveVideos = append(DemoLoveVideos, demoVideo)
	}

	return nil
}

func favoriteLike(user User) error {
	if err := favoriteFind(user.Id); err != nil {
		return err
	}

	for _, item := range DemoLoveAdmins {
		DemoVideos[int(item.VideoId)-1].IsFavorite = true
	}

	return nil
}

func favoriteClear() {
	for i := 0; i<len(DemoVideos); i++ {
		DemoVideos[i].IsFavorite = false
	} 
}

func addFavorite(user User, videoId int64, action int64) error {
	dsn := dbUserName+":"+dbUserPass+"@tcp(" + IpAddress + ":3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed to connect database")
		return err
	}

	if action == 1 {
		var fa1 = loveVideo{}
		db.Last(&fa1)

		var fa = loveVideo{
			Id: int64(fa1.Id + 1), UserId: user.Id, VideoId: videoId}
		db.Create(fa)
		DemoVideos[int(videoId-1)].IsFavorite = true
		DemoVideos[int(videoId-1)].FavoriteCount = DemoVideos[int(videoId-1)].FavoriteCount + 1
		db.Model(Video{}).Where("id = ?", videoId).Updates(Video{
			FavoriteCount: DemoVideos[int(videoId-1)].FavoriteCount})
	} else {
		db.Where("user_id = ? AND video_id = ?", user.Id, videoId).Delete(&loveVideo{})
		DemoVideos[int(videoId-1)].IsFavorite = false
		DemoVideos[int(videoId-1)].FavoriteCount = DemoVideos[int(videoId-1)].FavoriteCount - 1
		db.Model(Video{}).Where("id = ?", videoId).Updates(Video{
			FavoriteCount: DemoVideos[int(videoId-1)].FavoriteCount})
	}

	return nil
}

func findFollow(id int64) error {
	dsn := dbUserName+":"+dbUserPass+"@tcp(" + IpAddress + ":3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed to connect database")
		return err
	}

	DemoFollows = DemoFollows[:0]
	db.Where("follower_id = ?", id).Find(&Interests)
	for _, item := range Interests {
		var demoUser = User{}
		db.Where("id = ?", item.FollowId).Find(&demoUser)
		if DemoUsers[int(item.FollowId-1)].IsFollow {
			demoUser.IsFollow = true
		}
		DemoFollows = append(DemoFollows, demoUser)
	}

	return nil
}

func followLike(user User) error {
	if err := findFollow(user.Id); err != nil {
		return err
	}

	for i:=0 ; i<len(Interests); i++{
		DemoUsers[int(DemoFollows[i].Id)-1].IsFollow = true
	}

	return nil
}

func followClear() {
	for i := 0; i<len(DemoUsers); i++ {
		DemoUsers[i].IsFollow = false
	} 
}

func addFollow(user User, authorId int64, action int64) int {
	dsn := dbUserName+":"+dbUserPass+"@tcp(" + IpAddress + ":3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed to connect database")
		return 4
	}

	if (action == 1) {
		if (DemoUsers[int(authorId-1)].IsFollow){
			return 3
		}
		var fa1 = fan{}
		db.Last(&fa1)

		var fa = fan{
			Id: int64(fa1.Id + 1), FollowId: authorId, FollowerId: user.Id}
		db.Create(fa)
		DemoUsers[int(authorId-1)].IsFollow = true
		DemoUsers[int(authorId-1)].FollowerCount = DemoUsers[int(authorId-1)].FollowerCount + 1
		DemoUsers[int(user.Id-1)].FollowCount = DemoUsers[int(user.Id-1)].FollowCount + 1
		db.Model(User{}).Where("id = ?", authorId).Updates(User{
			FollowerCount: DemoUsers[int(authorId-1)].FollowerCount})
		db.Model(User{}).Where("id = ?", user.Id).Updates(User{
			FollowCount: DemoUsers[int(user.Id-1)].FollowCount})
	}else{
		if (!DemoUsers[int(authorId-1)].IsFollow){
			return 2
		}
		db.Where("follow_id = ? AND follower_id = ?", authorId, user.Id).Delete(&fan{})
		DemoUsers[int(authorId-1)].IsFollow = false
		DemoUsers[int(authorId-1)].FollowerCount = DemoUsers[int(authorId-1)].FollowerCount - 1
		DemoUsers[int(user.Id-1)].FollowCount = DemoUsers[int(user.Id-1)].FollowCount - 1
		db.Model(User{}).Where("id = ?", authorId).Updates(User{
			FollowerCount: DemoUsers[int(authorId-1)].FollowerCount})
		db.Model(User{}).Where("id = ?", user.Id).Updates(User{
			FollowCount: DemoUsers[int(user.Id-1)].FollowCount})
	}

	return 1
}

func findFollower(id int64) error {
	dsn := dbUserName+":"+dbUserPass+"@tcp(" + IpAddress + ":3306)/douyin?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		fmt.Println("failed to connect database")
		return err
	}

	DemoFollowers = DemoFollowers[:0]
	db.Where("follow_id = ?", id).Find(&Fans)
	for _, item := range Fans {
		var demoUser = User{}
		db.Where("id = ?", item.FollowerId).Find(&demoUser)
		if DemoUsers[int(item.FollowerId-1)].IsFollow {
			demoUser.IsFollow = true
		}
		DemoFollowers = append(DemoFollowers, demoUser)
	}

	return nil
}
