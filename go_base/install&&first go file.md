# 安装
 1. windows安装
   下载安装包直接安装
 2. vscode编译环境设置
    2.1 安装go插件
      安装失败，设置代理
      go env -w GO111MODULE=on    打开go module
      go env -w GOPROXY=https://goproxy.c  配置七牛云代理
      go get -v golang.org/x/tools/cmd/goimports
      go get -v golang.org/x/tools/gopls
# 运行第一个go文件
 1. 命令行运行 go run xxx.go
 2. vscode 安装code Runner插件运行