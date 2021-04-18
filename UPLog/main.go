package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/fsnotify.v1"
)

func main() {
	// 参数：日志路径
	fileWatch("C:\\Anti_log", "http://127.0.0.1:3000/api/v1/upload")
}

// FileWatch 是监听日志目录的函数
func fileWatch(path string, url string) {
	// 创建监听器
	watcher, _ := fsnotify.NewWatcher()
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if path+"\\Anti-log.txt" == event.Name && event.Op&fsnotify.Write == fsnotify.Write {
					if upload(path, "\\Anti-log.txt", url) == true {
						err := os.Remove(path + "\\Anti-log.txt")
						if err != nil {
							fmt.Println(err)
							return
						}
					}
				} else if path+"\\hook_log.txt" == event.Name && event.Op&fsnotify.Write == fsnotify.Write {
					if upload(path, "\\hook_log.txt", url) == true {
						err := os.Remove(path + "\\hook_log.txt")
						if err != nil {
							fmt.Println(err)
							return
						}
					}
				}
			case _, ok := <-watcher.Errors:
				if !ok {
					return
				}
			}
		}
	}()

	// 指定监听目录
	_ = watcher.Add(path)
	<-done
}

func upload(path string, filename string, url string) bool {
	// 读文件
	file, _ := os.Open(path + filename)
	defer file.Close()

	// 读玩家名
	name, _ := ioutil.ReadFile(path + "\\username.txt")

	// 创建表单
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	formFile, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))

	// 写入表单
	_, _ = io.Copy(formFile, file)
	writer.Close()

	// 发送请求
	client := &http.Client{}
	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", writer.FormDataContentType())
	req.Header.Add("PlayerName", string(name))

	res, _ := client.Do(req)

	if res == nil {
		return false
	}
	defer res.Body.Close()
	text, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(text))

	if string(text) == "Upload Success" {
		return true
	}
	return false
}
