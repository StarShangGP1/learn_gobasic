package global

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	FilePath       string
	FilePathRecord string
)

//func init() {
//	FilePath = "/Users/draken/Documents/chat-project/user.json"
//	FilePathRecord = "/Users/draken/Documents/chat-project/record.json"
//}

// 生成存储文件
func init() {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]

	FilePath = fmt.Sprintf("%s/user.json", ret)
	FilePathRecord = fmt.Sprintf("%s/record.json", ret)

	if _, err := os.Stat(FilePath); err != nil {
		_, err = os.Create(FilePath)
		if err != nil {
			fmt.Println("创建文件失败")
		}
	}

	if _, err := os.Stat(FilePathRecord); err != nil {
		_, err = os.Create(FilePathRecord)
		if err != nil {
			fmt.Println("创建文件失败")
		}
	}
}
