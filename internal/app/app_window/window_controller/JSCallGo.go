package window_controller

import (
	"datathink.top/Waigo/internal"
	"datathink.top/Waigo/internal/common"
	"datathink.top/Waigo/internal/common/kits"
	"fmt"
)

// ListJSCallGo js调用go方法的对照表
func (wct *WindowController) ListJSCallGo(key string, dataDict map[string]interface{}) map[string]interface{} {
	//
	var state = 0
	var msg = ""
	var content = map[string]interface{}{
		"key":       key,
		"data_dict": dataDict,
	}
	// 必要 ===
	if key == "test" || key == "Test" { // data_dict={test:""}
		//
		state = 1
		msg = "TEST"
		content["test"] = "test 123"
	} else if key == "stop_go_run_js_for_make_window_token" { // 停止go写入参数信号。 data_dict={}
		stateNum := common.InterfaceToInt(internal.GetConfigSetter("make_window_token_state"))
		internal.SetConfigSetter("make_window_token_state", int64(stateNum+1))
		//
		state = 1
		msg = "Stop Emit"
		content["state_num"] = stateNum
	} else if key == "make_window_token" { //。 data_dict={}
		windowToken := common.MakeRandID()
		jsCallGoApi := common.InterfaceToString(internal.GetConfigMap("gin", "view_url")) + "/api/js_call_go"
		//
		state = 1
		msg = "已生成"
		content["windowToken"] = windowToken
		content["jsCallGoApi"] = jsCallGoApi
	} else if key == "window_show" { // data_dict={}
		internal.WINDOW.Show()
		internal.WINDOW.Focus()
		//
		state = 1
		msg = "Show"
	} else if key == "window_hide" { // data_dict={}
		internal.WINDOW.Hide()
		//
		state = 1
		msg = "Hide"
	} else if key == "change_window_size" { // data_dict={size_tag: "", width: 0, height: 0}
		nowScreenInfo, _ := internal.APP.Window.Current().GetScreen() // 获取当前显示器的信息
		//fmt.Println("change_window_size=", nowScreenInfo, nowScreenInfo.Name, nowScreenInfo.ID, nowScreenInfo.IsPrimary, nowScreenInfo.Size.Width, nowScreenInfo.Size.Height)
		screenWidth := nowScreenInfo.Size.Width
		screenHeight := nowScreenInfo.Size.Height
		// 同init_window参数
		initWidth := 960
		initHeight := 700
		//
		width := 0
		height := 0
		sizeTag := dataDict["size_tag"]
		if sizeTag == "size_init" {
			width = initWidth
			height = initHeight
			internal.WINDOW.SetSize(width, height)
			// 居中
			//x := 0
			//y := 0
			//internal.WINDOW.SetPosition(x, y)
			internal.WINDOW.Center()
		} else if sizeTag == "size_full_height" {
			width = screenWidth / 3 * 2
			if width < initWidth {
				width = initWidth
			}
			height = screenHeight
			internal.WINDOW.SetSize(width, height)
			// 居中
			//x := (screenWidth - width) / 2
			//y := 0
			//internal.WINDOW.SetPosition(x, y)
			internal.WINDOW.Center()
		} else if sizeTag == "size_full_window" {
			width = screenWidth
			height = screenHeight
			internal.WINDOW.SetSize(width, height)
			// 居中
			//x := 0
			//y := 0
			//internal.WINDOW.SetPosition(x, y)
			internal.WINDOW.Center()
		} else if sizeTag == "size_full_screen" {
			internal.WINDOW.Fullscreen()
		} else if sizeTag == "size_center" {
			internal.WINDOW.Center()
		} else {
			// 自定义尺寸
			width = int(common.InterfaceToInt(dataDict["width"]))
			height = int(common.InterfaceToInt(dataDict["height"]))
			// 校验
			if width < 320 || width > screenWidth {
				width = screenWidth
			}
			if height < 320 || height > screenHeight {
				height = screenHeight
			}
			internal.WINDOW.SetSize(width, height)
			// 居中
			//x := (screenWidth - width) / 2
			//y := (screenHeight - height) / 2
			//internal.WINDOW.SetPosition(x, y)
			internal.WINDOW.Center()
		}

		//
		state = 1
		msg = "Size"
		content["size_tag"] = sizeTag
		content["window_size"] = []int{width, height}
		content["screen_size"], _ = internal.WINDOW.Size()
	} else if key == "window_title" { // data_dict={title:""}
		title := common.InterfaceToString(dataDict["title"])
		maxLen := 8
		if len(title) == 0 {
			title = "(null)"
		} else if len(title) >= maxLen {
			title = title[:maxLen] + ".."
		}
		internal.WINDOW.SetTitle(title)
		//
		state = 1
		msg = "Title"
		content["title"] = title
	} else if key == "open_url_with_default_browser" { // 用默认浏览器打开链接。 data_dict={}
		url := common.InterfaceToString(dataDict["url"])
		err := internal.APP.Browser.OpenURL(url)
		//
		if err != nil {
			state = 0
			msg = "打开浏览器时报错"
		} else {
			state = 1
			msg = "用默认浏览器打开"
		}
	} else if key == "open_url_with_master_window" { // data_dict={}
		url := common.InterfaceToString(dataDict["url"])
		internal.WINDOW.SetURL(url)
		//
		state = 1
		msg = "TEST"
	} else if key == "open_url_with_new_window" { // 用新窗口打开一个链接。 data_dict={}
		//
		state = 0
		msg = "（未完成的功能）"
	} else if key == "set_data" { // 新增或更新本地数据。 data_dict={data_key:"", data_value:"", data_timeout_s:3600}
		fmt.Println("SetData=====", dataDict)
		_key := common.InterfaceToString(dataDict["data_key"])
		_value := common.InterfaceToString(dataDict["data_value"])
		_timeoutS := common.InterfaceToInt(dataDict["data_timeout_s"])
		if _timeoutS <= 5*60 { // 最短5min
			_timeoutS = 5 * 60
		}
		db := kits.LocalDB{}
		data := db.SetData(_key, _value, _timeoutS)
		//
		state = 1
		msg = "OK"
		content["data"] = data
	} else if key == "get_data" { // 读取本地数据。 data_dict={data_key:""}
		db := kits.LocalDB{}
		_key := common.InterfaceToString(dataDict["data_key"])
		_value, _state := db.GetData(_key)
		//
		if _state == -1 {
			state = 0
			msg = "Not Set"
			content["data"] = ""
		} else {
			state = 1
			msg = "OK"
			content["data"] = _value
		}
	} else if key == "del_data" { // 删除本地数据 data_dict={data_key:""}
		_key := common.InterfaceToString(dataDict["data_key"])
		//
		db := kits.LocalDB{}
		_state := db.DelData(_key)
		if _state == -1 {
			state = 0
			msg = "No file"
		} else {
			state = 1
			msg = "Del OK"
		}
	} else {
		state = 0
		msg = "不在白名单的KEY"
	}
	return map[string]interface{}{
		"state":   state,
		"msg":     msg,
		"content": content,
	}
}
