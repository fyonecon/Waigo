package common

import (
	"datathink.top/internal"
)

// 设置一个会过期的临时数据

type CacheData struct{}

var cacheDataGap int64 = 2 // s，最小有效时间

// SetCacheData 新增一个数据
// key 要具有唯一性，timeout单位s（ >= cacheDataGap s）
func (cd *CacheData) SetCacheData(key string, mapData map[string]interface{}, timeout int64) bool {
	if timeout < cacheDataGap { // s
		timeout = cacheDataGap
	}
	cacheTime := GetTimeMS() + timeout*1000
	appClass := InterfaceToString(internal.ConfigMap["appClass"])
	// 新增或更新
	internal.SetGlobalMapInterface(appClass+key+"_cache_data", mapData)
	internal.SetGlobalMapInterface(appClass+key+"_cache_time", cacheTime) // ms

	return true
}

// GetCacheData 查询一个数据
func (cd *CacheData) GetCacheData(key string) (bool, map[string]interface{}) {
	appClass := InterfaceToString(internal.ConfigMap["appClass"])

	// 读取数据
	oldMap := internal.GetGlobalMapInterface(appClass + key + "_cache_data")
	oldMapCacheTime := internal.GetGlobalMapInterface(appClass + key + "_cache_time")
	_oldMap := InterfaceToMap(oldMap)
	cacheTime := InterfaceToInt(oldMapCacheTime)

	// 设置过期时间
	_nowTime := GetTimeMS()
	state := true
	//
	if _nowTime > cacheTime || cacheTime < cacheDataGap*1000 { // 过期
		_oldMap = map[string]interface{}{}
		state = false
		cd.DelCacheData(key)
	}

	return state, _oldMap
}

// DelCacheData 删除一个数据
func (cd *CacheData) DelCacheData(key string) bool {
	appClass := InterfaceToString(internal.ConfigMap["appClass"])
	//
	internal.DelGlobalMapInterface(appClass + key + "_cache_data")
	internal.DelGlobalMapInterface(appClass + key + "_cache_time")

	return true
}
