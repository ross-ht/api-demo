package main

import (
	"encoding/json"
	"fmt"
	"log"
	_ "net/url"
	"strings"

	"github.com/go-resty/resty/v2"
)

func jsonTostr(jsonparams string) interface{} {
	//转化json参数->参数字符串
	var paramsarr []string
	var arritem string
	m := make(map[string]interface{})
	err := json.Unmarshal([]byte(jsonparams), &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("map:%v\n", m)
	i := 0
	for key, value := range m {

		arritem = fmt.Sprintf("%s=%s", key, value)
		paramsarr = append(paramsarr, arritem)
		i++
		fmt.Println(i, len(m))
		if i > len(m) {
			break
		}

	}
	// paramsarr = append(paramsarr)
	// fmt.Println(paramsarr)
	paramsstr := strings.Join(paramsarr, "&")
	return paramsstr
}
func main() {
	var jsonstr string = `{
		"symbol":"BTCUSDT",
        "limit":"200"}`
	var paramsObj interface{} = jsonTostr(jsonstr)
	fmt.Println(paramsObj)

	var baseurl string = "https://api.mexc.com"
	var url string = "/api/v3/depth"

	var path string = fmt.Sprintf("%s%s?%s", baseurl, url, paramsObj)

	client := resty.New()

	resp, err := client.R().Get(path)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Response Info:", resp)

}
