package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "micro/api/system/v1"
	"micro/internal/model"
	"micro/internal/pkg"
	"strings"
)

type SysRoleRepo interface {
	GetListRoles(ctx context.Context, query model.SysRoleQuery) ([]*model.BizSysRole, error)
	ModifyRoleForUser(ctx context.Context, userId string, roleIds []string) error
	AuthUserSelectAll(ctx context.Context, roleId string, userIds []string) error
	GetPageSet(ctx context.Context, pageNum int64, pageSize int64, query model.SysRoleQuery) ([]*model.BizSysRole, int64, error)
	GetRoleInfo(ctx context.Context, m *model.SysRoleQuery) (*model.BizSysRole, error)
	Save(ctx context.Context, role *model.SysRole) error
	Update(ctx context.Context, role *model.SysRole) error
	Delete(ctx context.Context, roleId string) error
	AuthUserCancel(ctx context.Context, roleId string, userIds []string) error
}

type SysRoleUsecase struct {
	repo     SysRoleRepo
	menuRepo SysMenuRepo
	userRepo SysUserRepo
	deptRepo SysDeptRepo
	log      *log.Helper
}

func NewSysRoleUsecase(repo SysRoleRepo, menuRepo SysMenuRepo, deptRepo SysDeptRepo, userRepo SysUserRepo, logger log.Logger) *SysRoleUsecase {
	return &SysRoleUsecase{repo: repo, menuRepo: menuRepo, deptRepo: deptRepo, userRepo: userRepo, log: log.NewHelper(logger)}
}

func (uc *SysRoleUsecase) GetRoleByUserId(ctx context.Context, userId string) ([]*v1.RoleReply, error) {
	bizRoles, _ := uc.repo.GetListRoles(ctx, model.SysRoleQuery{UserId: userId})
	var roles []*v1.RoleReply
	for _, bizRole := range bizRoles {
		roles = append(roles, bizRole.BizToV1())
	}
	return roles, nil
}

func (uc *SysRoleUsecase) GetListRoles(ctx context.Context) ([]*v1.RoleReply, error) {
	bizRoles, err := uc.repo.GetListRoles(ctx, model.SysRoleQuery{})
	if err != nil {
		return nil, err
	}
	var roles []*v1.RoleReply
	for _, bizRole := range bizRoles {
		roles = append(roles, bizRole.BizToV1())
	}
	return roles, nil
}

func (uc *SysRoleUsecase) ModifyRoleForUser(ctx context.Context, userId string, roleIds []string) error {
	return uc.repo.ModifyRoleForUser(ctx, userId, roleIds)
}

func (uc *SysRoleUsecase) GetPageSet(ctx context.Context, req *v1.ListSysRoleRep) (*v1.ListSysRoleReply, error) {
	bizRoles, total, err := uc.repo.GetPageSet(ctx, req.GetPageNum(), req.GetPageSize(), model.SysRoleQuery{})
	if err != nil {
		return nil, err
	}
	var roles []*v1.RoleReply
	for _, bizRole := range bizRoles {
		toV1 := bizRole.BizToV1()
		if bizRole.DataScope == "1" {
			toV1.Admin = true
		} else {
			toV1.Admin = false
		}
		roles = append(roles, toV1)
	}
	return &v1.ListSysRoleReply{
		Rows:  roles,
		Total: total,
	}, nil
}

func (uc *SysRoleUsecase) GetRoleById(ctx context.Context, req *v1.GetSysRoleRep) (*v1.GetSysRoleReply, error) {
	info, err := uc.repo.GetRoleInfo(ctx, &model.SysRoleQuery{RoleID: req.GetRoleId()})
	if err != nil {
		return nil, err
	}
	return &v1.GetSysRoleReply{
		Role: info.BizToV1(),
	}, nil
}

func (uc *SysRoleUsecase) CreateSysRole(ctx context.Context, req *v1.ModifySysRoleRep) (*v1.EmptyReply, error) {
	roleId := pkg.GetID()
	err := uc.repo.Save(ctx, &model.SysRole{
		RoleID:            roleId,
		RoleSort:          req.GetRoleSort(),
		RoleName:          req.GetRoleName(),
		Remark:            req.GetRemark(),
		Status:            req.GetStatus(),
		RoleKey:           req.GetRoleKey(),
		MenuCheckStrictly: boolCoverInt(req.GetMenuCheckStrictly()),
	})
	if err != nil {
		return nil, err
	}
	return uc.modifyRoleMenus(ctx, req.GetMenuIds(), roleId)
}
func (uc *SysRoleUsecase) UpdateSysRole(ctx context.Context, req *v1.ModifySysRoleRep) (*v1.EmptyReply, error) {
	err := uc.repo.Update(ctx, &model.SysRole{
		RoleID:            req.RoleId,
		RoleName:          req.GetRoleName(),
		RoleSort:          req.GetRoleSort(),
		Remark:            req.GetRemark(),
		Status:            req.GetStatus(),
		RoleKey:           req.GetRoleKey(),
		MenuCheckStrictly: boolCoverInt(req.GetMenuCheckStrictly()),
	})
	if err != nil {
		return nil, err
	}
	return uc.modifyRoleMenus(ctx, req.GetMenuIds(), req.RoleId)
}
func boolCoverInt(b bool) int64 {
	if b {
		return 1
	} else {
		return 0
	}
}
func (uc *SysRoleUsecase) modifyRoleMenus(ctx context.Context, menuIds []string, roleId string) (*v1.EmptyReply, error) {
	if err := uc.menuRepo.ModifyMenuForRole(ctx, roleId, menuIds); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}

func (uc *SysRoleUsecase) Delete(ctx context.Context, roleId string) (*v1.EmptyReply, error) {
	if err := uc.repo.Delete(ctx, roleId); err != nil {
		return nil, err
	}
	// 删除sys_role_menu
	if err := uc.menuRepo.DeleteRoleMenu(ctx, roleId, ""); err != nil {
		return nil, err
	}
	// 删除sys_role_dept
	// 删除sys_user_role
	return &v1.EmptyReply{}, nil
}

func (uc *SysRoleUsecase) DataScope(ctx context.Context, req *v1.DataScopeSysRoleRep) (*v1.EmptyReply, error) {
	err := uc.repo.Update(ctx, &model.SysRole{RoleID: req.RoleId, DataScope: req.DataScope})
	if err != nil {
		return nil, err
	}
	err = uc.deptRepo.ModifyDeptForRole(ctx, req.RoleId, req.DeptIds)
	return &v1.EmptyReply{}, err
}

func (uc *SysRoleUsecase) ChangeStatus(ctx context.Context, req *v1.ChangeStatusSysRoleRep) (*v1.EmptyReply, error) {
	if err := uc.repo.Update(ctx, &model.SysRole{RoleID: req.RoleId, Status: req.Status}); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}

func (uc *SysRoleUsecase) IsAllocatedList(ctx context.Context, req *v1.IsAllocatedListRep, authUser bool) (*v1.IsAllocatedListReply, error) {
	query := model.SysUserQuery{
		RoleId: req.RoleId,
	}
	bizUsers, total, err := uc.userRepo.GetRoleUser(ctx, req.PageNum, req.PageSize, authUser, query)
	if err != nil {
		return nil, err
	}
	var users = make([]*v1.UserReply, 0)
	for _, bizUser := range bizUsers {
		user := &v1.UserReply{
			UserId:      bizUser.UserID,
			DeptId:      bizUser.DeptID,
			UserName:    bizUser.UserName,
			NickName:    bizUser.NickName,
			Email:       bizUser.Email,
			PhoneNumber: bizUser.PhoneNumber,
			Sex:         bizUser.Sex,
			Avatar:      bizUser.Avatar,
			Password:    bizUser.Password,
			Status:      bizUser.Status,
			LoginIp:     bizUser.LoginIP,
			LoginDate:   bizUser.LoginDate.Format("2006-01-02 15:04:05"),
			Remark:      bizUser.Remark,
			Dept:        bizUser.Dept.BizToV1(),
			CreatedAt:   bizUser.CreatedAt.Format("2006-01-02 15:04:05"),
		}
		users = append(users, user)
	}
	return &v1.IsAllocatedListReply{Rows: users, Total: total}, nil
}

func (uc *SysRoleUsecase) AuthUserSelectAll(ctx context.Context, req *v1.AuthUserSelectAllRep) (*v1.EmptyReply, error) {
	userIds := strings.Split(req.UserIds, ",")
	if err := uc.repo.AuthUserSelectAll(ctx, req.RoleId, userIds); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}

func (uc *SysRoleUsecase) AuthUserCancel(ctx context.Context, req *v1.AuthUserCancelRep) (*v1.EmptyReply, error) {
	var userIds = make([]string, 0)
	userIds = append(userIds, req.UserId)
	if err := uc.repo.AuthUserCancel(ctx, req.RoleId, userIds); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}

func (uc *SysRoleUsecase) AuthUserCancelAll(ctx context.Context, req *v1.AuthUserSelectAllRep) (*v1.EmptyReply, error) {
	userIds := strings.Split(req.UserIds, ",")
	if err := uc.repo.AuthUserCancel(ctx, req.RoleId, userIds); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}
