package hooklog

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
	"server/util"
)

// AddHookLogService is
type AddHookLogService struct {
	IP       string
	Player   string
	Contents []string
}

// Add 写入HookLog
func (service *AddHookLogService) Add() serializer.Response {
	var hooks []model.Log
	for _, content := range service.Contents {
		hooks = append(hooks, model.Log{
			IP:        service.IP,
			Player:    service.Player,
			Content:   content,
			CreatedAt: time.Now(),
		})
	}

	// 写入
	for _, hook := range hooks {
		_, err := database.DB.Exec(`INSERT INTO pvzac.hook(ip,player,content,create_at)VALUES (?,?,?,?)`, hook.IP, hook.Player, hook.Content, hook.CreatedAt)
		if err != nil {
			util.Log().Panic("日志写入数据库失败", err)
		}
	}
	return serializer.BuildLogResponse(hooks)
}
