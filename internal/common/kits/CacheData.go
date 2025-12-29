package kits

// 设置一个会过期的临时数据（软件开启时）

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type CacheData struct{}

var cacheDataGap int64 = 2 // s，最小有效时间

var cacheDataPrefix string = "waigo_" // 前缀

// SetCacheData 新增一个数据
// key 要具有唯一性，timeout单位s（ >= cacheDataGap s）
func (cd *CacheData) SetCacheData(key string, mapData map[string]interface{}, timeout int64) bool {
	cdf := CacheDataFuncs{}

	if timeout < cacheDataGap { // s
		timeout = cacheDataGap
	}
	cacheTime := cdf.GetTimeMS() + timeout*1000
	// 新增或更新
	cdf.SetConfigSetter(key+"_cache_data", mapData)
	cdf.SetConfigSetter(key+"_cache_time", cacheTime) // ms

	return true
}

// GetCacheData 查询一个数据
func (cd *CacheData) GetCacheData(key string) (bool, map[string]interface{}) {
	cdf := CacheDataFuncs{}

	// 读取数据
	oldMap := cdf.GetConfigSetter(key + "_cache_data")
	oldMapCacheTime := cdf.GetConfigSetter(key + "_cache_time")
	_oldMap := cdf.InterfaceToMap(oldMap)
	cacheTime := cdf.InterfaceToInt(oldMapCacheTime)

	// 设置过期时间
	_nowTime := cdf.GetTimeMS()
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
	cdf := CacheDataFuncs{}

	return cdf.DelConfigSetter(key+"_cache_data") && cdf.DelConfigSetter(key+"_cache_time")
}

//==================================================

type CacheDataFuncs struct{}

// map[string]interface{} 全局配置（自定义）参数。读写
var syncMapInterface sync.Map

// SetConfigSetter 新增或更新
func (cdf *CacheDataFuncs) SetConfigSetter(key string, value interface{}) interface{} {
	//
	syncMapInterface.Store(cacheDataPrefix+key, value)
	return value
}

// GetConfigSetter 读取
func (cdf *CacheDataFuncs) GetConfigSetter(key string) interface{} {
	//
	back, ok := syncMapInterface.Load(cacheDataPrefix + key)
	if ok {
		return back
	} else {
		return nil
	}
}

// DelConfigSetter 删除
func (cdf *CacheDataFuncs) DelConfigSetter(key string) bool {
	//
	syncMapInterface.Delete(cacheDataPrefix + key)
	_, ok := syncMapInterface.Load(cacheDataPrefix + key)
	if ok {
		return true
	} else {
		return false
	}
}

// GetTimeMS 获取毫秒时间戳，ms
func (cdf *CacheDataFuncs) GetTimeMS() int64 {
	timer := time.Now()
	timeMS := int64(timer.UnixNano() / 1e6)
	return timeMS
}

// InterfaceToMap interface{}转map[string]interface{}
func (cdf *CacheDataFuncs) InterfaceToMap(inter interface{}) map[string]interface{} {
	if inter == nil {
		return nil
	} else {
		return inter.(map[string]interface{})
	}
}

// InterfaceToInt interface{}，类似ValueInterfaceToInt
func (cdf *CacheDataFuncs) InterfaceToInt(value interface{}) int64 {
	return cdf.StringToInt(fmt.Sprintf("%v", value))
}

// StringToInt string转int
func (cdf *CacheDataFuncs) StringToInt(_str string) int64 {
	_int, err := strconv.ParseInt(_str, 10, 64) // string转int
	if err != nil {
		_int = 0
	}
	return _int
}
