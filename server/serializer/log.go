package serializer

import "server/model"

// Log 日志序列化器
type Log struct {
	ID        uint64 `json:"id"`
	IP        string `json:"ip"`
	Player    string `json:"player"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"created_at"`
}

// BuildLog 序列化日志
func BuildLog(log model.Log) Log {
	return Log{
		ID:        log.ID,
		IP:        log.IP,
		Player:    log.Player,
		Content:   log.Content,
		CreatedAt: log.CreatedAt.Unix(),
	}
}

// BuildLogs 序列化日志列表
func BuildLogs(items []model.Log) []Log {
	var logs []Log
	for _, item := range items {
		logs = append(logs, BuildLog(item))
	}
	return logs
}

// BuildLogResponse 序列化日志响应
func BuildLogResponse(log []model.Log) Response {
	return Response{
		Data: BuildLogs(log),
	}
}
