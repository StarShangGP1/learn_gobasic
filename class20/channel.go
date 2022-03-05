package main

import (
	"fmt"
	"time"
)

type Downloader struct {
	fileNameCh chan string
	finishedCh chan string
}

func (d *Downloader) Download() {
	for fileName := range d.fileNameCh {
		fmt.Println("开始下载文件：", fileName)
		time.Sleep(1 * time.Second)
		fmt.Println("开始处理文件：", fileName)
		time.Sleep(10 * time.Millisecond)
		fmt.Println("保存文件：", fileName)
		d.finishedCh <- fileName
	}
}

func ChannelStudy() {
	fileNameCh := make(chan string)
	finishedCh := make(chan string)

	workerCount := 50

	// step 1
	// 创建 50 个 goroutine 去 下载文件
	for i := 0; i < workerCount; i++ {
		go func() {
			downloader := &Downloader{
				fileNameCh: fileNameCh,
				finishedCh: finishedCh,
			}
			downloader.Download()
		}()
	}

	// step 2
	fileNum := 100
	fileNames := make([]string, 0, fileNum)
	// 创建数组
	for i := 0; i < fileNum; i++ {
		fileNames = append(fileNames, fmt.Sprintf("file_%d.txt", i))
	}

	// step 4
	// 接收 fileNameCh 传过来的文件
	finishedFileCount := 0
	go func() {
		for finishedFile := range finishedCh {
			fmt.Println("文件：", finishedFile, "处理完毕")
			finishedFileCount++
			if finishedFileCount == fileNum {
				close(finishedCh)
			}
		}
	}()

	// step 3
	//把数组里的文件传给 chan
	for _, fileItem := range fileNames {
		fileNameCh <- fileItem
	}
	close(fileNameCh)

	fmt.Println("所有文件处理完毕，结束")

}
