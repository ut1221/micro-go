package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	v1 "micro/api/system/v1"
	"micro/internal/model"
)

type SysPostRepo interface {
	GetPostList(ctx context.Context, query *model.SysPostQuery) ([]*model.BizSysPost, error)
	ModifyPostForUser(ctx context.Context, userId string, postIds []string) error
	GetPageSet(ctx context.Context, pageNum int64, pageSize int64, query model.SysPostQuery) ([]*model.BizSysPost, int64, error)
	GetPostInfo(ctx context.Context, query *model.SysPostQuery) (*model.BizSysPost, error)
	Save(ctx context.Context, post *model.SysPost) error
	Update(ctx context.Context, post *model.SysPost) error
	Delete(ctx context.Context, postId string) error
	DeleteSysUserPost(ctx context.Context, postId string) error
}

type SysPostUsecase struct {
	repo SysPostRepo
	log  *log.Helper
}

func NewSysPostUsecase(repo SysPostRepo, logger log.Logger) *SysPostUsecase {
	return &SysPostUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysPostUsecase) GetPostList(ctx context.Context, query *model.SysPostQuery) ([]*v1.PostReply, error) {
	list, err := uc.repo.GetPostList(ctx, query)
	if err != nil {
		return nil, err
	}
	var posts []*v1.PostReply
	for _, bizPost := range list {
		toV1 := bizPost.BizToV1()
		posts = append(posts, toV1)
	}
	return posts, nil
}

func (uc *SysPostUsecase) ModifyPostForUser(ctx context.Context, userId string, postIds []string) error {
	return uc.repo.ModifyPostForUser(ctx, userId, postIds)
}

func (uc *SysPostUsecase) GetPageSet(ctx context.Context, req *v1.ListSysPostRep) (*v1.ListSysPostReply, error) {
	bizPosts, total, err := uc.repo.GetPageSet(ctx, req.PageNum, req.PageSize, model.SysPostQuery{})
	if err != nil {
		return nil, err
	}
	var posts []*v1.PostReply
	for _, bizPost := range bizPosts {
		toV1 := bizPost.BizToV1()
		posts = append(posts, toV1)
	}
	return &v1.ListSysPostReply{
		Total: total,
		Rows:  posts,
	}, nil
}
func (uc *SysPostUsecase) CreateSysPost(ctx context.Context, req *v1.SysPostRep) (*v1.EmptyReply, error) {
	_, err := uc.repo.GetPostInfo(ctx, &model.SysPostQuery{PostCode: req.PostCode})
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("编码已存在")
	}
	if err := uc.repo.Save(ctx, uc.V1ToEntity(req)); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}
func (uc *SysPostUsecase) UpdateSysPost(ctx context.Context, req *v1.SysPostRep) (*v1.EmptyReply, error) {
	if err := uc.repo.Update(ctx, uc.V1ToEntity(req)); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}
func (uc *SysPostUsecase) DeleteSysPost(ctx context.Context, req *v1.DeleteSysPostRep) (*v1.EmptyReply, error) {
	if err := uc.repo.Delete(ctx, req.GetId()); err != nil {
		return nil, err
	}
	if err := uc.repo.DeleteSysUserPost(ctx, req.GetId()); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}
func (uc *SysPostUsecase) GetSysPost(ctx context.Context, req *v1.GetSysPostRep) (*v1.GetSysPostReply, error) {
	info, err := uc.repo.GetPostInfo(ctx, &model.SysPostQuery{PostID: req.GetId()})
	if err != nil {
		return nil, err
	}
	return &v1.GetSysPostReply{
		Post: info.BizToV1(),
	}, nil
}

func (uc *SysPostUsecase) V1ToEntity(v1 *v1.SysPostRep) *model.SysPost {
	return &model.SysPost{
		PostID:   v1.PostId,
		PostCode: v1.PostCode,
		PostName: v1.PostName,
		PostSort: v1.PostSort,
		Status:   v1.Status,
		Remark:   v1.Remark,
	}
}
