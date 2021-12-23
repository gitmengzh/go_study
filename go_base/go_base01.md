# Windows安装
# vscode环境配置
# go目录结构
# go语言依赖管理   http://c.biancheng.net/view/4774.html
    * set GO111MODULE=on 或者 set GO111MODULE=auto
    * go env -w GOPROXY=https://goproxy.cn,direct
    * Go语言的包与文件夹是一一对应的，它具有以下几点特性：
        一个目录下的同级文件属于同一个包。
        包名可以与其目录名不同。
        main 包是Go语言程序的入口包，一个Go语言程序必须有且仅有一个 main 包。如果一个程序没有 main 包，那么编译时将会出错，无法生成可执行文件。
# go编译
    * go build
    * go run