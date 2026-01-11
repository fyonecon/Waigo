package common

// LangDict 翻译对照表
// 调用：common.GetTranslate("test", lang)
var LangDict = map[string]interface{}{
	"test": map[string]string{ // 示例
		"zh": "测试",
		"en": "Test",
		"jp": "",
		"fr": "",
		"de": "",
		"ru": "",
		"es": "",
		"vi": "",
	},
	"_null": map[string]string{ // 必须
		"zh": " -空- ",
		"en": " -Null- ",
	},
	//

}
