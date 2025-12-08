package bootstrap

import (
	"runtime"
)

type CheckSYS struct {
}

// InitCheckSYS 系统运行条件检测
func InitCheckSYS() (int64, string, map[string]interface{}) {
	// 检查逻辑CPU数量
	minCPUs := 2 // 个逻辑CPU
	numCPUs := runtime.NumCPU()
	cpuState := numCPUs >= minCPUs
	//
	state := 0
	msg := ""
	content := map[string]interface{}{
		"error": "",
	}
	//
	if cpuState {
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
		}
	}
	//
	return int64(state), msg, content
}
