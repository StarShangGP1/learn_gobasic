package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "class/class22/file.txt"
	//readFile(filePath)
	//writeFile(filePath)
	writeFileWithAppend(filePath)
}

func writeFileWithAppend(filePath string) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.Write([]byte("this is first line11---"))
	_, err = file.Write([]byte("this is first line11\n"))
	_, err = file.WriteString("第二行内容111\n")
	//_, err = file.WriteAt([]byte("CHANGED111"), 0)
	fmt.Println(err)
}

func writeFile(filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write([]byte("this is first line---"))
	_, err = file.Write([]byte("this is first line\n"))
	_, err = file.WriteString("第二行内容\n")
	_, err = file.WriteAt([]byte("CHANGED"), 0)
	fmt.Println(err)
}

func readFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("无法打开文件", filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	b := make([]byte, 10)
	var n int
	for i := 0; i < 2; i++ {
		n, err = file.Read(b)
		if err != nil {
			fmt.Println("无法读取文件：", err)
			os.Exit(2)
		}
		fmt.Println("n的大小：", n)
		fmt.Println("读出的内容：", string(b[:n]))
		file.Seek(3, io.SeekCurrent)
	}
}
