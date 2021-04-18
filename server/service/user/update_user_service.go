package user

import (
	"strconv"
	"time"

	"server/database"
	"server/model"
	"server/serializer"
)

// UpdateUserService 用户更新服务
type UpdateUserService struct {
	UserName string `form:"user_name" json:"user_name" binding:"required,excludes= ,min=2,max=10"`
}

// Update 用户更新
func (service *UpdateUserService) Update(uid string) serializer.Response {
	var userVar model.User
	var count = 0
	var createAt time.Time

	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM user WHERE uid = ?`, uid).Scan(&count)
	_ = database.DB.QueryRow(`SELECT create_at FROM user WHERE uid = ?`, uid).Scan(&createAt)
	if count == 0 {
		return serializer.Response{
			Code: 404,
			Msg:  "用户不存在",
		}
	}
	intNum, _ := strconv.Atoi(uid)
	userVar.UID = uint64(intNum)
	userVar.UserName = service.UserName
	userVar.CreatedAt = createAt

	_, err := database.DB.Exec(`
				UPDATE user SET username = ? WHERE uid = ?`, userVar.UserName, uid)
	if err != nil {
		return serializer.Response{
			Code:  50003,
			Msg:   "用户更新失败",
			Error: err.Error(),
		}
	}

	return serializer.Response{
		Data: serializer.BuildUser(userVar),
		Msg:  "用户更新成功",
	}
}
