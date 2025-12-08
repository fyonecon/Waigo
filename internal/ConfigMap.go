package internal

import (
	"github.com/wailsapp/wails/v3/pkg/application"
)

// 全局变量
var APP *application.App
var WINDOW *application.WebviewWindow

func GetConfigMap(group string, key string) interface{} {
	//
	var inter = ConfigMap[group]
	if inter != nil {
		config := inter.(map[string]interface{}) // interface转map
		return config[key]
	} else {
		return nil
	}
}

// ConfigMap 全局配置（固定）参数。只读。静态数据。
// appConfig := common.InterfaceToMap(internal.ConfigMap["app"]); appClass := common.InterfaceToString(appConfig["appClass"])
var ConfigMap map[string]interface{} = map[string]interface{}{
	"app": map[string]interface{}{
		"appName":    "Waigo",         // AppName，推荐使用英文
		"appClass":   "Waigo_",        // 必须为string，且必须唯一，推荐使用英文
		"appRights":  "Datathink.Top", // AppRights，也是软件安装默认路径
		"appVersion": "1.0.0",         // AppVersion
	},
	"sys": map[string]interface{}{
		"appStateStartTime": 20251201010101,                // 最早时间
		"appStateEndTime":   20321201010101,                // 截止时间，一个版本：7年
		"cachePath":         "top.datathink.Wailgo/Cache/", // 缓存文件夹，必须/结尾
		"localPath":         "top.datathink.Wailgo/Files/", // 缓存文件夹，必须/结尾
	},
	"gin": map[string]interface{}{
		"ginDebug":    "on",   // on 显示Gin日志，off 不显示Gin日志
		"ginPort":     "9850", // 内网端口，默认：9850
		"maxFileSize": "32",   // 最大上传文件体积，GB
	},
	"wails": map[string]interface{}{
		"debug": "off",
	},
}
