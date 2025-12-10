package window_controller

import (
	"datathink.top/Waigo/internal"
	"datathink.top/Waigo/internal/common"
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
	if key == "test" || key == "Test" {
		//
		state = 1
		msg = "TEST"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
		}
	} else if key == "stop_go_run_js_for_make_window_token" {
		stateNum := common.InterfaceToInt(internal.GetConfigSetter("make_window_token_state"))
		internal.SetConfigSetter("make_window_token_state", int64(stateNum+1))
		//
		state = 1
		msg = "Stop Emit"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
			"state_num": stateNum,
		}
	} else if key == "window_show" {
		internal.WINDOW.Show()
		internal.WINDOW.Focus()
		//
		state = 1
		msg = "Show"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
		}
	} else if key == "window_hide" {
		internal.WINDOW.Hide()
		//
		state = 1
		msg = "Hide"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
		}
	} else if key == "window_title" {
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
		content = map[string]interface{}{
			"key":       key,
			"title":     title,
			"data_dict": dataDict,
		}
	} else if key == "open_url_with_default_browser" {
		url := common.InterfaceToString(dataDict["url"])
		err := internal.APP.Browser.OpenURL(url)
		if err != nil {
			state = 0
			msg = "打开浏览器时报错"
			content = map[string]interface{}{
				"key":       key,
				"data_dict": dataDict,
			}
		} else {
			state = 1
			msg = "用默认浏览器打开"
			content = map[string]interface{}{
				"key":       key,
				"data_dict": dataDict,
			}
		}
	} else if key == "open_url_with_master_window" {
		url := common.InterfaceToString(dataDict["url"])
		internal.WINDOW.SetURL(url)
		//
		state = 1
		msg = "TEST"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
		}
	} else if key == "open_url_with_new_window" {
		//
		state = 0
		msg = "（未完成的功能）"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
		}
	} else {
		state = 0
		msg = "不在白名单的KEY"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
		}
	}
	return map[string]interface{}{
		"state":   state,
		"msg":     msg,
		"content": content,
	}
}
