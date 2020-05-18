# go-micro demo项目

### 目录结构

    ├─ Api                      接口目录
    ├─ Library                  公共扩展库
    |   ├─ cache                缓存相关
    |   ├─ jwt                  jwt相关
    |   ├─ log                  日志相关
    |   ├─ orm                  数据库相关
    |   ├─ queue                消息队列相关
    |   ├─ tool                 公共函数相关
    │   └─ ...                  更多虚拟机目录
    ├─ Models
    ├─ Proto
    |   ├─ pb                   pb.go文件目录
    |   └─ protos               protobuf文件目录
    ├─ Servers                  微服务目录
    |   ├─ User                 用户服务目录
    |   |   ├─ Handler          业务层
    |   |   ├─ Servers          服务层
    |   |   ├─ Migrate          数据库迁移
    |   |   ├─ Subscriber       消息订阅层
    │   |   └─ ...              更多微服务分层
    |   ├─ Order                订单服务目录
    |   ├─ Wallet               钱包服务目录
    │   └─ ...                  更多微服务目录
    └─ Shell                    Shell命令文件目录
