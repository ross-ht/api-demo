# api-demo
## go版本 api的调用
说明：
**公共接口**
进入main.go中，无参数的请求直接调用
带参数的请求需要在params中输入json格式的参数 如：`{"symbol":"BTCUSDT",	"limit":"200"}`
**私有接口**
需要先在config.go中配置相关的api_key和sec_key
后续操作与公共接口相同

 
