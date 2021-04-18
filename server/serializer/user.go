package serializer

import "server/model"

// User 用户序列化器
type User struct {
	UID       uint64 `json:"uid"`
	UserName  string `json:"user_name"`
	CreatedAt int64  `json:"created_at"`
}

// BuildUser 序列化用户
func BuildUser(user model.User) User {
	return User{
		UID:       user.UID,
		UserName:  user.UserName,
		CreatedAt: user.CreatedAt.Unix(),
	}
}

// BuildUsers 序列化用户列表
func BuildUsers(items []model.User) []User {
	var users []User
	for _, item := range items {
		userV := BuildUser(item)
		users = append(users, userV)
	}
	return users
}

// BuildUserResponse 序列化用户响应
func BuildUserResponse(user model.User) Response {
	return Response{
		Data: BuildUser(user),
	}
}
