package Servers

import (
	"errors"
	"strconv"

	"micro.demo/Library/jwt"
	"micro.demo/Library/log"
	"micro.demo/Library/orm"
	"micro.demo/Library/tool"
	"micro.demo/Models"
	"micro.demo/Proto/pb"
)

// 用户注册
func Register(userName, password, phone, email, avatar string) error {
	user := Models.User{
		UserId:   tool.GenerateUUIDInt(),
		UserName: userName,
		Password: password,
		Phone:    phone,
		Email:    email,
		Avatar:   avatar,
	}

	// 验证参数
	if err := tool.Validate(user); err != nil {
		return err
	}

	err := orm.DB().Create(&user).Error
	if err != nil {
		return err
	}

	return nil
}

// 用户登录
func Login(email, password string) (token string, err error) {
	var (
		data  Models.User
		where = Models.User{
			Email: email,
		}
	)

	// 查询用户信息
	data, err = Find(where)
	if err != nil {
		err = errors.New("登录失败: " + err.Error())
		return
	}

	// 验证密码
	if data.Password != password {
		err = errors.New("用户密码错误")
		return
	}

	// 生成token
	token, err = jwt.CreateToken(strconv.FormatInt(data.UserId, 10))
	return
}

// 更新用户
func Update(userId int64, token string, user Models.User) (err error) {
	var (
		data  Models.User
		where Models.User
	)

	if userId < 1 && token == "" {
		err = errors.New("参数错误")
		return
	}

	// 验证参数
	if err := tool.Validate(user); err != nil {
		return err
	}

	if userId > 0 {
		where.UserId = userId
		data, err = Find(where)
	} else {
		data, err = FindByToken(token)
	}

	if err != nil {
		err = errors.New("查询用户失败: " + err.Error())
		return
	}
	err = orm.DB().Model(&data).Update(&user).Error
	if err != nil {
		return
	}

	return nil
}

// 查询用户
func Find(user Models.User) (data Models.User, err error) {
	if user.UserId < 1 && user.Phone == "" && user.Email == "" {
		err = errors.New("参数错误")
		log.Debug(err)
		return
	}

	err = orm.DB().Where(user).Find(&data).Error
	if err != nil {
		log.Debug(err)
		return
	}

	return
}

// 查询用户
func FindByToken(token string) (user Models.User, err error) {
	var where Models.User
	if token == "" {
		err = errors.New("参数错误")
		log.Debug(err)
		return
	}

	log.Info("token=", token)
	p, e := jwt.ParseToken(token)
	if e != nil {
		err = e
		log.Debug(err)
		return
	}

	// 获取用户Id
	if userId, e := strconv.ParseInt(p.UID, 10, 64); e != nil {
		err = e
		log.Info("p.UID=", p.UID)
		log.Debug(err)
		return
	} else {
		where.UserId = userId
		log.Debug("用户Id: ", userId)
	}

	user, err = Find(where)
	user.SetAvatar("")
	return
}

// 查询用户列表
func SelectByIds(ids []int64) (users []*pb.UserItem, err error) {
	var user []Models.User
	if err = orm.DB().Where(ids).Find(&user).Error; err != nil {
		return
	}

	for _, item := range user {
		u := pb.UserItem{
			UserId:   item.UserId,
			UserName: item.UserName,
			Phone:    item.Phone,
			Email:    item.Email,
			Avatar:   item.SetAvatar(""),
		}

		users = append(users, &u)
	}
	return
}
