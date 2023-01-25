# dou-sheng

gin+grpc+gorm+etcd+mysql 实现抖声后台接口


# 项目主要依赖
- gin
- gorm
- etcd
- grpc
- jwt-go
- logrus
- viper
- protobuf

# 项目结构

## 1. api-gateway 网关部分

```
api-gateway/
├── cmd                   // 启动入口
├── config                // 配置文件
├── discovery             // etcd服务注册、keep-alive、获取服务信息等等
├── internal              // 业务逻辑（不对外暴露）
│   ├── handler           // 视图层
│   └── service           // 服务层
│       └──pb             // 放置生成的pb文件
├── logs                  // 放置打印日志模块
├── middleware            // 中间件
├── pkg                   // 各种包
│   ├── e                 // 统一错误状态码
│   ├── res               // 统一response接口返回
│   └── util              // 各种工具、JWT、Logger等等..
├── routes                // http路由模块
└── wrappers              // 各个服务之间的熔断降级
```

## 2. 模块


```
user/
├── cmd                   // 启动入口
├── config                // 配置文件
├── discovery             // etcd服务注册、keep-alive、获取服务信息等等
├── internal              // 业务逻辑（不对外暴露）
│   ├── handler           // 视图层
│   ├── cache             // 缓存模块
│   ├── repository        // 持久层
│   └── service           // 服务层
│       └──pb             // 放置生成的pb文件
├── logs                  // 放置打印日志模块
└── pkg                   // 各种包
    ├── e                 // 统一错误状态码
    ├── res               // 统一response接口返回
    └── util              // 各种工具、JWT、Logger等等..
```

# 项目完善
- 添加熔断机制
- ....其他想法

# 项目文件配置

各模块下的`config/config.yml`文件


```yaml
server:
# 模块
  domain: user
  # 模块名称
  version: 1.0
  # 模块版本
  grpcAddress: "127.0.0.1:10001"
  # grpc地址

datasource:
# mysql数据源
  driverName: mysql
  host: 127.0.0.1
  port: 3306
  database: basicInfo
  # 数据库名
  username: root
  password: root
  charset: utf8mb4

etcd:
# etcd 配置
  address: 127.0.0.1:2379

redis:
# redis 配置
  address: 127.0.0.1:6379
  password:
```



# 项目启动
保证etcd处于运行状态。
- 在各模块下进行

```go
go mod tidy
```

- 在各模块下的cmd目录

```go
go run main.go
```