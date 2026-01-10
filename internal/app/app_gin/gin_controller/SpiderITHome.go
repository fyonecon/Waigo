package gin_controller

import (
	"datathink.top/Waigo/internal/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (gcl *GinController) SpiderITHome(ctx *gin.Context) {
	itURL := "https://m.ithome.com"
	var array []map[string]interface{}

	//
	info := map[string]interface{}{
		"news_index":   0,
		"news_id":      0,
		"news_href":    "",
		"news_title":   common.URLEncode(""),
		"news_time":    "",
		"comments_num": 0,
	}
	array = append(array, info)

	//
	ctx.JSON(http.StatusOK, gin.H{
		"state": 1,
		"msg":   "-功能还没写-",
		"content": map[string]interface{}{
			"it_url": itURL,
			"array":  map[string]interface{}{},
		},
	})

	return
}
