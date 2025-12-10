package window_controller

import (
	"datathink.top/Waigo/internal"
	"datathink.top/Waigo/internal/common"
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
	//
	var state = 0
	var msg = ""
	var content = map[string]interface{}{
		"key":       key,
		"data_dict": dataDict,
	}
	//
	if key == "test" || key == "Test" {
		//
		state = 1
		msg = "OK Test"
		content["dataTest"] = "test data 123"
	} else if key == "make_window_token" {
		windowToken := common.MakeRandID()
		//
		state = 1
		msg = "已生成"
		content["windowToken"] = windowToken
	} else {
		state = 0
		msg = "无白名单key"
	}
	//
	dataJSONString := common.MapToJsonString(map[string]interface{}{
		"state":   state,
		"msg":     msg,
		"content": content,
	}) // dataStr只能是string，可能是Emit函数的Bug
	//fmt.Println("ListGoRunJS=", key, dataDict, dataJSONString)
	internal.APP.Event.Emit(key, dataJSONString)
}
