package antilog

import (
	"server/database"
	"server/model"
	"server/serializer"
)

// CountAntiExceptionService Anti异常统计服务
type CountAntiExceptionService struct {
}

func (ce *CountAntiExceptionService) Count() serializer.Response {
	var count model.Count

	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM anti`).Scan(&count.CountAll)
	_ = database.DB.QueryRow(`SELECT COUNT(*) FROM anti WHERE left(content,1)='!'`).Scan(&count.CountException)

	return serializer.BuildCountResponse(count)
}
