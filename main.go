package main

import (
	"fmt"

	spotList "github.com/ross-ht/api-demo/spot"
)

func main() {
	//输入json格式的参数
	depthinfo := spotList.SpotMarketDepth(`{
		"symbol":"BTCUSDT",
		"limit":"200"
	}`)
	fmt.Println("返回信息:", depthinfo)

}
