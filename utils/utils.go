package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/go-resty/resty/v2"
)

//格式化参数字符串
func JsonToParamStr(jsonParams string) string {
	//转化json参数->参数字符串
	var paramsarr []string
	var arritem string
	m := make(map[string]string)
	err := json.Unmarshal([]byte(jsonParams), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("map:%v\n", m)
	i := 0
	for key, value := range m {

		arritem = fmt.Sprintf("%s=%s", key, value)
		paramsarr = append(paramsarr, arritem)
		i++
		fmt.Println("遍历：", i, "总共", len(m))
		if i > len(m) {
			break
		}
	}
	paramsstr := strings.Join(paramsarr, "&")
	return paramsstr
}

//定义公共get请求
func PublicGet(urlStr string, jsonParams string) interface{} {
	var path string
	if jsonParams == "" {
		path = urlStr
	} else {
		strParams := JsonToParamStr(jsonParams)
		path = urlStr + "?" + strParams
		fmt.Println("path:", path)
	}
	//创建请求
	client := resty.New()
	//发送请求
	resp, err := client.R().Get(path)

	if err != nil {
		log.Fatal("请求报错：", err)
	}

	// fmt.Println("Response Info:", resp)
	return resp
}
