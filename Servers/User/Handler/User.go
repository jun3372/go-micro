package Handler

import (
	"context"

	"micro.demo/Models"
	"micro.demo/Proto/pb"
	"micro.demo/Servers/User/Servers"
)

type User struct {
}

// 用户注册
func (u *User) UserRegister(ctx context.Context, req *pb.UserRegisterRequest, res *pb.UserRegisterResponse) error {
	// 开始注册
	if err := Servers.Register(
		req.GetUserName(),
		req.GetPassword(),
		req.GetPhone(),
		req.GetEmail(),
		req.GetAvatar(),
	); err != nil {
		res.Code = 1
		res.Msg = "注册失败: " + err.Error()
		return err
	}

	res.Code = 0
	res.Msg = "注册成功"
	return nil
}

// 用户登录
func (u *User) UserLogin(ctx context.Context, req *pb.UserLoginRequest, res *pb.UserLoginResponse) error {
	token, err := Servers.Login(req.GetEmail(), req.GetPassword())
	if err != nil {
		return err
	}

	// 返回成功
	res.Code = 0
	res.Msg = "登录成功"
	res.Data = &pb.UserLoginResponseData{
		Token: token,
	}
	return nil
}

// 用户更新
func (u *User) UserUpdateProfile(ctx context.Context, req *pb.UserUpdateProfileRequest, res *pb.UserUpdateProfileResponse) error {
	user := Models.User{
		UserName: req.GetUserName(),
		Password: req.GetPassword(),
		Phone:    req.GetPhone(),
		Email:    req.GetEmail(),
		Avatar:   req.GetAvatar(),
	}

	if err := Servers.Update(req.UserId, req.GetToken(), user); err != nil {
		res.Code = 1
		res.Msg = err.Error()
		return nil
	}

	res.Code = 0
	res.Msg = "更新成功"
	return nil
}

// 用户查询信息
func (u *User) UserFind(ctx context.Context, req *pb.UserFindByWhereRequest, res *pb.UserFindByWhereResponse) error {
	var (
		err  error
		data Models.User
		user = Models.User{
			UserId: req.GetUserId(),
			Phone:  req.GetPhone(),
			Email:  req.GetEmail(),
		}
	)

	// token查询
	if req.GetToken() != "" {
		if data, err = Servers.FindByToken(req.GetToken()); err != nil {
			return err
		}
	} else {
		// 其他参数查询
		if data, err = Servers.Find(user); err != nil {
			return err
		}
	}

	res.Code = 0
	res.Msg = "查询成功"
	res.Data = &pb.UserItem{
		UserId:   data.UserId,
		UserName: data.UserName,
		Email:    data.Email,
		Phone:    data.Phone,
		Avatar:   data.Avatar,
	}
	return nil
}

// 用户查询列表
func (u *User) UserFindIds(ctx context.Context, in *pb.UserFindByIdsRequest, out *pb.UserFindByIdsResponse) (err error) {
	users, err := Servers.SelectByIds(in.GetIds())
	if err != nil {
		out.Code = 1
		out.Msg = "获取用户列表失败: " + err.Error()
		return
	}

	out.Code = 0
	out.Msg = "获取用户列表成功"
	out.Data = users
	return
}
