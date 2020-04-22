package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"
)

type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	Username string `ini:"username"`
	Password string `ini:"password"`
}

type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func loadIni(fileName string, data interface{}) (err error) {
	t := reflect.TypeOf(data)

	if t.Kind() != reflect.Ptr {
		err = fmt.Errorf("")
		return
	}

	b, _ := ioutil.ReadFile(fileName)
	//string(b)//将文件内容转换成字符串
	lineSlice := strings.Split(string(b), "\n")
	//fmt.Printf("%#v\n", lineSlice)
	var structName string
	for _, line := range lineSlice {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			sectionName := strings.TrimSpace(line[1 : len(line)-1])
			println(sectionName)
			//v := reflect.ValueOf(data)
			for i := 0; i < t.Elem().NumField(); i++ {
				field := t.Elem().Field(i)
				println(field.Tag.Get("ini"))
				if sectionName == field.Tag.Get("ini") {
					structName = field.Name
					fmt.Printf("找到%s对应的结构体%s", sectionName, structName)
					break
				}
			}
		} else {
			return nil
		}
	}
	return nil
}
func main() {
	var cfg MysqlConfig
	_ = loadIni(`E:\data\go\go\src\go_learn\day06\ini_demo\conf.ini`, &cfg)
	fmt.Println(cfg.Address, cfg.Port)
}
