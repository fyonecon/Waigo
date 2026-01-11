package kits

// 设置一个会过期的临时数据（软件开启时，软件关闭数据消失）

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type CacheData struct{}

var cacheDataGap int64 = 2 // s，最小有效时间

var cacheDataPrefix string = "waigo_" // 前缀

// SetData 新增一个数据
// key 要具有唯一性，timeout单位s（ >= cacheDataGap s）
func (cd *CacheData) SetData(key string, mapData map[string]interface{}, timeout int64) bool {
	cdf := CacheDataFunc{}

	if timeout < cacheDataGap { // s
		timeout = cacheDataGap
	}
	cacheTime := cdf.GetTimeMS() + timeout*1000
	// 新增或更新
	cdf.SetConfigSetter(key+"_cache_data", mapData)
	cdf.SetConfigSetter(key+"_cache_time", cacheTime) // ms

	return true
}

// GetData 查询一个数据
func (cd *CacheData) GetData(key string) (bool, map[string]interface{}) {
	cdf := CacheDataFunc{}

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
		cd.DelData(key)
	}

	return state, _oldMap
}

// DelData 删除一个数据
func (cd *CacheData) DelData(key string) bool {
	cdf := CacheDataFunc{}

	return cdf.DelConfigSetter(key+"_cache_data") && cdf.DelConfigSetter(key+"_cache_time")
}

//==================================================

type CacheDataFunc struct{}

// map[string]interface{} 全局配置（自定义）参数。读写
var syncMapInterface sync.Map

// SetConfigSetter 新增或更新
func (cdf *CacheDataFunc) SetConfigSetter(key string, value interface{}) interface{} {
	//
	syncMapInterface.Store(cacheDataPrefix+key, value)
	return value
}

// GetConfigSetter 读取
func (cdf *CacheDataFunc) GetConfigSetter(key string) interface{} {
	//
	back, ok := syncMapInterface.Load(cacheDataPrefix + key)
	if ok {
		return back
	} else {
		return nil
	}
}

// DelConfigSetter 删除
func (cdf *CacheDataFunc) DelConfigSetter(key string) bool {
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
func (cdf *CacheDataFunc) GetTimeMS() int64 {
	timer := time.Now()
	timeMS := int64(timer.UnixNano() / 1e6)
	return timeMS
}

// InterfaceToMap interface{}转map[string]interface{}
func (cdf *CacheDataFunc) InterfaceToMap(inter interface{}) map[string]interface{} {
	if inter == nil {
		return nil
	} else {
		return inter.(map[string]interface{})
	}
}

// InterfaceToInt interface{}，类似ValueInterfaceToInt
func (cdf *CacheDataFunc) InterfaceToInt(value interface{}) int64 {
	return cdf.StringToInt(fmt.Sprintf("%v", value))
}

// StringToInt string转int
func (cdf *CacheDataFunc) StringToInt(_str string) int64 {
	_int, err := strconv.ParseInt(_str, 10, 64) // string转int
	if err != nil {
		_int = 0
	}
	return _int
}
