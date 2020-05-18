package handler

import (
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	s_users "micro.demo/Api/servers/users"
	"micro.demo/Proto/pb"
)

type UserHandler struct {
}

// 登录用户
func (u *UserHandler) Login(ctx *gin.Context) {
	var (
		err error
		req *pb.UserLoginRequest
		res *pb.UserLoginResponse
	)

	if err = ctx.Bind(&req); err != nil {
		res = &pb.UserLoginResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		ctx.JSON(200, res)
		return
	}

	// 调用登录
	if res, err = s_users.Login(req.GetEmail(), req.GetPassword()); err != nil {
		res = &pb.UserLoginResponse{
			Code: 1,
			Msg:  err.Error(),
		}
		ctx.JSON(200, res)
		return
	}

	ctx.JSON(200, res)
}

// 查询用户
func (u *UserHandler) Find(ctx *gin.Context) {
	var (
		err error
		req pb.UserFindByWhereRequest
		res *pb.UserFindByWhereResponse
	)

	if err = ctx.Bind(&req); err != nil {
		res = &pb.UserFindByWhereResponse{
			Code: 1,
			Msg:  "参数错误",
		}

		ctx.JSON(200, res)
		return
	}

	if res, err = s_users.Find(req.GetUserId(), req.GetEmail(), req.GetPhone(), req.GetToken()); err != nil {
		res = &pb.UserFindByWhereResponse{
			Code: 1,
			Msg:  "获取用户失败:" + err.Error(),
		}

		ctx.JSON(200, res)
		return
	}

	ctx.JSON(200, res)
}

func (u *UserHandler) FindByIds(ctx *gin.Context) {
	ids, _ := ctx.GetQuery("id")
	var (
		users []int64
	)
	for _, id := range strings.Split(ids, ",") {
		i, _ := strconv.ParseInt(id, 10, 64)
		users = append(users, i)
	}

	res, err := s_users.SelectByIds(users)
	if err != nil {
		res = &pb.UserFindByIdsResponse{
			Code: 1,
			Msg:  "获取用户列表错误: " + err.Error(),
		}

		ctx.JSON(200, res)
		return
	}

	ctx.JSON(200, res)
}

// 注册用户
func (u *UserHandler) Register(ctx *gin.Context) {
	var (
		err error
		req pb.UserRegisterRequest
		res *pb.UserRegisterResponse
	)

	if err = ctx.Bind(&req); err != nil {
		res = &pb.UserRegisterResponse{
			Code: 1,
			Msg:  "参数错误",
		}

		ctx.JSON(200, res)
		return
	}

	// 注册
	if res, err = s_users.Register(
		req.GetUserName(),
		req.GetPassword(),
		req.GetPhone(),
		req.GetEmail(),
		req.GetAvatar(),
	); err != nil {
		res = &pb.UserRegisterResponse{
			Code: 1,
			Msg:  "获取用户失败:" + err.Error(),
		}

		ctx.JSON(200, res)
		return
	}

	ctx.JSON(200, res)
}

// 更新用户
func (u *UserHandler) Update(ctx *gin.Context) {
	var (
		err error
		req pb.UserUpdateProfileRequest
		res *pb.UserUpdateProfileResponse
	)

	// 绑定参数
	if err = ctx.Bind(&req); err != nil {
		res = &pb.UserUpdateProfileResponse{
			Code: 1,
			Msg:  "参数错误",
		}

		ctx.JSON(200, res)
		return
	}

	// 调用更新
	res, err = s_users.Update(req.GetUserId(), req.GetUserName(), req.GetPassword(), req.GetPhone(), req.GetEmail(), req.GetAvatar(), req.GetToken())
	if err != nil {
		res = &pb.UserUpdateProfileResponse{
			Code: 1,
			Msg:  "更新错误: " + err.Error(),
		}

		ctx.JSON(200, res)
		return
	}

	ctx.JSON(200, res)
}
