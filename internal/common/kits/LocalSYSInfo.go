package kits

import (
	"os"
	"os/user"
	"runtime"
	"strconv"
	"strings"
)

// 本地系统基本信息

type LocalSYSInfo struct {
}

// GetSYSInfo 获取系统信息（仅限 win、mac、Linux）
func (lsi *LocalSYSInfo) GetSYSInfo() map[string]string {
	var locale string = "en_US.UTF-8"
	var platform string = ""
	var username string = ""
	var homeDir string = ""
	var numCPUs int = runtime.NumCPU()

	//
	userInfo, _ := user.Current()
	username = userInfo.Username
	homeDir = userInfo.HomeDir

	// 根据操作系统获取
	switch runtime.GOOS {
	case "windows":
		locale = os.Getenv("LANG")
		if locale == "" {
			// Windows 使用其他变量
			locale = "en_US.UTF-8"
		}
		platform = "win"
	case "darwin":
		// 尝试多个环境变量
		for _, env := range []string{"LANG", "LANGUAGE", "LC_ALL"} {
			locale = os.Getenv(env)
			if locale != "" {
				break
			}
		}
		platform = "mac"
	case "linux":
		// 尝试多个环境变量
		for _, env := range []string{"LANG", "LANGUAGE", "LC_ALL"} {
			locale = os.Getenv(env)
			if locale != "" {
				break
			}
		}
		platform = "linux"
	}

	//
	lsf := LocalSYSInfoFunc{}
	localeInfo := lsf.parseLocale(locale)

	//
	return map[string]string{
		"homeDir":  homeDir,                               //
		"cacheDir": lsf.CachePath(),                       // 缓存存放目录
		"dataDir":  lsf.DataPath(),                        // 数据存放目录
		"numCPU":   strconv.FormatInt(int64(numCPUs), 10), // 逻辑CPU数
		"username": username,                              // 计算机用户名
		"platform": platform,                              // 平台
		"language": localeInfo.Language,                   // 语言
		"country":  localeInfo.Country,                    // 国别
		"encoding": localeInfo.Encoding,                   // 计算机编码方式
	}
}

// =========================================================

// SystemLocale 系统语言信息
type SystemLocale struct {
	Language string
	Country  string
	Encoding string
}

type LocalSYSInfoFunc struct {
}

func (lsf *LocalSYSInfoFunc) parseLocale(locale string) SystemLocale {
	// 示例: "zh_CN.UTF-8" -> Language: "zh", Country: "CN", Encoding: "UTF-8"
	result := SystemLocale{
		Language: "en",
		Country:  "US",
		Encoding: "UTF-8",
	}

	if locale == "" {
		return result
	}

	// 分割编码部分
	parts := strings.Split(locale, ".")
	if len(parts) > 1 {
		result.Encoding = parts[1]
	}

	// 分割语言和国家
	langParts := strings.Split(parts[0], "_")
	if len(langParts) > 0 {
		result.Language = langParts[0]
	}
	if len(langParts) > 1 {
		result.Country = langParts[1]
	}

	return result
}

// CachePath 当前平台存储程序缓存的路径，结尾无/
func (lsf *LocalSYSInfoFunc) CachePath() string {
	cachePath, err := os.UserCacheDir()
	if err != nil {
		return ""
	} else {
		return cachePath
	}
}

// IsMac 是mac平台
func (lsf *LocalSYSInfoFunc) IsMac() bool {
	sys := runtime.GOOS
	return sys == "darwin"
}

// IsWin 是win平台
func (lsf *LocalSYSInfoFunc) IsWin() bool {
	sys := runtime.GOOS
	return sys == "windows"
}

// IsLinux 是linux平台
func (lsf *LocalSYSInfoFunc) IsLinux() bool {
	sys := runtime.GOOS
	return sys == "linux"
}

// DataPath 当前平台存储程序持久化数据的路径，结尾无/
func (lsf *LocalSYSInfoFunc) DataPath() string {
	dataPath, err := os.UserHomeDir()
	if lsf.IsMac() {
		dataPath = dataPath + "/Library/Application Support"
	} else if lsf.IsWin() {
		dataPath = dataPath + "/AppData/Local"
	} else if lsf.IsLinux() {
		dataPath = dataPath + "/.local/share"
	} else {
		dataPath = dataPath + "/.wgo_data"
	}
	if err != nil {
		return ""
	} else {
		return dataPath
	}
}
