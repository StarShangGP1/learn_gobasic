package format

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/ghodss/yaml"
	"google.golang.org/protobuf/proto"
	"learn_gobasic/pkg/apis"
	"log"
	"os"
)

type Record struct {
	filePath         string
	yamlFilePath     string
	protobufFilePath string
}

func (r *Record) NewRecord(filePath string) *Record {
	return &Record{
		filePath:         filePath,
		yamlFilePath:     filePath + ".yaml",
		protobufFilePath: filePath + ".proto.base64",
	}
}

func (r *Record) SavePInfo(pInfo *apis.PersonalInformation) error {
	{
		data, err := json.Marshal(pInfo)
		if err != nil {
			fmt.Println("marshal 出错：", err)
			return err
		}
		if err := r.WriteFileWithJson(data); err != nil {
			log.Println("写入JSON时出错：", err)
			return err
		}
	}
	{
		data, err := yaml.Marshal(pInfo)
		if err != nil {
			fmt.Println("marshal 出错：", err)
			return err
		}
		if err := r.WriteFileWithYaml(data); err != nil {
			log.Println("写入JSON时出错：", err)
			return err
		}
	}
	{
		data, err := proto.Marshal(pInfo)
		if err != nil {
			fmt.Println("marshal 出错：", err)
			return err
		}
		if err := r.WriteFileWithProtobuf(data); err != nil {
			log.Println("写入JSON时出错：", err)
			return err
		}
	}

	return nil
}

func (r *Record) WriteFileWithJson(data []byte) error {
	file, err := os.OpenFile(r.filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("无法打开文件", r.filePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.Write(append(data, '\n'))
	if err != nil {
		return err
	}
	return nil
}

func (r *Record) WriteFileWithYaml(data []byte) error {
	file, err := os.OpenFile(r.yamlFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("无法打开文件", r.yamlFilePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()
	tmpData := append([]byte("---\n"), data...)
	_, err = file.Write(tmpData)
	if err != nil {
		return err
	}
	return nil
}

func (r *Record) WriteFileWithProtobuf(data []byte) error {
	file, err := os.OpenFile(r.protobufFilePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		fmt.Println("无法打开文件", r.yamlFilePath, "错误信息是：", err)
		os.Exit(1)
	}
	defer file.Close()

	tmpData := []byte(base64.StdEncoding.EncodeToString(data))
	_, err = file.Write(tmpData)
	if err != nil {
		return err
	}
	return nil
}
