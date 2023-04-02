package logic

import (
	"gin_forum/dao/mysql"
	"gin_forum/dao/redis"
	"gin_forum/models"
	"gin_forum/pkg/snowflake"
	"strconv"
)

func GetPostList(page, size int64) (data []*models.ApiPostList, err error) {

	posts, err := mysql.GetPostList(page, size)
	if err != nil {
		return
	}
	data = make([]*models.ApiPostList, 0, len(posts))
	for _, post := range posts {

		user, err := mysql.GetPostAuthorName(post.AuthorID)
		if err != nil {
			return nil, err
		}
		likenum := redis.GetPostLikeNum(strconv.Itoa(int(post.PostID)))

		postDetail := &models.ApiPostList{
			Title:      post.Title,
			LikeNum:    int64(likenum),
			AuthorName: user.Username,
			PostID:     post.PostID,
		}
		data = append(data, postDetail)
	}

	return
}

func CreatePost(p *models.Post) (err error) {
	//1.生成post id
	p.PostID = snowflake.GenID()
	p.LikeNum = 0
	//2. 保存到数据库
	err = mysql.CreatePost(p)
	if err != nil {
		return err
	}
	redis.CreatePost(p.PostID)
	return
	//3. 返回

}

func GetPostDetail(postid int64) (*models.ApiPostDetail, error) {

	PostDetail, err := mysql.GetPostDetail(postid)
	if err != nil {
		return nil, err
	}

	return PostDetail, nil

}
