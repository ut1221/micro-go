package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "micro/api/api/system/v1"
	"micro/internal/model"
	"micro/internal/pkg/constants"
	"strconv"
	"strings"
)

type SysMenuRepo interface {
	GetMenuByUserId(ctx context.Context, userId string, isAdmin bool) ([]*model.BizSysMenu, error)
	GetMenusList(ctx context.Context, query *model.SysMenuQuery) ([]*model.BizSysMenu, error)
	ModifyMenuForRole(ctx context.Context, roleId string, menuIds []string) error
	Save(ctx context.Context, menu *model.SysMenu) error
	Update(ctx context.Context, menu *model.SysMenu) error
	Delete(ctx context.Context, menuId string) error
	DeleteRoleMenu(ctx context.Context, menuId string, roleId string) error
	GetMenuInfo(ctx context.Context, query model.SysMenuQuery) (*model.BizSysMenu, error)
}

type SysMenuUsecase struct {
	repo SysMenuRepo
	log  *log.Helper
}

func NewSysMenuUsecase(repo SysMenuRepo, logger log.Logger) *SysMenuUsecase {
	return &SysMenuUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysMenuUsecase) GetMenuByUserId(ctx context.Context, userId string, isAdmin bool) ([]*v1.Router, error) {
	bizMenus, err := uc.repo.GetMenuByUserId(ctx, userId, isAdmin)
	if err != nil {
		return nil, err
	}
	var menus []*v1.Router
	for _, bizMenu := range bizMenus {
		var router v1.Router
		router.MenuId = bizMenu.MenuID
		router.ParentId = bizMenu.ParentID
		router.Path = bizMenu.Path
		router.Component = bizMenu.Component
		if bizMenu.ParentID != "0" && bizMenu.IsFrame == constants.YesFrame {
			router.Redirect = bizMenu.Path
			router.Component = constants.InnerLink
		}
		if bizMenu.ParentID == "0" && bizMenu.IsFrame == constants.YesFrame {
			router.Redirect = bizMenu.Path
			router.Component = constants.Layout
		}
		if bizMenu.ParentID == "0" && bizMenu.IsFrame == constants.NoFrame {
			router.Path = "/" + bizMenu.Path
			router.Redirect = "noRedirect"
			router.Component = constants.Layout
		}
		if bizMenu.ParentID != "0" && bizMenu.MenuType == constants.TypeDir {
			router.Component = constants.ParentView
		}
		router.Name = strings.ToUpper(bizMenu.Path[:1]) + bizMenu.Path[1:]
		router.Meta = &v1.Router_Meta{
			Title:   bizMenu.MenuName,
			NoCache: bizMenu.IsCache == 1,
			Icon:    bizMenu.Icon,
			Link:    router.Redirect,
		}
		menus = append(menus, &router)
	}
	tree := Tree(menus, "0")
	return tree, nil
}

func (uc *SysMenuUsecase) GetMenusList(ctx context.Context, req *v1.ListSysMenuRep) (*v1.ListSysMenuReply, error) {
	bizMenus, err := uc.repo.GetMenusList(ctx, &model.SysMenuQuery{})
	if err != nil {
		return nil, err
	}
	var menus []*v1.MenuReply
	for _, bizMenu := range bizMenus {
		toV1 := bizMenu.BizToV1()
		menus = append(menus, toV1)
	}
	return &v1.ListSysMenuReply{
		Menus: menus,
	}, err
}

func (uc *SysMenuUsecase) GetSysRoleMenu(ctx context.Context, req *v1.GetSysRoleMenuRep) (*v1.GetSysRoleMenuReply, error) {
	userList, err := uc.repo.GetMenusList(ctx, &model.SysMenuQuery{RoleId: req.GetRoleId()})
	if err != nil {
		return nil, err
	}
	var checkedKeys []string
	for _, menu := range userList {
		checkedKeys = append(checkedKeys, menu.MenuID)
	}
	menusList, err := uc.repo.GetMenusList(ctx, &model.SysMenuQuery{})
	if err != nil {
		return nil, err
	}
	var menus []*v1.RoleMenuReply
	for _, bizMenu := range menusList {
		var menu v1.RoleMenuReply
		menu.Id = bizMenu.MenuID
		menu.Pid = bizMenu.ParentID
		menu.Label = bizMenu.MenuName
		menus = append(menus, &menu)
	}
	tree := RoleMenuTree(menus, "0")
	return &v1.GetSysRoleMenuReply{
		Menus:       tree,
		CheckedKeys: checkedKeys,
	}, nil
}

func (uc *SysMenuUsecase) GetTreeSelect(ctx context.Context, req *v1.GetTreeSelectRep) (*v1.GetTreeSelectReply, error) {
	menusList, err := uc.repo.GetMenusList(ctx, &model.SysMenuQuery{})
	if err != nil {
		return nil, err
	}
	var menus []*v1.RoleMenuReply
	for _, bizMenu := range menusList {
		var menu v1.RoleMenuReply
		menu.Id = bizMenu.MenuID
		menu.Pid = bizMenu.ParentID
		menu.Label = bizMenu.MenuName
		menus = append(menus, &menu)
	}
	tree := RoleMenuTree(menus, "0")
	return &v1.GetTreeSelectReply{
		Menus: tree,
	}, nil
}

func (uc *SysMenuUsecase) Save(ctx context.Context, req *v1.SysMenuRep) (*v1.EmptyReply, error) {
	if err := uc.repo.Save(ctx, uc.V1ToEntity(req)); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}

func (uc *SysMenuUsecase) Update(ctx context.Context, req *v1.SysMenuRep) (*v1.EmptyReply, error) {
	if err := uc.repo.Update(ctx, uc.V1ToEntity(req)); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}

func (uc *SysMenuUsecase) Delete(ctx context.Context, menuId string) (*v1.EmptyReply, error) {
	if err := uc.repo.Delete(ctx, menuId); err != nil {
		return nil, err
	}
	//删除用户已分配菜单
	if err := uc.repo.DeleteRoleMenu(ctx, menuId, ""); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}

func (uc *SysMenuUsecase) GetSysMenu(ctx context.Context, req *v1.GetSysMenuRep) (*v1.GetSysMenuReply, error) {
	bizMenu, err := uc.repo.GetMenuInfo(ctx, model.SysMenuQuery{MenuID: req.GetId()})
	if err != nil {
		return nil, err
	}
	return &v1.GetSysMenuReply{
		Menu: bizMenu.BizToV1(),
	}, nil
}

func Tree(node []*v1.Router, pid string) []*v1.Router {
	res := make([]*v1.Router, 0)
	for _, v := range node {
		if v.ParentId == pid {
			v.Children = Tree(node, v.MenuId)
			res = append(res, v)
		}
	}
	return res
}
func RoleMenuTree(node []*v1.RoleMenuReply, pid string) []*v1.RoleMenuReply {
	res := make([]*v1.RoleMenuReply, 0)
	for _, v := range node {
		if v.Pid == pid {
			v.Children = RoleMenuTree(node, v.Id)
			res = append(res, v)
		}
	}
	return res
}
func (uc *SysMenuUsecase) V1ToEntity(req *v1.SysMenuRep) *model.SysMenu {
	isFrame, _ := strconv.ParseInt(req.IsFrame, 10, 64)
	isCache, _ := strconv.ParseInt(req.IsCache, 10, 64)
	return &model.SysMenu{
		MenuID:    req.MenuId,
		ParentID:  req.ParentId,
		MenuType:  req.MenuType,
		OrderNum:  req.OrderNum,
		MenuName:  req.MenuName,
		Status:    req.Status,
		Path:      req.Path,
		IsFrame:   isFrame,
		IsCache:   isCache,
		Visible:   req.Visible,
		Query:     req.Query,
		Component: req.Component,
		Perms:     req.Perms,
		Icon:      req.Icon,
	}
}
