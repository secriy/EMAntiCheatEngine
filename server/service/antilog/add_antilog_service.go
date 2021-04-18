package antilog

import (
	"time"

	"server/database"
	"server/model"
	"server/serializer"
	"server/util"
)

// AddAntiLogService is
type AddAntiLogService struct {
	IP       string
	Player   string
	Contents []string
}

// Add 写入AntiLog
func (service *AddAntiLogService) Add() serializer.Response {
	var antis []model.Log
	for _, content := range service.Contents {
		antis = append(antis, model.Log{
			IP:        service.IP,
			Player:    service.Player,
			Content:   content,
			CreatedAt: time.Now(),
		})
	}

	// 写入
	for _, anti := range antis {
		_, err := database.DB.Exec(`INSERT INTO pvzac.anti(ip,player,content,create_at)VALUES (?,?,?,?)`, anti.IP, anti.Player, anti.Content, anti.CreatedAt)
		if err != nil {
			util.Log().Println(anti.Content)

			util.Log().Panic("日志写入数据库失败", err)
		}
	}
	return serializer.BuildLogResponse(antis)
}
