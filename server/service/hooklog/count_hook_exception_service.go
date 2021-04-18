package hooklog

import (
	"server/database"
	"server/model"
	"server/serializer"
)

// CountHookExceptionService Hook异常统计服务
type CountHookExceptionService struct {
}

func (ce *CountHookExceptionService) Count() serializer.Response {
	var count model.Count

	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM hook`).Scan(&count.CountAll)
	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM hook WHERE left(content,1)='!'`).Scan(&count.CountException)

	return serializer.BuildCountResponse(count)
}
