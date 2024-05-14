package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"micro/internal/biz"
	"micro/internal/model"
)

type sysPostRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysPostRepo(data *Data, logger log.Logger) biz.SysPostRepo {
	return &sysPostRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s sysPostRepo) GetPostList(ctx context.Context, query *model.SysPostQuery) ([]*model.BizSysPost, error) {
	var post []model.SysPost
	db := s.data.Db.Table("sys_post")
	if query.UserId != "" {
		db.Joins("INNER JOIN sys_user_post ON sys_post.post_id = sys_user_post.post_id and sys_user_post.deleted_at is null").
			Where("sys_user_post.user_id = ?", query.UserId)
	}
	err := db.Find(&post).Error
	if err != nil {
		return nil, err
	}
	var posts []*model.BizSysPost
	for _, sysPost := range post {
		toBiz := sysPost.EntityToBiz()
		posts = append(posts, toBiz)
	}
	return posts, nil
}

func (s sysPostRepo) GetPageSet(ctx context.Context, pageNum int64, pageSize int64, query model.SysPostQuery) ([]*model.BizSysPost, int64, error) {
	var sysPosts []model.SysPost
	var total int64
	db := s.data.Db.Model(&model.SysPost{}).Count(&total)
	tx := db.Limit(int(pageSize)).Offset(int((pageNum - 1) * pageSize)).Find(&sysPosts)
	if tx.Error != nil {
		return nil, 0, tx.Error
	}
	var bizPosts []*model.BizSysPost
	for _, sysPost := range sysPosts {
		toBiz := sysPost.EntityToBiz()
		bizPosts = append(bizPosts, toBiz)
	}
	return bizPosts, total, nil
}

func (s sysPostRepo) ModifyPostForUser(ctx context.Context, userId string, postIds []string) error {
	var sysUserPost = make([]model.SysUserPost, 0)
	for _, postId := range postIds {
		sysUserPost = append(sysUserPost, model.SysUserPost{
			UserID: userId,
			PostID: postId,
		})
	}
	tx := s.data.Db.Where("user_id = ?", userId).Delete(&model.SysUserPost{})
	if tx.Error != nil {
		return tx.Error
	}
	return s.data.Db.CreateInBatches(sysUserPost, 100).Error
}

func (s sysPostRepo) GetPostInfo(ctx context.Context, query *model.SysPostQuery) (*model.BizSysPost, error) {
	var sysPost model.SysPost
	db := s.data.Db.Model(&sysPost)
	if query.PostID != "" {
		db.Where("post_id = ?", query.PostID)
	}
	if query.PostCode != "" {
		db.Where("post_code = ?", query.PostCode)
	}
	if err := db.First(&sysPost).Error; err != nil {
		return nil, err
	}
	return sysPost.EntityToBiz(), nil
}

func (s sysPostRepo) Save(ctx context.Context, post *model.SysPost) error {
	return s.data.Db.Save(&post).Error
}

func (s sysPostRepo) Update(ctx context.Context, post *model.SysPost) error {
	return s.data.Db.Model(&model.SysPost{}).Where("post_id = ?", post.PostID).Updates(&post).Error
}

func (s sysPostRepo) Delete(ctx context.Context, postId string) error {
	return s.data.Db.Where("post_id = ?", postId).Delete(&model.SysPost{}).Error
}

func (s sysPostRepo) DeleteSysUserPost(ctx context.Context, postId string) error {
	return s.data.Db.Where("post_id = ?", postId).Delete(&model.SysUserPost{}).Error
}
