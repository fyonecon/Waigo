package bootstrap

import (
	"datathink.top/Waigo/internal"
	"datathink.top/Waigo/internal/common"
	"runtime"
)

type CheckSYS struct {
}

// InitCheckSYS 系统运行条件检测
func InitCheckSYS() (int64, string, map[string]interface{}) {
	//
	state := 0
	msg := ""
	content := map[string]interface{}{
		"error": "",
	}

	// 检查逻辑CPU数量
	minCPUs := 2 // 个逻辑CPU
	numCPUs := runtime.NumCPU()
	cpuState := numCPUs >= minCPUs

	// 检测强制更新时间（这是软件及扩展更新的要求）
	startTime := common.InterfaceToInt(internal.GetConfigMap("sys", "appStateStartTime"))
	endTime := common.InterfaceToInt(internal.GetConfigMap("sys", "appStateEndTime"))
	nowTime := common.StringToInt(common.GetTimeDate("YmdHis"))
	timeState := (nowTime >= startTime) && (nowTime <= endTime)

	// 判断
	if cpuState && timeState {
		// 生成本次软件运行的唯一ID
		internal.RUNNNINGID = common.GoRandString(64, 128)
		//
		state = 1
		msg = "系统检测通告"
		content = map[string]interface{}{
			"error": "",
		}
	} else {
		state = 0
		msg = "系统运行条件不满足"
		content = map[string]interface{}{
			"error":   "",
			"numCPUs": numCPUs,
			"nowTime": nowTime,
		}
	}
	//
	return int64(state), msg, content
}
