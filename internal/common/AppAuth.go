package common

import (
	"datathink.top/Waigo/internal"
	"datathink.top/Waigo/internal/common/kits"
	"strings"
)

// MakeRandToken 生成Token
func MakeRandToken(salt string, timeoutS int64) string {
	appClass := InterfaceToString(internal.GetConfigMap("app", "appClass"))
	//appVersion := InterfaceToString(internal.GetConfigMap("app", "appVersion"))
	//
	if timeoutS < 10 {
		timeoutS = 10
	}
	_timeoutMS := timeoutS * 1000
	_nowTimeMS := GetTimeMS()
	_salt := MD5(salt + "2025WaiG0" + appClass)
	_key := "goWaI2025" // 8位key
	if len(_key) < 8 {
		_key = "goWaI2520"
	} else {
		_key = _key[:8]
	}
	secret := kits.Secret{}
	//
	token := appClass + "#@" + _salt + "#@" + IntToString(_nowTimeMS) + "#@" + IntToString(_timeoutMS) + "#@" + GoRandString(12, 18)
	//
	return URIEncode(secret.StringEncode(token, _key))
}

// CheckRandToken 检查Token
func CheckRandToken(salt string, checkString string) bool {
	appClass := InterfaceToString(internal.GetConfigMap("app", "appClass"))
	//appVersion := InterfaceToString(internal.GetConfigMap("app", "appVersion"))
	_nowTimeMS := GetTimeMS()
	_salt := MD5(salt + "2025WaiG0" + appClass)
	_key := "goWaI2025" // 8位key
	if len(_key) < 8 {
		_key = "goWaI2520"
	} else {
		_key = _key[:8]
	}
	secret := kits.Secret{}
	//
	token := secret.StringDecode(URIDecode(checkString), _key)
	tokenArray := strings.Split(token, "#@")
	if len(tokenArray) >= 5 {
		//fmt.Println("CheckRandToken=", token)
		__appClass := tokenArray[0]
		__salt := tokenArray[1]
		__timeoutMS := StringToInt(tokenArray[2]) + StringToInt(tokenArray[3])
		return appClass == __appClass && __salt == _salt && (_nowTimeMS <= __timeoutMS)
	} else {
		return false
	}
}
