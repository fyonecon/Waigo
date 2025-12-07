package app_window

// ListJSCallGo js调用go方法的对照表
func (awd *AppWindow) ListJSCallGo(key string, dataDict map[string]interface{}) map[string]interface{} {
	//
	var state = 0
	var msg = ""
	var content = map[string]interface{}{
		"key":       key,
		"data_dict": dataDict,
		"result":    "",
	}
	// 必要 ===
	if key == "test" {
		state = 1
		msg = "TEST"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
			"result":    "TEST is OK",
		}
	} else if key == "window_show" {
		state = 1
		msg = "Show"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
			"result":    "",
		}
	} else if key == "window_hide" {
		state = 1
		msg = "Hide"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
			"result":    "",
		}
	} else if key == "window_title" {
		state = 1
		msg = "Title"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
			"result":    "",
		}
	} else if key == "open_url_with_default_browser" {
		state = 1
		msg = "TEST"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
			"result":    "",
		}
	} else if key == "open_url_with_new_window" {
		state = 1
		msg = "TEST"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
			"result":    "",
		}
	} else {
		state = 0
		msg = "不在白名单的KEY"
		content = map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
			"result":    "Others",
		}
	}
	return map[string]interface{}{
		"state":   state,
		"msg":     msg,
		"content": content,
	}
}
