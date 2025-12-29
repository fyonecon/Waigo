package common

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
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
	// 1. 路由参数
	if value := ctx.Param(key); value != "" && value != "/" {
		return strings.TrimPrefix(value, "/")
	}

	// 2. 查询参数
	if value, exists := ctx.GetQuery(key); exists {
		return value
	}

	// 3. 处理 JSON 请求
	contentType := ctx.ContentType()
	if strings.Contains(contentType, "application/json") {
		return getJSONInput(ctx, key)
	}

	// 4. 表单数据
	if strings.Contains(contentType, "application/x-www-form-urlencoded") ||
		strings.Contains(contentType, "multipart/form-data") {
		if value, exists := ctx.GetPostForm(key); exists {
			return value
		}
	}

	// 5. 其他来源
	if value := ctx.PostForm(key); value != "" {
		return value
	}

	return ""
}

func getJSONInput(ctx *gin.Context, key string) string {
	// 使用缓存避免重复解析
	const cacheKey = "_parsed_json_body"

	// 检查缓存
	if cached, exists := ctx.Get(cacheKey); exists {
		if bodyMap, ok := cached.(map[string]interface{}); ok {
			return extractValue(bodyMap, key)
		}
	}

	// 读取并解析 JSON
	var requestBody map[string]interface{}

	// 方法1: 使用 ShouldBindJSON
	if err := ctx.ShouldBindJSON(&requestBody); err != nil {
		// 方法2: 使用 GetRawData 手动解析
		bodyBytes, err := ctx.GetRawData()
		if err != nil {
			return ""
		}

		// 恢复 body
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// 解析 JSON
		if err := json.Unmarshal(bodyBytes, &requestBody); err != nil {
			return ""
		}
	}

	// 缓存结果
	ctx.Set(cacheKey, requestBody)

	// 提取值
	return extractValue(requestBody, key)
}

func extractValue(data map[string]interface{}, key string) string {
	if data == nil {
		return ""
	}

	// 直接查找
	if value, exists := data[key]; exists {
		return formatValue(value)
	}

	// 嵌套查找 (user.name)
	if strings.Contains(key, ".") {
		parts := strings.Split(key, ".")
		current := interface{}(data)

		for _, part := range parts {
			if m, ok := current.(map[string]interface{}); ok {
				if val, exists := m[part]; exists {
					current = val
				} else {
					return ""
				}
			} else {
				return ""
			}
		}

		return formatValue(current)
	}

	return ""
}

func formatValue(v interface{}) string {
	if v == nil {
		return ""
	}

	// 对于字符串、数字、布尔值，直接格式化
	switch val := v.(type) {
	case string:
		return val
	case json.Number:
		return val.String()
	case bool:
		if val {
			return "true"
		}
		return "false"
	case int, int8, int16, int32, int64,
		uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", val)
	case float32, float64:
		// 使用 %v 自动选择合适格式
		return fmt.Sprintf("%v", val)
	default:
		// 对于数组和对象，返回 JSON 字符串
		if bytes, err := json.Marshal(val); err == nil {
			return string(bytes)
		}
		// 如果 JSON 转换失败，返回空字符串而不是 Go 的 map 字符串
		return ""
	}
}
