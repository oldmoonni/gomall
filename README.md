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
#### 服务端
安装consul
```
go get github.com/kitex-contrib/registry-consul
```
然后在demo_proto的main.go中添加服务注册
启动docker容器的consul
启动demo_proto，注意检查conf.yaml中的端口号 
#### 客户端
编写cmd/client/main.go 参考kitex/consul的官网文档

### 配置管理
包括文件配置、环境配置和配置中心
#### 环境配置
将需要配置的参数写入服务根目录下的.env文件中
然后利用godotenv包读取环境变量
```
# 安装
go get github.com/joho/godotenv
```
```
# 使用
package main

import (
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
  err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }

  s3Bucket := os.Getenv("S3_BUCKET")
  secretKey := os.Getenv("SECRET_KEY")

  // now do something with s3 or whatever
}
```

### 使用ORM进行数据操作
安装MySQL插件
GORM官方文档：
https://gorm.io/zh_CN/docs/models.html

在biz目录下创建model目录，存放数据库模型
在数据库启动代码中加入自动迁移
```
db.AutoMigrate(&User{})
```
启动demo_proto服务后，会自动在数据库中创建user表
想要改变表的名字的话，使用TableName函数
```
type User struct {
	gorm.Model
	Email    string `gorm:"uniqueIndex;type:varchar(128) not null"`
	Password string `gorm:"type:varchar(64) not null"`
}

func (User) TableName() string {
	return "user"
}
```