package s_users

import (
	"context"

	"github.com/micro/go-micro/client"

	"micro.demo/Proto/pb"
)

var (
	userClient = pb.NewUserService("micro.demo.user", client.DefaultClient)
)

// 用户登录
func Login(email, password string) (*pb.UserLoginResponse, error) {
	req := pb.UserLoginRequest{
		Email:    email,
		Password: password,
	}

	// 请求接口
	return userClient.UserLogin(context.TODO(), &req)
}

// 注册用户
func Register(userName, password, phone, email, avatar string) (res *pb.UserRegisterResponse, err error) {
	request := pb.UserRegisterRequest{
		UserName: userName,
		Password: password,
		Phone:    phone,
		Email:    email,
		Avatar:   avatar,
	}

	return userClient.UserRegister(context.TODO(), &request)
}

// 更新用户
func Update(userId int64, userName, password, phone, email, avatar, token string) (*pb.UserUpdateProfileResponse, error) {
	request := pb.UserUpdateProfileRequest{
		UserId:   userId,
		UserName: userName,
		Password: password,
		Email:    email,
		Phone:    phone,
		Avatar:   avatar,
		Token:    token,
	}

	return userClient.UserUpdateProfile(context.TODO(), &request)
}

// 查询用户信息
func Find(userId int64, email, phone, token string) (res *pb.UserFindByWhereResponse, err error) {
	request := pb.UserFindByWhereRequest{
		UserId: userId,
		Email:  email,
		Phone:  phone,
		Token:  token,
	}

	res, err = userClient.UserFind(context.TODO(), &request)
	if err != nil {
		return
	}

	return
}

// 查询用户列表
func SelectByIds(id []int64) (res *pb.UserFindByIdsResponse, err error) {
	var (
		req = pb.UserFindByIdsRequest{
			Ids: id,
		}
	)

	return userClient.UserFindIds(context.TODO(), &req)
}
