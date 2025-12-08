package internal

import (
	"fmt"
	"sync"
)

// map[string]interface{} 全局配置（自定义）参数。读写
var syncMapInterface sync.Map

// SetConfigSetter 新增或更新
func SetConfigSetter(key string, value interface{}) interface{} {
	appClass := fmt.Sprintf("%v", GetConfigMap("app", "appClass"))
	//
	syncMapInterface.Store(appClass+key, value)
	return value
}

// GetConfigSetter 读取
func GetConfigSetter(key string) interface{} {
	appClass := fmt.Sprintf("%v", GetConfigMap("app", "appClass"))
	//
	back, ok := syncMapInterface.Load(appClass + key)
	if ok {
		return back
	} else {
		return nil
	}
}

// DelConfigSetter 删除
func DelConfigSetter(key string) bool {
	appClass := fmt.Sprintf("%v", GetConfigMap("app", "appClass"))
	//
	syncMapInterface.Delete(appClass + key)
	_, ok := syncMapInterface.Load(appClass + key)
	if ok {
		return true
	} else {
		return false
	}
}

// 初始化数据：全局配置（自定义）参数。读、写。动态数据。
// 读取common.InterfaceToString(internal.GetConfigSetterInterface("xxx"))
// 设置internal.SetConfigSetter("xxx", any)
func init() {
	// 子程序参数
	SetConfigSetter("sys_start_time", int64(0))
	SetConfigSetter("max_file_size", int64(20))
}
