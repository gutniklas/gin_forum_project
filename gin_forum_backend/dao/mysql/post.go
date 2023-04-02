package mysql

import (
	"gin_forum/models"
)

// GetPostList 查询帖子列表函数
func GetPostList(page, size int64) (posts []*models.Post, err error) {
	sqlStr := `select post_id, title, content, author_id, likenum,create_time
	from post
	order by create_time
	DESC
	limit ?,?
	`
	posts = make([]*models.Post, 0, 2)
	err = db.Select(&posts, sqlStr, (page-1)*size, size)
	return
}
// GetPostAuthorName 查询帖子作者名称函数
func GetPostAuthorName(AuthorID int64) (user *models.User, err error) {
	user = new(models.User)
	sqlStr := `select username from user where user_id=?`
	err = db.Get(user, sqlStr, AuthorID)
	return

}
// CreatePost 创建帖子函数
func CreatePost(p *models.Post) (err error) {
	sqlStr := `insert into post(post_id,author_id,likenum,title,content) values (?, ?, ?, ?, ?)`
	_, err = db.Exec(sqlStr, p.PostID, p.AuthorID, p.LikeNum, p.Title, p.Content)
	return

}
//  GetPostDetail 进入帖子函数
func GetPostDetail(postid int64) (postdetail *models.ApiPostDetail, err error) {
	postdetail = new(models.ApiPostDetail)
	sqlStr := `select   likenum, post_id, title, content from post where post_id=?`
	err = db.Get(postdetail, sqlStr, postid)
	return
}
