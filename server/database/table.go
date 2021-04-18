package database

import "server/util"

// 创建表
func createTable() {
	// 用户表
	_, err := DB.Exec(`
		CREATE TABLE IF NOT EXISTS user(
			uid MEDIUMINT(8) UNSIGNED  AUTO_INCREMENT,
			username VARCHAR(20) NOT NULL UNIQUE,
			password VARCHAR(200) NOT NULL,
			create_at TIMESTAMP NOT NULL,
			PRIMARY KEY(uid)
			);`)
	if err != nil {
		util.Log().Panic("创建用户表失败", err)
	}

	// HookLog表
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS hook(
			id INT UNSIGNED AUTO_INCREMENT,
			ip VARCHAR(15) NOT NULL,
			player VARCHAR(15) NOT NULL,
			content TEXT NOT NULL,
			create_at TIMESTAMP NOT NULL,
			PRIMARY KEY(id)
		)`)
	if err != nil {
		util.Log().Panic("创建HookLog表失败", err)
	}

	// AntiLog表
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS anti(
			id INT UNSIGNED AUTO_INCREMENT,
			ip VARCHAR(15) NOT NULL,
			player VARCHAR(15) NOT NULL,
			content TEXT NOT NULL,
			create_at TIMESTAMP NOT NULL,
			PRIMARY KEY(id)
		)`)
	if err != nil {
		util.Log().Panic("创建AntiLog表失败", err)
	}
}
