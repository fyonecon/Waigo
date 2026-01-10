package kits

import (
	"crypto/md5"
	"datathink.top/Waigo/internal"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

type LocalDB struct{}

const CodeKey string = "fyWaiG05" // 8位
const CodeSalt string = "2025"

// LocalDBSetData 新增或更新
func (ld *LocalDB) LocalDBSetData(dataKey string, dataValue string, dataTimeoutS int64) string {
	ldf := LocalDBFunc{}
	var LocalPath string = ldf.DataPath() + "/" + ldf.InterfaceToString(internal.GetConfigMap("sys", "dataPath"))

	secret := Secret{}
	filename := ldf.InterfaceToString(internal.GetConfigMap("app", "appClass")) + ldf.MD5(dataKey+CodeSalt) + ".lcl"
	_theFilepath := LocalPath + "/local_database"
	theFilepath := _theFilepath + "/" + filename
	//
	theTimer := ldf.GetTimeS() + dataTimeoutS // 截止日期
	//
	theValue := ldf.URLEncode(dataKey) + "\n" + ldf.IntToString(theTimer) + "\n" + secret.StringEncode(dataValue, CodeKey) // 写入3行内容
	// 无文件夹就创建
	if !ldf.IsDir(_theFilepath) {
		_, err := ldf.MakeDir(_theFilepath)
		if err != nil {
			return ""
		}
	}
	// 写入信息
	err := os.WriteFile(theFilepath, []byte(theValue), 0644)
	if err != nil {
		return ""
	}
	return theValue
}

// LocalDBGetData 获取
// 1有文件，-1无文件
func (ld *LocalDB) LocalDBGetData(dataKey string) (string, int64) {
	ldf := LocalDBFunc{}
	var LocalPath string = ldf.DataPath() + "/" + ldf.InterfaceToString(internal.GetConfigMap("sys", "dataPath"))

	secret := Secret{}
	filename := ldf.InterfaceToString(internal.GetConfigMap("app", "appClass")) + ldf.MD5(dataKey+CodeSalt) + ".lcl"
	_theFilepath := LocalPath + "/local_database"
	theFilepath := _theFilepath + "/" + filename
	//
	nowTime := ldf.GetTimeS()
	//
	if ldf.IsFile(theFilepath) {
		// 读取整个文件
		content, err := os.ReadFile(theFilepath)
		if err != nil {
			return "", -1
		} else {
			// 按换行符分割
			lines := strings.Split(string(content), "\n")
			if len(lines) == 3 {
				//_key := common.URLDecode(lines[0])
				_timeoutS := ldf.StringToInt(lines[1])
				_value := lines[2]
				if nowTime <= _timeoutS { // 有效
					return secret.StringDecode(_value, CodeKey), 1
				} else { // 过期
					return "", ld.LocalDBDelData(dataKey) // 删除文件
				}
			} else {
				return "", -1
			}
		}
	} else {
		return "", -1
	}
}

// LocalDBDelData 删除该DB文件
// 1有文件，-1无文件
func (ld *LocalDB) LocalDBDelData(dataKey string) int64 {
	ldf := LocalDBFunc{}
	var LocalPath string = ldf.DataPath() + "/" + ldf.InterfaceToString(internal.GetConfigMap("sys", "dataPath"))

	filename := ldf.InterfaceToString(internal.GetConfigMap("app", "appClass")) + ldf.MD5(dataKey+CodeSalt) + ".lcl"
	_theFilepath := LocalPath + "/local_database"
	theFilepath := _theFilepath + "/" + filename
	if ldf.IsFile(theFilepath) {
		err := ldf.DelFile(theFilepath)
		if err != nil {
			return -1
		} else {
			return 1
		}
	} else {
		return -1
	}
}

//==================================================

type LocalDBFunc struct{}

// MD5 生成md5
func (ldf *LocalDBFunc) MD5(_string string) string {
	md := md5.New()
	_, err := io.WriteString(md, _string)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md.Sum(nil))
}

// InterfaceToString interface{}，类似ValueInterfaceToString
func (ldf *LocalDBFunc) InterfaceToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

// ConvertedPath win下\转/，win兼容mac的path，统一转成mac path
func (ldf *LocalDBFunc) ConvertedPath(path string) string {
	path = ldf.URLEncode(path)
	path = strings.ReplaceAll(path, "%5C", "%2F")
	return ldf.URLDecode(path)
}

// IsDir 是目录
func (ldf *LocalDBFunc) IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 是文件
func (ldf *LocalDBFunc) IsFile(path string) bool {
	return !ldf.IsDir(path)
}

// DelFile 删除文件
func (ldf *LocalDBFunc) DelFile(oldFile string) error {
	oldFile = ldf.ConvertedPath(oldFile)
	return os.Remove(oldFile)
}

// StringToInt string转int
func (ldf *LocalDBFunc) StringToInt(_str string) int64 {
	_int, err := strconv.ParseInt(_str, 10, 64) // string转int
	if err != nil {
		_int = 0
	}
	return _int
}

// IntToString int转string
func (ldf *LocalDBFunc) IntToString(_int int64) string {
	_str := strconv.FormatInt(_int, 10)
	return _str
}

// GetTimeMS 获取毫秒时间戳，ms
func (ldf *LocalDBFunc) GetTimeMS() int64 {
	timer := time.Now()
	timeMS := int64(timer.UnixNano() / 1e6)
	return timeMS
}

// GetTimeS 获取秒时间戳
func (ldf *LocalDBFunc) GetTimeS() int64 {
	return time.Now().Unix()
}

// URLEncode URI加密，大写
func (ldf *LocalDBFunc) URLEncode(uri string) string {
	return strings.ReplaceAll(url.QueryEscape(uri), "+", "%20") // 已js的空格转换为标准
}

// URLDecode URI解密
func (ldf *LocalDBFunc) URLDecode(uri string) string {
	res, err := url.QueryUnescape(uri)
	if err != nil {
		return ""
	} else {
		return string(res)
	}
}

// DataPath 当前平台存储程序持久化数据的路径，结尾无/
func (ldf *LocalDBFunc) DataPath() string {
	dataPath, err := os.UserHomeDir()
	if ldf.IsMac() {
		dataPath = dataPath + "/Library/Application Support"
	} else if ldf.IsWin() {
		dataPath = dataPath + "/AppData/Local"
	} else if ldf.IsLinux() {
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

// IsMac 是mac平台
func (ldf *LocalDBFunc) IsMac() bool {
	sys := runtime.GOOS
	return sys == "darwin"
}

// IsWin 是win平台
func (ldf *LocalDBFunc) IsWin() bool {
	sys := runtime.GOOS
	return sys == "windows"
}

// IsLinux 是linux平台
func (ldf *LocalDBFunc) IsLinux() bool {
	sys := runtime.GOOS
	return sys == "linux"
}

// MakeDir 创建文件夹
func (ldf *LocalDBFunc) MakeDir(dir string) (string, error) {
	// 检查目录是否存在
	_, err := os.Stat(dir)
	if err != nil { // 如果目录不存在，则创建它
		if os.IsNotExist(err) {
			errDir := os.MkdirAll(dir, 0755)
			if errDir != nil {
				return "Error1", errors.New("创建文件夹出错：" + errDir.Error())
			}
			return dir, nil
		} else {
			return "Error2", nil
		}
	} else {
		//fmt.Println("目录已存在：", dir)
		return dir, nil
	}
}
