package hooklog

import (
	"server/database"
	"server/model"
	"server/serializer"
)

// ListHookLogService HookLog列表服务
type ListHookLogService struct {
}

// List HookLog
func (service *ListHookLogService) List() serializer.Response {
	var hookLog model.Log
	var hookLogs []model.Log

	rows, err := database.DB.Query(`SELECT * FROM hook ORDER BY id`)
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}
	for rows.Next() {
		err = rows.Scan(&hookLog.ID, &hookLog.IP, &hookLog.Player, &hookLog.Content, &hookLog.CreatedAt)
		if err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库操作错误",
				Error: err.Error(),
			}
		}
		hookLogs = append(hookLogs, hookLog)
	}
	return serializer.Response{
		Data: serializer.BuildLogs(hookLogs),
	}
}

// ListExp HookLog Exception
func (service *ListHookLogService) ListExp() serializer.Response {
	var hookLog model.Log
	var hookLogs []model.Log

	rows, err := database.DB.Query(`SELECT * FROM hook WHERE left(content,1)='!' ORDER BY id`)
	if err != nil {
		return serializer.Response{
			Code:  50000,
			Msg:   "数据库连接错误",
			Error: err.Error(),
		}
	}
	for rows.Next() {
		err = rows.Scan(&hookLog.ID, &hookLog.IP, &hookLog.Player, &hookLog.Content, &hookLog.CreatedAt)
		if err != nil {
			return serializer.Response{
				Code:  50000,
				Msg:   "数据库操作错误",
				Error: err.Error(),
			}
		}
		hookLogs = append(hookLogs, hookLog)
	}
	return serializer.Response{
		Data: serializer.BuildLogs(hookLogs),
	}
}
