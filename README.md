# gomall

### IDL
echo.thrift 安装cwgo和thriftgo
`go install github.com/cloudwego/cwgo@latest`
`GO111MODULE=on go install github.com/cloudwego/thriftgo@latest`
cwgo zsh 自动补全命令行
```
mkdir autocomplete # You can choose any location you like
cwgo completion zsh > ./autocomplete/zsh_autocomplete
source ./autocomplete/zsh_autocomplete
```
cwgo根据idl生成demo_thrift服务的代码
```
cwgo server --type RPC --module github.com/oldmoonni/gomall/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift
```
修改一下biz/service/echo.go的代码(业务逻辑代码)
```
go mod tidy
go work use .
go run .
```
protobuf同理

最后创建一个Makefile文件 为了减少命令的重复输入
 

### 服务注册与服务发现
安装consul
```
go get github.com/kitex-contrib/registry-consul
```
然后在demo_proto的main.go中添加服务注册
启动docker容器的consul
启动demo_proto，注意检查conf.yaml中的端口号