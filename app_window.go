package main

import "context"

// RunAppStartService 软件启动时执行，页面刷新不执行
func RunAppStartService(wailsCTX context.Context) {

}

// RunPageLoadedService 页面Loaded以后执行，刷新页面后执行
func RunPageLoadedService(wailsCTX context.Context) {

}

// RunPageCloseService 在前端被销毁之后，应用程序终止之前
func RunPageCloseService(wailsCTX context.Context) {

}
