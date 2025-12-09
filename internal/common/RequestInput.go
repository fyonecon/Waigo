package common

import (
	"github.com/gin-gonic/gin"
	"strings"
)

// RequestFullURL 获取完整的请求域名
func RequestFullURL(ctx *gin.Context) string {
	//
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	} else {
		if ctx.GetHeader("X-Forwarded-Ssl") == "on" {
			scheme = "https"
		}
	}
	//
	host := ctx.Request.Host
	param := ctx.Request.RequestURI
	// 拼接成完整域名
	fullDomain := scheme + "://" + host + param
	//
	return fullDomain
}

// RequestInput 万能获取请求参数
func RequestInput(ctx *gin.Context, key string) string {
	// 1. 路由参数（最明确）
	if value := ctx.Param(key); value != "" {
		return value
	}
	// 2. 查询参数（所有方法都可能有）
	if value, exists := ctx.GetQuery(key); exists {
		return value
	}
	// 3. 根据Content-Type处理body
	contentType := ctx.ContentType()
	switch {
	case strings.Contains(contentType, "application/json"):
		// 使用缓存避免重复解析
		if value, found := getFromJSONWithCache(ctx, key); found {
			return value
		}
		//var requestBody map[string]interface{}
		//err := ctx.ShouldBindJSON(&requestBody)
		//if err != nil {
		//	return ""
		//} else {
		//	return InterfaceToString(requestBody[key])
		//}
	case strings.Contains(contentType, "application/x-www-form-urlencoded"), strings.Contains(contentType, "multipart/form-data"):
		if value, exists := ctx.GetPostForm(key); exists {
			return value
		}
	}
	//
	return ""
}

// JSON解析缓存键
const jsonBodyCacheKey = "_cached_json_body"

// getFromJSONWithCache 带缓存的JSON参数获取
// 性能优化
func getFromJSONWithCache(ctx *gin.Context, key string) (string, bool) {
	// 检查是否已经有缓存的解析结果
	if cached, exists := ctx.Get(jsonBodyCacheKey); exists {
		if cached == nil {
			// 缓存表示之前解析失败
			return "", false
		}

		if body, ok := cached.(map[string]interface{}); ok {
			if value, exists := body[key]; exists {
				return InterfaceToString(value), true
			}
		}
		return "", false
	}

	// 第一次解析
	var requestBody map[string]interface{}
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		// 解析失败，缓存nil避免重复尝试
		ctx.Set(jsonBodyCacheKey, nil)
		return "", false
	}

	// 缓存解析结果
	ctx.Set(jsonBodyCacheKey, requestBody)

	// 返回请求的值
	if value, exists := requestBody[key]; exists {
		return InterfaceToString(value), true
	}

	return "", false
}
