package v1

import (
	"bytes"
	"io"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"server/service/antilog"
	"server/service/hooklog"
	"server/util"
)

// UploadLog Log上传接口
func UploadLog(c *gin.Context) {
	var addAntilogService antilog.AddAntiLogService
	var addHooklogService hooklog.AddHookLogService

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "Request Failed")
		return
	}

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		util.Log().Error("读取上传文件错误", err.Error())
	}

	if header.Filename == "Anti-log.txt" {
		// 传入 Antilog Contents
		var antiResult = &(addAntilogService).Contents
		for {
			line, err := buf.ReadString('\n')
			line = strings.TrimSpace(line)
			if err != nil {
				if err == io.EOF { // 读取结束
					break
				}
				util.Log().Error("按行读取文件错误", err)
			}
			if line != "" {
				*antiResult = append(*antiResult, line)
			}
		}

		// 传入IP和玩家名
		(&addAntilogService).IP = c.ClientIP()
		(&addAntilogService).Player = c.Request.Header.Get("PlayerName")
		_ = addAntilogService.Add()
		// ret, _ := json.Marshal(res)
		// util.Log().Println(string(ret) + "\n")
	} else {
		// 传入 Hooklog Contents
		var hookResult = &(addHooklogService).Contents
		for {
			line, err := buf.ReadString('\n')
			line = strings.TrimSpace(line)
			if err != nil {
				if err == io.EOF { // 读取结束
					break
				}
				util.Log().Error("按行读取文件错误", err)
			}
			if line != "" {
				*hookResult = append(*hookResult, line)
			}
		}

		// 传入IP和玩家名
		(&addHooklogService).IP = c.ClientIP()
		(&addHooklogService).Player = c.Request.Header.Get("PlayerName")
		_ = addHooklogService.Add()
		// ret, _ := json.Marshal(res)
		// util.Log().Println(string(ret) + "\n")
	}

	// 返回成功消息
	c.String(http.StatusCreated, "Upload Success")
}
