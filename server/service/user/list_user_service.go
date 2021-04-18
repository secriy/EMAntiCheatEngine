package user

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
)

// ListUsersService 用户列表服务
type ListUsersService struct {
}

// List 用户列表
func (service *ListUsersService) List() serializer.Response {
	var userVar model.User
	var users []model.User

	rows, err := database.DB.Query(`SELECT * FROM user ORDER BY uid`)
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}
	for rows.Next() {
		var (
			uid       uint64
			userName  string
			password  string
			createdAt time.Time
		)

		err = rows.Scan(&uid, &userName, &password, &createdAt)
		if err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库操作错误",
				Error: err.Error(),
			}
		}
		userVar.UID = uid
		userVar.UserName = userName
		userVar.CreatedAt = createdAt

		users = append(users, userVar)
	}

	return serializer.Response{
		Data: serializer.BuildUsers(users),
	}
}
