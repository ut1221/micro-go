package service

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/mojocn/base64Captcha"
	pb "micro/api/system/v1"
	"micro/internal/biz"
	"micro/internal/pkg"
	"micro/internal/pkg/constants"
)

type AuthService struct {
	pb.UnimplementedAuthServer
	uc  *biz.SysUserUsecase
	rc  *biz.SysRoleUsecase
	mc  *biz.SysMenuUsecase
	ac  *biz.CoreUsecase
	log *log.Helper
}

var Store = base64Captcha.DefaultMemStore

func NewAuthService(uc *biz.SysUserUsecase, rc *biz.SysRoleUsecase, mc *biz.SysMenuUsecase, ac *biz.CoreUsecase, logger log.Logger) *AuthService {
	return &AuthService{uc: uc, rc: rc, mc: mc, ac: ac, log: log.NewHelper(logger)}
}
func (s *AuthService) Logout(ctx context.Context, req *pb.LogoutReq) (*pb.LogoutReply, error) {
	return &pb.LogoutReply{}, nil
}
func (s *AuthService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {
	//判断验证码是否正确
	data, err := s.ac.RedisGetData(ctx, constants.CacheCaptcha+req.Uuid)
	if err != nil {
		return nil, errors.New("验证码失效")
	}
	if data.(string) != req.Code {
		//验证码错误则重新获取验证码
		_ = s.ac.RedisRemoveData(ctx, constants.CacheCaptcha+req.Uuid)
		return nil, errors.New("验证码验证失败")
	}
	//根据用户名查询用户信息
	user, err := s.uc.GetUserByUserName(ctx, req.Username)
	if err != nil {
		return nil, err
	}
	//判断密码是否匹配
	if !pkg.Verify(user.Password, req.Password) {
		return nil, errors.New("密码验证失败")
	}
	//返回token
	token, err := pkg.GenerateToken(user.UserID, user.UserName)
	if err != nil {
		return nil, err
	}
	return &pb.LoginReply{Token: token}, nil
}
func (s *AuthService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterReply, error) {
	//注册用户
	//角色权限分配
	return &pb.RegisterReply{}, nil
}
func (s *AuthService) Captcha(ctx context.Context, req *pb.CaptchaReq) (*pb.CaptchaReply, error) {
	driver := base64Captcha.NewDriverDigit(80, 250, 4, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, Store)
	id, b64s, answer, err := cp.Generate()
	if err != nil {
		return nil, err
	}
	//将验证码添加到redis中
	err = s.ac.RedisSetData(ctx, constants.CacheCaptcha+id, answer)
	if err != nil {
		return nil, err
	}

	return &pb.CaptchaReply{
		Img:           b64s,
		Uuid:          id,
		CaptchaEnable: true,
	}, nil
}
func (s *AuthService) UserInfo(ctx context.Context, req *pb.UserInfoReq) (*pb.UserInfoReply, error) {
	userId := pkg.GetLoginUserId(ctx)
	user, _ := s.uc.GetUserInfoById(ctx, userId)
	roles, _ := s.rc.GetRoleByUserId(ctx, userId)
	user.Roles = roles
	return &pb.UserInfoReply{
		Permissions: []string{constants.AllPermission},
		Roles:       []string{constants.SuperAdmin},
		User:        user}, nil
}
func (s *AuthService) Routers(ctx context.Context, req *pb.RoutersReq) (*pb.RoutersReply, error) {
	menus, err := s.mc.GetMenuByUserId(ctx, pkg.GetLoginUserId(ctx), true)
	if err != nil {
		return nil, err
	}
	return &pb.RoutersReply{Routers: menus}, nil
}
