# Waigo是用Golang写的桌面框架
```text
Waigo基于Wails3、Ginvel3等。

前端使用Svelte。

开发条件（🔥）：go1.25+、win10+、macOS12+ 。

（方+——爬说明，仅在Git——-h>ub发布，20251205）

Github：https://github.com/fyonecon/Waigo 。

```

Python版请戳：https://github.com/fyonecon/Ginthon 。
Python版视窗功能没有Golang版全，差异在于各自依赖的生态不同。

==============================

（wails3的文档和实际代码区别较大，正在写）

==============================

### 拉取项目
> git clone -b main https://github.com/fyonecon/Waigo.git Waigo-Main

### 打包成桌面可运行程序包
```text
（正在测试）
```

==============================

### 运行效果：
![运行效果](./show.png)

### Wails3新项目搭建:
0.安装wails3:
> go install github.com/wailsapp/wails/v3/cmd/wails3@latest
> 
> wails3 doctor

1.查看支持哪些前端框架：
> wails3 init -l

2.初始化一个新项目Waigo，前端框架使用sveltekit：
> wails3 init -n Waigo -t sveltekit

3.运行：
> wails3 dev

### Golang环境搭建:
开启mod模式:
> go env -w GO111MODULE=on

go get大陆地区代理： 

> go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

### 2025-12-05