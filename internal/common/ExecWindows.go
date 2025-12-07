//go:build windows

package common

import (
	"os/exec"
	"syscall" // 1/2 Win真机代码片段
)

// RunMacShell 运行命令行
// 举例：out, err := RunMacShell("defaults write com.apple.finder AppleShowAllFiles -boolean true ; killall Finder")
func RunMacShell(shell string) (string, error) {
	cmd := exec.Command("bash", "-c", shell) // mac or linux
	out, err := cmd.CombinedOutput()
	back := ""
	if err != nil {
		back = "Error"
	}
	back = string(out)
	return back, err
}

// RunWinShell 运行命令行
func RunWinShell(shell string) (string, error) {
	cmd := exec.Command("cmd", "/C", shell) // windows

	// 2/2 Win真机代码片段
	// 这个特性会出现Goland与Mac系统不匹配的差异，导致build编译时不能正确识别该语法。解决办法：在win真实环境中编译即可。
	cmd.SysProcAttr = &syscall.SysProcAttr{
		HideWindow:    true,
		CreationFlags: 0x08000000,
	}

	out, err := cmd.CombinedOutput()
	back := ""
	if err != nil {
		back = "Error"
	}
	back = string(out)
	return back, err
}
