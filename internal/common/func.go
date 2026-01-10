//go:build darwin || windows || linux

package common

import (
	"crypto/md5"
	"datathink.top/Waigo/internal"
	kits "datathink.top/Waigo/internal/common/kits"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"os/user"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// StringToInt string转int
func StringToInt(_str string) int64 {
	_int, err := strconv.ParseInt(_str, 10, 64) // string转int
	if err != nil {
		_int = 0
	}
	return _int
}

// IntToString int转string
func IntToString(_int int64) string {
	_str := strconv.FormatInt(_int, 10)
	return _str
}

// GetTimeMS 获取毫秒时间戳，ms
func GetTimeMS() int64 {
	timer := time.Now()
	timeMS := int64(timer.UnixNano() / 1e6)
	return timeMS
}

// GetOS 获取系统平台
func GetOS() string {
	sys := runtime.GOOS
	back := ""
	if sys == "darwin" {
		back = "mac"
	} else if sys == "windows" {
		back = "win"
	} else if sys == "linux" {
		back = "linux"
	}
	return back
}

// IsMac 是mac平台
func IsMac() bool {
	sys := runtime.GOOS
	return sys == "darwin"
}

// IsWin 是win平台
func IsWin() bool {
	sys := runtime.GOOS
	return sys == "windows"
}

// IsLinux 是linux平台
func IsLinux() bool {
	sys := runtime.GOOS
	return sys == "linux"
}

// IsLanIPv4 属于白名单局域网IPv4
func IsLanIPv4(host string) bool {
	whiteIPv4 := []string{
		"127.0.0.1", "0.0.0.0", "192.168.", // C 类
		"172.16.", "172.17.", "172.18.", "172.19.", "172.20.", "172.21.", "172.22.", "172.23.", "172.24.", "172.25.", "172.26.", "172.27.", "172.28.", "172.29.", "172.30.", "172.31.", // B 类
		"10.",       // A 类
		"purehome.", // 绑定网址类
	}
	for _, ip := range whiteIPv4 {
		if strings.HasPrefix(host, ip) {
			return true
		}
	}
	return false
}

// Ping ping网址、ip，返回平局耗时ms
// url="www.google.com" count=执行次数
func Ping(url string, count int64) string {
	back := ""
	if count < 1 {
		count = 3
	} else if count > 1000 {
		count = 1000
	}
	_count := IntToString(count)

	startTime := GetTimeMS()
	//cmd := exec.Command("ping", url, "-c", _count) // -c <完成次数> 设置完成要求回应的次数。 W <timeout> 在等待 timeout 毫秒后开始执行。
	//err := cmd.Run()
	if IsMac() {
		ping := "ping " + url + " -c " + _count
		back, _ = kits.RunMacShell(ping)
	} else if IsWin() {
		ping := "ping " + url + " -n " + _count
		back, _ = kits.RunWinShell(ping)
	} else {
		ping := "ping " + url + " -c " + _count
		back, _ = kits.RunMacShell(ping)
	}
	endTime := GetTimeMS()
	spendTime := (endTime - startTime) / count / 10 // 平局耗时, ms
	back = IntToString(spendTime)
	if spendTime > 410 {
		back = "Long"
	}
	return back
}

// MakeDir 创建文件夹
func MakeDir(dir string) (string, error) {
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

// ReadFileDir 读取文件的绝对路径
func ReadFileDir(file string) string {
	relativePath := file
	absolutePath, _ := filepath.Abs(relativePath)
	return absolutePath
}

// DownloadFile 下载文件 。saveDir是文件夹，不能含有文件名
func DownloadFile(saveDir string, fileURL string, filename string) error {
	if len(filename) < 1 {
		filename = "null-filename"
	}
	var err error = nil
	res, err1 := http.Get(fileURL)
	if err1 != nil {
		err = err1
	}
	defer func(Body io.ReadCloser) {
		err5 := Body.Close()
		if err5 != nil {
			err = err5
		}
	}(res.Body)

	if res != nil {
		file, err2 := os.Create(saveDir + filename)
		if err2 != nil {
			fmt.Println("os.Create:", saveDir, filename)
			err = err2
		}
		defer func(file *os.File) {
			err6 := file.Close()
			if err6 != nil {
				err = err6
			}
		}(file)

		_, err3 := io.Copy(file, res.Body)
		if err3 != nil {
			err = err3
		}
	} else {
		fmt.Print("url获取内容为空：", fileURL)
	}

	return err
}

// GetFileFatherDirName 获取文件/文件夹所在的父级文件夹路径
func GetFileFatherDirName(file string) string {
	file = ConvertedPath(file)
	file = strings.TrimRight(file, "/") // 删除最后一位是/
	//
	sep := "/"
	parts := strings.Split(file, sep)
	newDir := ""
	for i := 0; i < len(parts)-1; i++ {
		if len(parts[0]) != 0 && i == 0 { // 第一位不为空时
			newDir = parts[0]
		} else {
			newDir = newDir + sep + parts[i]
		}
	}
	return newDir
}

// GetFileLastDirName 获取文件/文件夹最后一个文件夹或文件的名称
func GetFileLastDirName(file string) string {
	file = ConvertedPath(file)
	file = strings.TrimRight(file, "/") // 删除最后一位是/
	//
	sep := "/"
	parts := strings.Split(file, sep)
	if len(parts) == 0 {
		return file
	} else {
		return parts[len(parts)-1]
	}
}

// HasFileOrDir 是否已存在文件/文件夹
func HasFileOrDir(file string) (bool, int64) {
	f, err := os.Stat(file)
	if err != nil {
		if os.IsNotExist(err) {
			return false, 0
		}
	}
	size := f.Size()
	return true, size
}

// IsDir 是目录
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// IsFile 是文件
func IsFile(path string) bool {
	return !IsDir(path)
}

// MoveFile 移动文件或重命名文件
func MoveFile(oldFile string, newFile string) error {
	oldFile = ConvertedPath(oldFile)
	newFile = ConvertedPath(newFile)
	return os.Rename(oldFile, newFile)
}

// MoveDir 移动文件夹或重命名文件夹
func MoveDir(oldDir string, newDir string) error {
	oldDir = ConvertedPath(oldDir)
	newDir = ConvertedPath(newDir)
	return os.Rename(oldDir, newDir)
}

// RenameFile 移动文件或重命名文件
func RenameFile(oldFile string, newFile string) error {
	oldFile = ConvertedPath(oldFile)
	newFile = ConvertedPath(newFile)
	return os.Rename(oldFile, newFile)
}

// RenameDir 移动文件夹或重命名文件夹
func RenameDir(oldDir string, newDir string) error {
	oldDir = ConvertedPath(oldDir)
	newDir = ConvertedPath(newDir)
	return os.Rename(oldDir, newDir)
}

// DelFile 删除文件
func DelFile(oldFile string) error {
	oldFile = ConvertedPath(oldFile)
	return os.Remove(oldFile)
}

// DelDir 删除文件夹
func DelDir(oldDir string) error {
	oldDir = ConvertedPath(oldDir)
	return os.RemoveAll(oldDir)
}

// CopyFile 复制文件到新文件夹（支持大文件）
// oldFilePath 老路径+老文件名
// newPath 新路径，以/结尾
// newFilename 新文件名，只需要文件名
func CopyFile(oldFile string, newFile string) error {
	oldFile = ConvertedPath(oldFile)
	newFile = ConvertedPath(newFile)
	// 创建新文件夹
	if !IsDir(GetFileFatherDirName(newFile)) {
		_, err := MakeDir(newFile)
		if err != nil {
			return err
		}
	}
	// 打开源文件
	srcFile, err := os.Open(oldFile)
	if err != nil {
		return err
	}
	defer func() {
		err = srcFile.Close()
		if err != nil {
			fmt.Println("源文件关闭失败,原因是:", err)
		}
	}()

	// 创建目标文件,稍后会向这个目标文件写入拷贝内容
	distFile, err := os.Create(newFile)
	if err != nil {
		return err
	}
	defer func() {
		err = distFile.Close()
		if err != nil {
			fmt.Println("目标文件关闭失败,原因是:", err)
		}
	}()
	//定义指定长度的字节切片,每次最多读取指定长度
	var tmp = make([]byte, 1024*4)
	//循环读取并写入
	for {
		n, err := srcFile.Read(tmp)
		n, _ = distFile.Write(tmp[:n])
		if err != nil {
			if err == io.EOF { //读到了文件末尾,并且写入完毕,任务完成返回(关闭文件的操作由defer来完成)
				return nil
			} else {
				fmt.Println("拷贝过程中发生错误,错误原因为:", err)
			}
		}
	}
}

// CopyDir 复制文件夹
func CopyDir(oldDirPath string, newDirPath string) error {
	oldDirPath = ConvertedPath(oldDirPath)
	newDirPath = ConvertedPath(newDirPath)
	// 检查目录是否正确
	if IsDir(oldDirPath) {
		if !IsDir(newDirPath) { // 无新路径就直接创建
			_, err := MakeDir(newDirPath)
			if err != nil {
				return err
			}
		}
	} else {
		return errors.New("旧路径不存在，请重新选择要复制的路径")
	}
	//
	if strings.TrimSpace(oldDirPath) == strings.TrimSpace(newDirPath) {
		return errors.New("源路径与目标路径不能相同！")
	}
	// 遍历路径
	err := filepath.Walk(oldDirPath, func(path string, info os.FileInfo, err error) error {
		if info == nil {
			return err
		}
		//复制目录是将源目录中的子目录复制到目标路径中，不包含源目录本身
		if path == oldDirPath {
			return nil
		}
		//生成新路径
		newPath := strings.Replace(path, oldDirPath, newDirPath, -1)
		if !info.IsDir() { // file
			err := CopyFile(path, newPath)
			if err != nil {
				return err
			}
		} else { // dir
			if !IsDir(newPath) {
				_, err := MakeDir(newDirPath)
				if err != nil {
					return err
				}
			}
		}
		return nil
	})
	return err
}

// FormatTimeToDate 将time格式化成日期
func FormatTimeToDate(_format string, timer time.Time) (date string) {
	if len(_format) == 0 {
		_format = "YmdHisMS"
	}
	date = _format
	if timer.IsZero() {
		timer = time.Now()
	}

	var year int64 = int64(timer.Year())
	var month int64 = int64(timer.Month())
	var day int64 = int64(timer.Day())
	var hour int64 = int64(timer.Hour())
	var minute int64 = int64(timer.Minute())
	var second int64 = int64(timer.Second())
	var week int64 = int64(timer.Weekday())
	var ms int64 = int64(timer.UnixNano() / 1e6)
	var ns int64 = int64(timer.UnixNano() / 1e9)
	msTmp := IntToString(int64(math.Floor(float64(ms / 1000))))
	nsTmp := IntToString(int64(math.Floor(float64(ns / 1000000))))

	var _year string
	var _month string
	var _day string
	var _hour string
	var _minute string
	var _second string
	var _week string // 英文星期
	var _Week string // 中文星期
	var _ms string   // 毫秒
	var _ns string   // 纳秒

	_year = IntToString(year)
	if month < 10 {
		_month = IntToString(month)
		_month = "0" + _month
	} else {
		_month = IntToString(month)
	}
	if day < 10 {
		_day = IntToString(day)
		_day = "0" + _day
	} else {
		_day = IntToString(day)
	}
	if hour < 10 {
		_hour = IntToString(hour)
		_hour = "0" + _hour
	} else {
		_hour = IntToString(hour)
	}
	if minute < 10 {
		_minute = IntToString(minute)
		_minute = "0" + _minute
	} else {
		_minute = IntToString(minute)
	}
	if second < 10 {
		_second = IntToString(second)
		_second = "0" + _second
	} else {
		_second = IntToString(second)
	}
	_week = IntToString(week)
	WeekZh := [...]string{"日", "一", "二", "三", "四", "五", "六"} // 默认从"日"开始
	_Week = WeekZh[week]
	_ms = strings.Replace(IntToString(ms), msTmp, "", -1)
	_ns = strings.Replace(IntToString(ns), nsTmp, "", -1)

	// 替换关键词
	date = strings.Replace(date, "MS", _ms, -1)
	date = strings.Replace(date, "NS", _ns, -1)
	date = strings.Replace(date, "Y", _year, -1)
	date = strings.Replace(date, "m", _month, -1)
	date = strings.Replace(date, "d", _day, -1)
	date = strings.Replace(date, "H", _hour, -1)
	date = strings.Replace(date, "i", _minute, -1)
	date = strings.Replace(date, "s", _second, -1)
	date = strings.Replace(date, "W", _Week, -1)
	date = strings.Replace(date, "w", _week, -1)

	return
}

// GetTimeDate 获取日期时间戳，s
// Y年m月d号 H:i:s.MS.NS 星期W
func GetTimeDate(_format string) string {
	if len(_format) == 0 {
		_format = "YmdHisMS"
	}
	return FormatTimeToDate(_format, time.Now())
}

// StringHasString 判断字符串中是否包含某个字符串
// -1代表不包含，其他代表第一次出现的索引位置
// 请用string.contains代替
func StringHasString(bigString string, minString string) int64 {
	index := strings.Index(bigString, minString)
	return int64(index)
}

// ArrayHasString 数组中含有某字符串
func ArrayHasString(array []string, value string) bool {
	for _, e := range array {
		if e == value {
			return true
		}
	}
	return false
}

// ArrayInString array里面的值在字符串里面(包含)
func ArrayInString(array []string, bigStr string) int64 {
	for _, e := range array {
		index := strings.Index(bigStr, e)
		if index != -1 {
			return int64(index)
		}
	}
	return -1
}

// ArrayHasInt 数组中含有某数值
func ArrayHasInt(array []int64, value int64) bool {
	for _, e := range array {
		if e == value {
			return true
		}
	}
	return false
}

// MapHasKey map[string]interface{}中是否含有某个键
func MapHasKey(_map map[string]interface{}, key string) bool {
	if len(InterfaceToString(_map[key])) > 0 {
		return true
	} else {
		return false
	}
}

// GetUrlParam 获取url中的参数（非解码）
// 仅支持restful形式的参数
func GetUrlParam(_url string, _key string) (value string) {
	has := StringHasString(_url, "?")
	if has == -1 { // 不是url正常形式的话，就自动拼装完整，否则Get方法解析不了
		_url = "url-get?" + _url
	}
	u, err := url.Parse(_url)
	values := u.Query()
	if err != nil {
		value = ""
	} else {
		value = values.Get(_key)
	}
	return
}

// InterfaceToJsonString interface{}，类似ValueInterfaceToJsonString
func InterfaceToJsonString(value interface{}) string {
	bytes, _ := json.Marshal(value)
	return string(bytes)
}

// InterfaceToString interface{}，类似ValueInterfaceToString
func InterfaceToString(value interface{}) string {
	return fmt.Sprintf("%v", value)
}

// InterfaceToInt interface{}，类似ValueInterfaceToInt
func InterfaceToInt(value interface{}) int64 {
	return StringToInt(fmt.Sprintf("%v", value))
}

// StringStrip 字符串只留下数字、字母、下划线
func StringStrip(str string) string {
	if str == "" {
		return ""
	}
	str = strings.TrimSpace(str)
	reg := regexp.MustCompile(`[\W|_]{1,}`)
	return reg.ReplaceAllString(str, "")
}

// DateToTimeS 秒日期时间戳转时间戳，s
func DateToTimeS(_date string, format string) int64 {
	var date string
	if len(_date) == 0 { //给一个默认值
		date = GetTimeDate("YmdHis")
	} else {
		date = _date
	}

	var layout string
	if format == "YmdHis" || format == "" {
		layout = "20060102150405" // 转化所需内定模板
	} else if format == "Y-m-d H:i:s" {
		layout = "2006-01-02 15:04:05"
	} else if format == "Y年m月d日 H:i:s" {
		layout = "2006年01月02日 15:04:05"
	} else {
		layout = "20060102150405"
	}

	//日期转化为时间戳
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(layout, date, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64

	return timestamp
}

// TimeSToDate 秒时间戳转秒日期，ms
func TimeSToDate(_timeS int64, format string) string {
	var timeS int64
	if _timeS == 0 { //给一个默认值
		timeS = GetTimeS()
	} else {
		timeS = _timeS
	}

	var layout string
	if format == "YmdHis" || format == "" {
		layout = "20060102150405" // 转化所需内定模板
	} else if format == "Y-m-d H:i:s" {
		layout = "2006-01-02 15:04:05"
	} else if format == "Y年m月d日 H:i:s" {
		layout = "2006年01月02日 15:04:05"
	} else {
		layout = "20060102150405"
	}

	date := time.Unix(timeS, 0).Format(layout)
	return date
}

// DateToDate 将日期时间戳YmdHis转成日期时间戳Y-m-d H:i:s
func DateToDate(_date string) string {
	var date string
	if len(_date) == 0 {
		date = GetTimeDate("YmdHis")
	} else {
		date = _date
	}
	timeS := DateToTimeS(date, "")
	return TimeSToDate(timeS, "Y-m-d H:i:s")
}

// GetTimeS 获取秒时间戳
func GetTimeS() int64 {
	return time.Now().Unix()
}

// SlotTimeMSToYearDate 毫秒时间段转年月日
// timerMS 时间段，format格式YmdHisMS。SlotTimeMSToYearDate(back, "系统已经运行Y年m月d天H小时i分钟s秒")
func SlotTimeMSToYearDate(timerMS int64, format string) string {
	var year int64 = 0
	var month int64 = 0
	var day int64 = 0
	var hour int64 = 0
	var minutes int64 = 0
	var seconds int64 = 0
	var ms int64 = 0

	var _year int64 = 12 * 30 * 24 * 60 * 60 * 1000 // 按一年360天
	var _month int64 = 30 * 24 * 60 * 60 * 1000
	var _day int64 = 24 * 60 * 60 * 1000
	var _hour int64 = 60 * 60 * 1000
	var _minutes int64 = 60 * 1000
	var _seconds int64 = 1000

	year = (timerMS - 0) / _year
	month = (timerMS - year*_year) / _month
	day = (timerMS - year*_year - month*_month) / _day
	hour = (timerMS - year*_year - month*_month - day*_day) / _hour
	minutes = (timerMS - year*_year - month*_month - day*_day - hour*_hour) / _minutes
	seconds = (timerMS - year*_year - month*_month - day*_day - hour*_hour - minutes*_minutes) / _seconds
	ms = timerMS - year*_year - month*_month - day*_day - hour*_hour - minutes*_minutes - seconds*_seconds

	var __year string = IntToString(year)
	var __month string = IntToString(month)
	var __day string = IntToString(day)
	var __hour string = IntToString(hour)
	var __minutes string = IntToString(minutes)
	var __seconds string = IntToString(seconds)
	var __ms string = IntToString(ms)

	format = strings.ReplaceAll(format, "Y", __year)
	format = strings.ReplaceAll(format, "m", __month)
	format = strings.ReplaceAll(format, "d", __day)
	format = strings.ReplaceAll(format, "H", __hour)
	format = strings.ReplaceAll(format, "i", __minutes)
	format = strings.ReplaceAll(format, "s", __seconds)
	format = strings.ReplaceAll(format, "MS", __ms)

	return format
}

// SlotTimeMSToDayDate 毫秒时间段转年月日
// timerMS 时间段，format格式dHisMS。SlotTimeMSToDayDate(back, "系统已经运行d天H小时i分钟s秒")
func SlotTimeMSToDayDate(timerMS int64, format string) string {
	var day int64 = 0
	var hour int64 = 0
	var minutes int64 = 0
	var seconds int64 = 0
	var ms int64 = 0

	var _day int64 = 24 * 60 * 60 * 1000
	var _hour int64 = 60 * 60 * 1000
	var _minutes int64 = 60 * 1000
	var _seconds int64 = 1000

	day = (timerMS - 0) / _day
	hour = (timerMS - day*_day) / _hour
	minutes = (timerMS - day*_day - hour*_hour) / _minutes
	seconds = (timerMS - day*_day - hour*_hour - minutes*_minutes) / _seconds
	ms = timerMS - day*_day - hour*_hour - minutes*_minutes - seconds*_seconds

	var __day string = IntToString(day)
	var __hour string = IntToString(hour)
	var __minutes string = IntToString(minutes)
	var __seconds string = IntToString(seconds)
	var __ms string = IntToString(ms)

	format = strings.ReplaceAll(format, "d", __day)
	format = strings.ReplaceAll(format, "H", __hour)
	format = strings.ReplaceAll(format, "i", __minutes)
	format = strings.ReplaceAll(format, "s", __seconds)
	format = strings.ReplaceAll(format, "MS", __ms)

	return format
}

// MapToJsonString map转jsonString
func MapToJsonString(param map[string]interface{}) string {
	dataType, _ := json.Marshal(param)
	dataString := string(dataType)
	return dataString
}

// JsonStringToMap jsonString转map，只能是 {} 格式，不能是map开头格式
func JsonStringToMap(jsonString string) map[string]interface{} {
	var tempMap map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &tempMap)
	if err != nil {
		fmt.Println("JsonStringToMap=", err, jsonString)
		return map[string]interface{}{}
	} else {
		return tempMap
	}
}

// InterfaceToMap interface{}转map[string]interface{}
func InterfaceToMap(inter interface{}) map[string]interface{} {
	if inter == nil {
		return nil
	} else {
		return inter.(map[string]interface{})
	}
}

// InterfaceToArrayString interface{}转[]string{}，前提是格式分布格式是[]string{}
func InterfaceToArrayString(inter interface{}) []string {
	if inter == nil {
		return nil
	} else {
		return inter.([]string)
	}
}

// StaticTwoNumber 固定输出2位数字长度
func StaticTwoNumber(str string) string {
	_str := StringToInt(str)
	if _str < 10 {
		return "0" + str
	} else {
		return str
	}
}

// ReplaceNumberString 数字替换成定长的数字符号
func ReplaceNumberString(numString string) string {
	numString = strings.ReplaceAll(numString, "0", "𝟬")
	numString = strings.ReplaceAll(numString, "1", "𝟭")
	numString = strings.ReplaceAll(numString, "2", "𝟮")
	numString = strings.ReplaceAll(numString, "3", "𝟯")
	numString = strings.ReplaceAll(numString, "4", "𝟰")
	numString = strings.ReplaceAll(numString, "5", "𝟱")
	numString = strings.ReplaceAll(numString, "6", "𝟲")
	numString = strings.ReplaceAll(numString, "7", "𝟳")
	numString = strings.ReplaceAll(numString, "8", "𝟴")
	numString = strings.ReplaceAll(numString, "9", "𝟵")
	return numString
}

// Base64Encode base64加密
func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// Base64Decode base64解密
func Base64Decode(str string) string {
	res, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return ""
	} else {
		return string(res)
	}
}

// URLEncode URI加密，大写
func URLEncode(uri string) string {
	return strings.ReplaceAll(url.QueryEscape(uri), "+", "%20") // 已js的空格转换为标准
}

// URLDecode URI解密
func URLDecode(uri string) string {
	res, err := url.QueryUnescape(uri)
	if err != nil {
		return ""
	} else {
		return string(res)
	}
}

// MD5 生成md5
func MD5(_string string) string {
	md := md5.New()
	_, err := io.WriteString(md, _string)
	if err != nil {
		return ""
	}
	return fmt.Sprintf("%x", md.Sum(nil))
}

// LocalIPv4 获取本地局域网IPv4地址
// 优先级：127.0.0.1 > 192.168 > 172. > 10. > (169.254、0.0.0.0)
func LocalIPv4() ([]string, error) {
	var ips []string
	var theIP127 []string
	var theIP192 []string
	var theIP172 []string
	var theIP10 []string
	var theIP169 []string
	ipData, err := net.InterfaceAddrs()
	if err != nil {
		return ips, err
	}
	for _, a := range ipData {
		if ipNet, ok := a.(*net.IPNet); ok && !ipNet.IP.IsLoopback() && ipNet.IP.To4() != nil {
			theIP := ipNet.IP.String()
			//ips = append(ips, theIP)
			// 增加IP的优先级
			if StringHasString(theIP, "127.0.0.1") == 0 {
				theIP127 = append(theIP127, theIP)
			}
			if StringHasString(theIP, "192.168.") == 0 {
				theIP192 = append(theIP192, theIP)
			}
			if StringHasString(theIP, "172.") == 0 {
				theIP172 = append(theIP172, theIP)
			}
			if StringHasString(theIP, "10.") == 0 {
				theIP10 = append(theIP10, theIP)
			}
			if StringHasString(theIP, "169.254.") == 0 {
				theIP169 = append(theIP169, theIP)
			}
		}
	}
	// 按IP的优先级排序
	if len(theIP127) > 0 {
		ips = append(ips, theIP127...)
	}
	if len(theIP192) > 0 {
		ips = append(ips, theIP192...)
	}
	if len(theIP172) > 0 {
		ips = append(ips, theIP172...)
	}
	if len(theIP10) > 0 {
		ips = append(ips, theIP10...)
	}
	if len(theIP169) > 0 {
		ips = append(ips, theIP169...)
	}
	if len(ips) == 0 {
		ips = append(ips, "127.0.0.1")
	}
	return ips, nil
}

// GetHostIPv4 获取host=local_ip:port
func GetHostIPv4() string {
	host := ""
	ip, _ := LocalIPv4()
	if len(ip) > 0 {
		host = ip[0] + ":" + InterfaceToString(internal.ConfigMap["webServerPort"])
	} else {
		host = "127.0.0.1:" + InterfaceToString(internal.ConfigMap["webServerPort"])
	}
	return host
}

// ComputerUser 获取电脑本机的用名
func ComputerUser() map[string]string {
	info, _ := user.Current()
	nickname := info.Name
	username := info.Username
	uid := info.Uid
	dir := info.HomeDir

	if len(username) > 0 {
		username = ConvertedPath(username)
		array := strings.Split(username, "/")
		if len(array) >= 2 { // win
			username = array[1]
		}
	}

	return map[string]string{
		"username": username, // win/mac 都有
		"nickname": nickname, // win可能没有
		"uid":      uid,
		"dir":      dir,
	}
}

// ConvertedPath win下\转/，win兼容mac的path，统一转成mac path
func ConvertedPath(path string) string {
	path = URLEncode(path)
	path = strings.ReplaceAll(path, "%5C", "%2F")
	path = URLDecode(path)
	path = strings.ReplaceAll(path, "//", "/")
	if len(path) >= 2 && path[len(path)-1] == '/' { // 删除最后一位是/，但不包括只有一个/
		path = path[:len(path)-1]
	}
	return path
}

// MacTheme 获取mac当前的主题色
//func MacTheme() string {
//	dm := exec.Command("osascript", "-e", `tell application "System Events"`, "-e", `tell appearance preferences`, "-e", `return dark mode`, "-e", `end tell`, "-e", `end tell`)
//	var out bytes.Buffer
//	dm.Stdout = &out
//
//	err := dm.Run()
//	if err != nil { // 默认dark
//		return "dark"
//	}
//
//	if strings.TrimSpace(out.String()) == "false" { // light
//		return "light"
//	} else if strings.TrimSpace(out.String()) == "true" { // dark
//		return "dark"
//	} else {
//		return "dark"
//	}
//}

// CycleEventStateXXi00s 返回周期时间是否已经达到，"is"，每小时的第i分钟运行一次。周期10s刷新。
// 周期数：每分钟、每小时。
// i为分钟时刻，举例：""每分钟运行一次、"00"每小时第00分钟运行一次、"01"每小时第01分钟运行一次、"30"每小时第30分钟运行一次。
// 注意：i值为空或两位字符串数字
func CycleEventStateXXi00s(i string) bool {
	nowTime := GetTimeDate("is")
	if i == "" {
		nowTime = GetTimeDate("s")
	}
	return nowTime == i+"00" || // 周期1s
		nowTime == i+"01" || // 周期2s
		nowTime == i+"02" || // 周期3s
		nowTime == i+"03" || // 周期4s
		nowTime == i+"04" || // 周期5s
		nowTime == i+"05" ||
		nowTime == i+"06" || // 周期6s
		nowTime == i+"07" ||
		nowTime == i+"08" ||
		nowTime == i+"09" || // 周期10s
		nowTime == i+"10" ||
		nowTime == i+"11" // 周期12s
	//nowTime == i+"12" ||
	//nowTime == i+"13" ||
	//nowTime == i+"14" // 周期15s
	//nowTime == i+"15" ||
	//nowTime == i+"16" ||
	//nowTime == i+"17" ||
	//nowTime == i+"18" ||
	//nowTime == i+"19" // 周期20s
}

// CycleEventState00s 返回周期时间是否已经达到，"s"，周期10s刷新。
// 周期数：每分钟
func CycleEventState00s() bool {
	return CycleEventStateXXi00s("")
}

// GoRand 获取范围内的随机数
func GoRand(min int64, max int64) int64 {
	rand.Seed(time.Now().UnixNano())
	return rand.Int63n(max-min+1) + min
}

// GoRandString 获取长度范围内的随机字母数字
func GoRandString(minLen int64, maxLen int64) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_-%"
	rand.Seed(time.Now().UnixNano() + int64(rand.Intn(1000)))
	// 随机确定长度
	length := GoRand(minLen, maxLen)
	// 生成字符串
	result := make([]byte, length)
	for i := int64(0); i < length; i++ {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// StringToFloat string转float64
func StringToFloat(str string) float64 {
	fl, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0.0
	} else {
		return fl
	}
}

// GetFileContentType 获取文件类型
func GetFileContentType(filename string) string {
	fileExt := path.Ext(filename) //后缀
	fileExt = strings.ToLower(fileExt)
	fileContentType := kits.FileContentTypeDict[fileExt]
	if len(fileContentType) == 0 {
		fileContentType = "application/octet-stream"
	}
	return fileContentType
}

// GetFilename 获取
func GetFilename(fullPath string) string {
	return filepath.Base(fullPath)
}

// CachePath 当前平台存储程序缓存的路径，结尾无/
func CachePath() string {
	cachePath, err := os.UserCacheDir()
	if err != nil {
		return ""
	} else {
		return cachePath
	}
}

// DataPath 当前平台存储程序持久化数据的路径，结尾无/
func DataPath() string {
	dataPath, err := os.UserHomeDir()
	if IsMac() {
		dataPath = dataPath + "/Library/Application Support"
	} else if IsWin() {
		dataPath = dataPath + "/AppData/Local"
	} else if IsLinux() {
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

// CreateDataDirLevel1 创建必要目录
func CreateDataDirLevel1(dirName string) (bool, string) {
	_localDataPath := DataPath() + "/" + InterfaceToString(internal.GetConfigMap("sys", "dataPath")) // 结尾无/
	fullPath := _localDataPath + "/" + dirName
	if !IsDir(fullPath) {
		_, err := MakeDir(fullPath)
		if err != nil {
			return false, fullPath
		}
	}
	return true, fullPath
}

// FormatFileSize 格式化文件大小成可读大小
func FormatFileSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	// 预定义单位
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"}
	// 计算应该使用哪个单位
	exp := int(math.Log(float64(bytes)) / math.Log(float64(unit)))
	if exp >= len(units) {
		exp = len(units) - 1
	}
	// 计算值
	value := float64(bytes) / math.Pow(float64(unit), float64(exp))
	// 格式化输出，保留2位小数
	format := "%.2f %s"
	return fmt.Sprintf(format, value, units[exp])
}
