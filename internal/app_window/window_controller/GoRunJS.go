package window_controller

import (
	"datathink.top.Waigo/internal"
	"datathink.top.Waigo/internal/common"
	"datathink.top.Waigo/internal/common/kits"
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
)

// 1️⃣/2️⃣先注册Event
func init() {
	application.RegisterEvent[string]("test")              // 注册Event
	application.RegisterEvent[string]("make_window_token") // 注册Event
}

// ListGoRunJS 2️⃣/2️⃣Go调用前端js对照表
// // GoRunJS
//
//	wct := window_controller.WindowController{}
//	wct.ListGoRunJS("test", map[string]interface{}{
//		"data1": 1,
//	})
//
// js接收
// // GoRunJS（Emit传递参数到前端js）
//
//	Events.On("test", (result) => {
//	 	console.log("GoEmitToJS=", result.data);
//	});
func (awd *WindowController) ListGoRunJS(key string, dataDict map[string]interface{}) {
	//
	if dataDict == nil {
		dataDict = map[string]interface{}{
			"key":       key,
			"ddataDict": dataDict,
			"error":     "空值",
		}
	}
	dataJSONString := common.MapToJsonString(dataDict) // dataStr只能是string，可能是Emit函数的Bug
	fmt.Println("ListGoRunJS=", key, dataDict, dataJSONString)
	//
	if key == "test" || key == "Test" {
		internal.APP.Event.Emit(key, dataJSONString)
	} else if key == "make_window_token" {
		secret := kits.Secret{}
		windowToken := secret.StringEncode("123", "")
		internal.APP.Event.Emit(key, windowToken)
	} else {
		fmt.Println("ListGoRunJS=else=", dataDict)
		internal.APP.Event.Emit(key, dataJSONString)
	}
}
