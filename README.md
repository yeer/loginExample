## 说明

1. MACOS 系统依赖下面的软件
   - docker (Docker Desktop)
   - Kubernetes (Docker Desktop)
     - 需要修改镜像地址
     - `vim ~/.docker/daemon.json`
     - 把下面内容添加到 json 中: `"registry-mirrors":["https://u99q7fs9.mirror.aliyuncs.com"]`
   - Make
   - helm
   - golang v1.16

## 运行方式

-

## 测试方式

#### curl 方式

1. 创建用户：
   - `curl -H "Content-Type: application/json" -X POST -d 'username=foo&password=xx22@4' "http://127.0.0.1:10010/v1/users/create"`
2. 用户登录
   - `curl -X POST -d 'username=foo&password=xx22@4' "http://127.0.0.1:10010/v1/login`

#### 单测方式

1. 简单运行单元测试
   `go test -v`

### 目录结构：

```
├── README.md
├── app
│   ├── controllers
│   ├── models
│   └── service
├── cmd
│   ├── root.go
│   └── server.go
├── conf
│   └── conf.yaml
├── go.mod
├── go.sum
├── main.go
├── server
│   └── server.go
└── util
    ├── cache
    ├── db
    ├── file
    ├── logger
    └── response
```
