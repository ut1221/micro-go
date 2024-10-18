package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/net/context"
	v1 "micro/api/system/v1"
	"micro/internal/model"
	"micro/internal/pkg"
	"strings"
	"time"
)

type SysUserRepo interface {
	GetUserByUserName(ctx context.Context, username string) (*model.BizSysUser, error)
	GetUserById(ctx context.Context, userId string) (*model.BizSysUser, error)
	GetPageSet(ctx context.Context, pageNum int64, pageSize int64, query model.SysUserQuery) ([]*model.BizSysUser, int64, error)
	Save(ctx context.Context, user *model.SysUser) error
	Update(ctx context.Context, user *model.BizSysUser) error
	Delete(ctx context.Context, userId string) error
	GetRoleUser(ctx context.Context, pageNum int64, pageSize int64, authUser bool, query model.SysUserQuery) ([]*model.BizSysUser, int64, error)
}

type SysUserUsecase struct {
	repo     SysUserRepo
	roleRepo SysRoleRepo
	log      *log.Helper
}

func NewSysUserUsecase(repo SysUserRepo, roleRepo SysRoleRepo, logger log.Logger) *SysUserUsecase {
	return &SysUserUsecase{repo: repo, roleRepo: roleRepo, log: log.NewHelper(logger)}
}

func (uc *SysUserUsecase) GetUserByUserName(ctx context.Context, username string) (*model.BizSysUser, error) {
	return uc.repo.GetUserByUserName(ctx, username)
}

func (uc *SysUserUsecase) GetUserInfoById(ctx context.Context, userId string) (*v1.UserReply, error) {
	var userInfo *v1.UserReply
	user, err := uc.repo.GetUserById(ctx, userId)
	if err != nil {
		return nil, err
	}
	userInfo = user.BizToV1()
	userInfo.Dept = user.Dept.BizToV1()
	return userInfo, nil
}

func (uc *SysUserUsecase) GetPageSet(ctx context.Context, req *v1.ListSysUserRep, userId int64) (*v1.ListSysUserReply, error) {
	query := model.SysUserQuery{
		DeptId:      req.GetDeptId(),
		UserName:    req.GetUserName(),
		BeginTime:   req.GetParams().GetBeginTime(),
		EndTime:     req.GetParams().GetEndTime(),
		Status:      req.GetStatus(),
		PhoneNumber: req.GetPhoneNumber(),
	}
	bizUsers, total, err := uc.repo.GetPageSet(ctx, req.PageNum, req.PageSize, query)
	if err != nil {
		return nil, err
	}
	var users []*v1.ListSysUser
	for _, user := range bizUsers {
		var listUser v1.ListSysUser
		user := &v1.UserReply{
			UserId:      user.UserID,
			DeptId:      user.DeptID,
			UserName:    user.UserName,
			NickName:    user.NickName,
			Email:       user.Email,
			PhoneNumber: user.PhoneNumber,
			Sex:         user.Sex,
			Avatar:      user.Avatar,
			Password:    user.Password,
			Status:      user.Status,
			LoginIp:     user.LoginIP,
			LoginDate:   user.LoginDate.Format("2006-01-02 15:04:05"),
			Remark:      user.Remark,
			Dept:        user.Dept.BizToV1(),
			CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		listUser.User = user
		users = append(users, &listUser)
	}
	return &v1.ListSysUserReply{Rows: users, Total: total}, nil
}

func (uc *SysUserUsecase) Save(ctx context.Context, user *model.BizSysUser) error {
	sysUser := &model.SysUser{
		UserID:      user.UserID,
		DeptID:      user.DeptID,
		UserName:    user.UserName,
		NickName:    user.NickName,
		UserType:    user.UserType,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		Sex:         user.Sex,
		Avatar:      user.Avatar,
		Password:    pkg.Encrypt(user.Password),
		Status:      user.Status,
		LoginDate:   time.Now(),
		LoginIP:     user.LoginIP,
		Remark:      user.Remark,
	}
	return uc.repo.Save(ctx, sysUser)
}

func (uc *SysUserUsecase) Delete(ctx context.Context, userId string) error {
	return uc.repo.Delete(ctx, userId)
}

func (uc *SysUserUsecase) Update(ctx context.Context, user *model.BizSysUser) error {
	return uc.repo.Update(ctx, user)
}

func (uc *SysUserUsecase) GetAuthRoleSysUser(ctx context.Context, req *v1.GetAuthRoleSysUserRep) (*v1.GetAuthRoleSysUserReply, error) {
	req.GetId()
	bizRoles, err := uc.roleRepo.GetListRoles(ctx, model.SysRoleQuery{UserId: req.GetId()})
	if err != nil {
		return nil, err
	}
	var roles []*v1.RoleReply
	var roleIds []string
	for _, bizRole := range bizRoles {
		toV1 := bizRole.BizToV1()
		roles = append(roles, toV1)
		roleIds = append(roleIds, toV1.RoleId)
	}
	user, err := uc.repo.GetUserById(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	userInfo := user.BizToV1()
	userInfo.Dept = user.Dept.BizToV1()
	userInfo.Roles = roles
	userInfo.RoleIds = roleIds
	userInfo.DeptId = user.Dept.BizToV1().DeptId
	listRoles, err := uc.roleRepo.GetListRoles(ctx, model.SysRoleQuery{})
	if err != nil {
		return nil, err
	}
	var allRoles []*v1.RoleReply
	for _, role := range listRoles {
		toV1 := role.BizToV1()
		for _, id := range roleIds {
			if toV1.RoleId == id {
				toV1.Flag = true
			}
		}
		allRoles = append(allRoles, toV1)
	}
	return &v1.GetAuthRoleSysUserReply{Roles: allRoles, User: userInfo}, nil
}

func (uc *SysUserUsecase) AuthRoleSysUser(ctx context.Context, req *v1.AuthRoleSysUserRep) (*v1.AuthRoleSysUserReply, error) {
	roleIds := strings.Split(req.GetRoleIds(), ",")
	err := uc.roleRepo.ModifyRoleForUser(ctx, req.GetUserId(), roleIds)
	if err != nil {
		return nil, err
	}
	return &v1.AuthRoleSysUserReply{}, err
}
