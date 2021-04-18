package user

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
	"server/util"
)

// AddUserService 增加用户服务
type AddUserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,excludes= ,min=2,max=10"`
	Password string `form:"password" json:"password" binding:"required,excludes= ,min=6,max=16"`
}

// valid 验证表单
func (service *AddUserService) valid() *serializer.Response {
	count := 0
	_ = database.DB.QueryRow(`SELECT  COUNT(username) FROM user WHERE username = ?`, service.UserName).Scan(&count)
	if count > 0 {
		return &serializer.Response{
			Code: 40001,
			Msg:  "用户名已被占用",
		}
	}
	return nil
}

// Create 用户创建
func (service *AddUserService) Create() serializer.Response {
	user := model.User{
		UserName:  service.UserName,
		Password:  service.Password,
		CreatedAt: time.Now(),
	}

	// 表单验证
	if err := service.valid(); err != nil {
		return *err
	}

	// 加密密码
	if err := user.SetPassword(service.Password); err != nil {
		return serializer.Err(
			serializer.CodeEncryptError,
			"密码加密失败",
			err,
		)
	}

	// 创建用户
	res, err := database.DB.Exec(`INSERT INTO pvzac.user(username,password,create_at )
		VALUES (?,?,?)`, user.UserName, user.Password, user.CreatedAt)
	if res != nil {
		tmp, _ := res.LastInsertId()
		user.UID = uint64(tmp)
	}

	if err != nil {
		util.Log().Panic("创建用户失败", err)
	}

	return serializer.BuildUserResponse(user)
}
