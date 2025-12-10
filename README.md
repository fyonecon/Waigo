# Waigo是用Golang写的“视图窗口+稳定服务”的桌面端多功能程序基建
```text
Waigo基于Wails3、Ginvel3等。

前端默认Svelte。

（方+——爬说明，仅在Git——-h>ub发布，20251205）

开源地址：https://github.com/fyonecon/Waigo 。

```
Python版基座请戳：https://github.com/fyonecon/Ginthon 。

Golang版基座请戳：https://github.com/fyonecon/Waigo 。

### 程序目标：
```
需要 go1.24+

需要 webview2或webkit基础环境

已适配macOS12+、Win10+。Linux环境尚未验证。

开发IDE Goland、VSCode。

```

### 拉取项目：
> git clone -b main https://github.com/fyonecon/Waigo.git Waigo-Main

### 官方文档：
（⚠️注意wails3的文档和实际代码区别较大，效果以实际代码为准。）
> wails：https://v3alpha.wails.io/quick-start/installation/
> 
> gin：https://gin-gonic.com/zh-cn/docs/

### 项目结构：
```text
Waigo-Main
├── build 配置Mac、Win、Linux环境的软件图标、名称、版本、打包信息等
├── frontend 视图UI，默认SvelteKit
│   ├── src 视图发开发文件
│   │   └── lib 页面、公共文件、公共函数。
│   │   └── routes 路由、layout。
│   ├── static 静态文件
├── ginassets 额外的web资源、视图。默认端口9850，外部浏览器可访问
│   ├── files 放文件
│   │   └── test.txt
│   ├── html 放HTML文件或文件
│   │   ├── favicon.ico
│   │   └── index.html
├── go.mod
├── go.sum
├── internal 框架逻辑
│   ├── app 自定义程序
│   │   ├── app_gin 服务：Gin
│   │   │   ├── gin_controller
│   │   │   │   ├── Files.go
│   │   │   │   └── init.go 空间命名的Struct
│   │   │   ├── init.go 空间命名的Struct
│   │   │   ├── Middlewares.go
│   │   │   ├── RouteAPI.go
│   │   │   ├── RouteFile.go
│   │   │   └── RouteHTML.go
│   │   ├── app_services 服务：定时器等
│   │   │   ├── init.go 空间命名的Struct
│   │   │   └── TimeInterval.go
│   │   ├── app_tray 服务：状态栏托盘
│   │   │   ├── init.go 空间命名的Struct
│   │   │   └── Tray.go
│   │   ├── app_window 服务：Wails主视图相关
│   │   │   ├── init.go 空间命名的Struct
│   │   │   ├── window_controller
│   │   │   │   ├── GoRunJS.go Go调用JS
│   │   │   │   ├── init.go 空间命名的Struct
│   │   │   │   └── JSCallGo.go JS调用GO
│   │   │   └── Window.go
│   ├── bootstrap 框架逻辑的实现（服务启动）
│   │   ├── app_gin.go
│   │   ├── app_tray.go
│   │   ├── init_check_sys.go 运行系统环境检测
│   │   └── init_window.go
│   ├── common 公用函数、封装的kit
│   │   ├── Func.go 公用函数
│   │   ├── kits 封装的各种kit
│   │   │   ├── _7Z.go
│   │   │   ├── CacheData.go 软件启动后的自定义参数数据，关闭软件后消失
│   │   │   ├── ExecDarwin.go 自动在Mac环境下编译
│   │   │   ├── ExecWindows.go 自动在Win环境下编译
│   │   │   ├── FileContentTypeDict.go 文件格式对照表
│   │   │   ├── ICON.go 二进制图标
│   │   │   ├── IconData.go 二进制图标
│   │   │   └── Secret.go 对称加密
│   │   └── Language.go
│   ├── ConfigMap.go 配置文件：读
│   └── ConfigSetter.go 配置文件：读、写
├── LICENSE
├── main.go
├── README.md
├── show.png
└── Taskfile.yml Wails官方的命令行任务

```

### 运行效果：
![运行效果](./show.png)

### Golang环境搭建:
```text
Go下载：https://golang.google.cn/dl/

Node下载：https://nodejs.org/en/download

Git下载：https://git-scm.com/install/mac

开发工具（IDE）：VSCode、Goland
```

开启mod模式:
> go env -w GO111MODULE=on

go get大陆地区代理：

> go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

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

### 打包成桌面可运行程序包：
```text
（正在测试）
```

### 2025-12-05