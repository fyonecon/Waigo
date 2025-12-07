package main

type WailsServices struct{}

// JSCallGo js调用go
func (asv *WailsServices) JSCallGo(key string, dataDict map[string]interface{}) map[string]interface{} {
	return map[string]interface{}{
		"state": 1,
		"msg":   "调用成功",
		"content": map[string]interface{}{
			"key":       key,
			"data_dict": dataDict,
			"result":    "OKOKOK",
		},
	}
}
