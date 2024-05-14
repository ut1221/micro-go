package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	v1 "micro/api/api/system/v1"
	"micro/internal/model"
	"strings"
)

type SysDeptRepo interface {
	GetSysDeptInfo(ctx context.Context, query model.SysDeptQuery) (*model.BizSysDept, error)
	GetDeptList(ctx context.Context, query model.SysDeptQuery) ([]*model.BizSysDept, error)
	ModifyDeptForRole(ctx context.Context, roleId string, deptIds []string) error
	Save(ctx context.Context, dept *model.SysDept) error
	Update(ctx context.Context, dept *model.SysDept) error
	Delete(ctx context.Context, deptId string) error
	DeleteSysDept(ctx context.Context, deptId string) error
}

type SysDeptUsecase struct {
	repo SysDeptRepo
	log  *log.Helper
}

func NewSysDeptUsecase(repo SysDeptRepo, logger log.Logger) *SysDeptUsecase {
	return &SysDeptUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SysDeptUsecase) GetDeptTree(ctx context.Context) ([]*v1.DeptTree, error) {
	var deptTree []*v1.DeptTree
	list, err := uc.repo.GetDeptList(ctx, model.SysDeptQuery{})
	if err != nil {
		return nil, err
	}
	for _, dept := range list {
		var deptTreeReply = &v1.DeptTree{
			Id:    dept.DeptID,
			Pid:   dept.ParentID,
			Label: dept.DeptName,
		}
		deptTree = append(deptTree, deptTreeReply)
	}
	dept := uc.TreeDept(deptTree, "0")
	return dept, err
}

func (uc *SysDeptUsecase) ExcludeDept(ctx context.Context, req *v1.ExcludeDeptRep) (*v1.ListSysDeptReply, error) {
	bizDepts, err := uc.repo.GetDeptList(ctx, model.SysDeptQuery{DeptId: req.GetId()})
	if err != nil {
		return nil, err
	}
	var depts []*v1.DeptReply
	for _, bizDept := range bizDepts {
		if req.GetId() == bizDept.DeptID {
			continue
		}
		ancestors := strings.Split(bizDept.Ancestors, ",")
		found := false
		for _, ancestor := range ancestors {
			if ancestor == req.GetId() {
				found = true
				break
			}
		}
		if found {
			continue
		}
		toV1 := bizDept.BizToV1()
		depts = append(depts, toV1)
	}
	return &v1.ListSysDeptReply{Dept: depts}, err
}

func (uc *SysDeptUsecase) GetDeptList(ctx context.Context, req *v1.ListSysDeptRep) (*v1.ListSysDeptReply, error) {
	bizDepts, err := uc.repo.GetDeptList(ctx, model.SysDeptQuery{})
	if err != nil {
		return nil, err
	}
	var depts []*v1.DeptReply
	for _, bizDept := range bizDepts {
		toV1 := bizDept.BizToV1()
		depts = append(depts, toV1)
	}
	return &v1.ListSysDeptReply{Dept: depts}, err
}
func (uc *SysDeptUsecase) GetSysRoleDept(ctx context.Context, req *v1.GetSysRoleDeptRep) (*v1.GetSysRoleDeptReply, error) {
	tree, err := uc.GetDeptTree(ctx)
	if err != nil {
		return nil, err
	}
	var checkedKeys []string
	list, err := uc.repo.GetDeptList(ctx, model.SysDeptQuery{RoleId: req.RoleId})
	for _, dept := range list {
		checkedKeys = append(checkedKeys, dept.DeptID)
	}
	return &v1.GetSysRoleDeptReply{Dept: tree, CheckedKeys: checkedKeys}, err
}

func (uc *SysDeptUsecase) Delete(ctx context.Context, req *v1.DeleteSysDeptRep) (*v1.EmptyReply, error) {
	if err := uc.repo.Delete(ctx, req.Id); err != nil {
		return nil, err
	}
	if err := uc.repo.DeleteSysDept(ctx, req.Id); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}

func (uc *SysDeptUsecase) Sava(ctx context.Context, req *v1.SysDeptRep) (*v1.EmptyReply, error) {

	_, err := uc.repo.GetSysDeptInfo(ctx, model.SysDeptQuery{DeptName: req.GetDeptName()})
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, errors.New("该部门已经存在")
	}
	parentInfo, _ := uc.repo.GetSysDeptInfo(ctx, model.SysDeptQuery{DeptId: req.GetParentId()})
	entity := uc.V1ToEntity(req)
	entity.Ancestors = parentInfo.Ancestors + "," + req.GetParentId()
	if err := uc.repo.Save(ctx, entity); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}

func (uc *SysDeptUsecase) Update(ctx context.Context, req *v1.SysDeptRep) (*v1.EmptyReply, error) {
	deptInfo, err := uc.repo.GetSysDeptInfo(ctx, model.SysDeptQuery{DeptName: req.GetDeptName()})
	if !errors.Is(err, gorm.ErrRecordNotFound) && deptInfo.DeptName != req.DeptName {
		return nil, errors.New("该部门已经存在")
	}
	parentInfo, _ := uc.repo.GetSysDeptInfo(ctx, model.SysDeptQuery{DeptId: req.GetParentId()})
	entity := uc.V1ToEntity(req)
	entity.Ancestors = parentInfo.Ancestors + "," + req.GetParentId()
	if err := uc.repo.Update(ctx, entity); err != nil {
		return nil, err
	}
	return &v1.EmptyReply{}, nil
}

func (uc *SysDeptUsecase) GetSysDeptData(ctx context.Context, req *v1.GetSysDeptRep) (*v1.GetSysDeptReply, error) {
	bizDept, err := uc.repo.GetSysDeptInfo(ctx, model.SysDeptQuery{DeptId: req.GetId()})
	if err != nil {
		return nil, err
	}
	return &v1.GetSysDeptReply{Dept: bizDept.BizToV1()}, nil
}

func (uc *SysDeptUsecase) TreeDept(node []*v1.DeptTree, pid string) []*v1.DeptTree {
	res := make([]*v1.DeptTree, 0)
	for _, v := range node {
		if v.Pid == pid {
			v.Children = uc.TreeDept(node, v.Id)
			res = append(res, v)
		}
	}
	return res
}

func (uc *SysDeptUsecase) V1ToEntity(req *v1.SysDeptRep) *model.SysDept {
	return &model.SysDept{
		DeptID:   req.DeptId,
		ParentID: req.ParentId,
		DeptName: req.DeptName,
		OrderNum: req.OrderNum,
		Leader:   req.Leader,
		Phone:    req.Phone,
		Email:    req.Email,
		Status:   req.Status,
	}
}
