package serializer

import "server/model"

// Count 日志统计序列化器
type Count struct {
	CountAll       uint `json:"count_all"`
	CountException uint `json:"count_exp"`
}

// BuildCount 序列化日志
func BuildCount(ct model.Count) Count {
	return Count{
		CountAll:       ct.CountAll,
		CountException: ct.CountException,
	}
}

// BuildCountResponse 序列化日志响应
func BuildCountResponse(ct model.Count) Response {
	return Response{
		Data: BuildCount(ct),
	}
}
