package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"server/database"
)

// 用户模型
type User struct {
	UID       uint64
	UserName  string
	Password  string
	CreatedAt time.Time
}

const PassWordCost = 12

// GetUser 用ID获取用户
func GetUser(ID interface{}) (User, error) {
	var user User
	result := database.DB.QueryRow(`SELECT * FROM user WHERE UID = ?`, ID).Scan(&user.UID, &user.UserName, &user.Password, &user.CreatedAt)
	return user, result
}

// SetPassword 设置密码
func (user *User) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), PassWordCost)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword 校验密码
func (user *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	return err == nil
}
