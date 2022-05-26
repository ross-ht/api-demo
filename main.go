package main

import (
	"fmt"

	spotList "github.com/ross-ht/api-demo/spot"
)

func main() {
	//说明：
	//在``中输入json格式的参数 如：{"symbol":"BTCUSDT",	"limit":"200"}
	//----------------------------------------------//
	// depthinfo := spotList.SpotMarketDepth(`{
	// 	"symbol":"BTCUSDT",
	// 	"limit":"200"
	// }`)
	// fmt.Println("返回信息:", depthinfo)

	// accountInfo := spotList.SpotmyTrade(`{
	// 	"symbol":"MXUSDT"
	// }`)
	// fmt.Println("返回信息:", accountInfo)
	subAccount := spotList.CreateSub(`{ "subAccount":"subAccount333","note":"no3"}`)
	fmt.Println("返回信息:", subAccount)

}
