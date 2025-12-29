package app_window

import (
	"datathink.top/Waigo/internal"
	"datathink.top/Waigo/internal/app/app_window/window_controller"
	"datathink.top/Waigo/internal/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/wailsapp/wails/v3/pkg/application"
	"time"
)

// JSMustData 必要参数
func (awd *AppWindow) JSMustData(ctx *gin.Context) {
	//
	windowToken := common.MakeRandID()
	jsCallGoApi := common.InterfaceToString(internal.GetConfigMap("gin", "view_url")) + "/api/js_call_go"
	windowTokenTimer := common.GetTimeDate("YmdHis")
	appToken := common.MakeRandToken("20251229", 2*365*24*60*60)
	//
	js := `
		localStorage.setItem("window_token", "` + windowToken + `"); 
 		localStorage.setItem("app_token", "` + appToken + `");
		localStorage.setItem("js_call_go_api", "` + jsCallGoApi + `");
		localStorage.setItem("window_token_timer", "` + windowTokenTimer + `");
		`
	ctx.Header("Content-Type", "text/javascript; charset=utf-8")
	ctx.String(200, js)

	return
}

// WindowView Wails3视窗
func (awd *AppWindow) WindowView(app *application.App, window *application.WebviewWindow) {
	fmt.Println("WindowView=")

	// 启动周期定时器服务，实现向web端写入数据
	ticker := time.NewTicker(2 * time.Second) // 周期2s
	num := 0                                  // 计数
	for _ = range ticker.C {
		// 检查是否继续执行
		makeWindowTokenState := common.InterfaceToInt(internal.GetConfigSetter("make_window_token_state"))
		//fmt.Println("makeWindowTokenState=", makeWindowTokenState, num)
		if makeWindowTokenState >= 1 || num >= 20 { // 跳出
			fmt.Println("WindowToken已写入，跳出周期=", makeWindowTokenState, num)
			break
		}
		// 执行
		wct := window_controller.WindowController{}
		wct.ListGoRunJS("make_window_token", map[string]interface{}{})
		//
		num++
	}

	//

}
