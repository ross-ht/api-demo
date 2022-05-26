package spotList

import (
	"fmt"

	"github.com/ross-ht/api-demo/config"
	"github.com/ross-ht/api-demo/utils"
)

func SpotMarketDepth(jsonParams string) interface{} {
	caseUrl := "/depth"
	requestUrl := config.BASE_URL + caseUrl
	fmt.Println("requestUrl:", requestUrl)
	response := utils.PublicGet(requestUrl, jsonParams)
	return response
}
