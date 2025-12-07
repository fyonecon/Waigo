package internal

import (
	"fmt"
	"sync"
)

// map[string]interface{} 全局配置（自定义）参数。读写
var syncMapInterface sync.Map

// SetGlobalMapInterface 新增或更新
func SetGlobalMapInterface(key string, value interface{}) {
	syncMapInterface.Store(key, value)
}

// GetGlobalMapInterface 读取
func GetGlobalMapInterface(key string) interface{} {
	back, _ := syncMapInterface.Load(key)
	return back
}

// DelGlobalMapInterface 删除
func DelGlobalMapInterface(key string) interface{} {
	syncMapInterface.Delete(key)
	back, _ := syncMapInterface.Load(key)
	return back
}

// 初始化数据：全局配置（自定义）参数。读、写。动态数据。
// 读取common.InterfaceToString(internal.GetGlobalMapInterface("xxx"))
// 设置internal.SetGlobalMapInterface("xxx", any)
func init() {
	appClass := fmt.Sprintf("%v", ConfigMap["appClass"])
	// 子程序参数
	SetGlobalMapInterface(appClass+"switch_sys_up_time", int64(0))
}
