package database

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"server/util"
)

// 创建默认管理员
func defaultAdmin(password string) error {
	// 查询是否已经存在默认管理员，
	var username string
	row := DB.QueryRow(`
			SELECT username FROM user WHERE username = ?;`, "admin")
	_ = row.Scan(&username)

	// 密码加密
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return err
	}

	// 创建默认管理员
	if username == "" {
		_, err := DB.Exec(`INSERT INTO pvzac.user(username,password,create_at )
		VALUES (?,?,?)`, "admin", string(bytes), time.Now())
		if err != nil {
			util.Log().Panic("创建默认管理员失败", err)
		}

	}
	return nil
}
