# api-demo
## v3 现货 api的调用
说明：
* **公共接口**
    1. 进入main.go中，无参数的请求直接调用
    2. 带参数的请求需要在params中输入json格式的参数 如：`{"symbol":"BTCUSDT",	"limit":"200"}`
* **私有接口**
    1. 需要先在config.go中配置相关的api_key和sec_key
    2. 后续操作与公共接口相同

## 现货ws的调用
**说明：**
1. 进入publicws.go,在payload参数中输入相应的json格式的request payload
2. 私有频道用privateWs.go进行订阅
