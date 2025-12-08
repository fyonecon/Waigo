//go:build darwin || linux

package kits

import "os/exec"

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

	out, err := cmd.CombinedOutput()
	back := ""
	if err != nil {
		back = "Error"
	}
	back = string(out)
	return back, err
}
