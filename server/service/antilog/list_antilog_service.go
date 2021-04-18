package antilog

import (
	"server/database"
	"server/model"
	"server/serializer"
)

// ListAntiLogService AntiLog列表服务
type ListAntiLogService struct {
}

// List AntiLog
func (service *ListAntiLogService) List() serializer.Response {
	var antiLog model.Log
	var antiLogs []model.Log

	rows, err := database.DB.Query(`SELECT * FROM anti ORDER BY id`)
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}
	for rows.Next() {
		err = rows.Scan(&antiLog.ID, &antiLog.IP, &antiLog.Player, &antiLog.Content, &antiLog.CreatedAt)
		if err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库操作错误",
				Error: err.Error(),
			}
		}
		antiLogs = append(antiLogs, antiLog)
	}
	return serializer.Response{
		Data: serializer.BuildLogs(antiLogs),
	}
}

// List AntiLog Exception
func (service *ListAntiLogService) ListExp() serializer.Response {
	var antiLog model.Log
	var antiLogs []model.Log

	rows, err := database.DB.Query(`SELECT * FROM anti WHERE left(content,1)='!' ORDER BY id`)
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}
	for rows.Next() {
		err = rows.Scan(&antiLog.ID, &antiLog.IP, &antiLog.Player, &antiLog.Content, &antiLog.CreatedAt)
		if err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库操作错误",
				Error: err.Error(),
			}
		}
		antiLogs = append(antiLogs, antiLog)
	}
	return serializer.Response{
		Data: serializer.BuildLogs(antiLogs),
	}
}
