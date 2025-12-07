package bootstrap

type CheckSYS struct {
}

// InitCheckSYS 系统运行条件检测
func InitCheckSYS() (int64, string, map[string]interface{}) {
	state := 1
	msg := "ok"
	content := map[string]interface{}{
		"error": "",
	}
	return int64(state), msg, content
}
