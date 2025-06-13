gin-web for Blog Project

#### 项目框架结构
```html
├── /cmd                                # 项目入口 (main.go)
│   └── main.go                         # 整个应用的入口，在这里启动应用 
├── /configs                            # 存放应用的配置文件
│   └── config.yaml                     # 应用的配置文件，通常包含数据库连接信息、服务器设置等
├── /docs                               # 存放应用的文档，如 API 文档、用户手册等
├── /internal                           # 私有代码 (禁止外部导入)
│   ├── /api                            # 路由定义
│   │   ├── /v1
│   │   │   └── routes.go
│   ├── /app                            # 核心应用逻辑
    │   ├── /controllers                # 包含控制器逻辑，处理请求并返回响应
│   │   │   ├── user_controller.go
│   │   │   └── ...
│   │   ├── /models                     # 定义应用的数据模型，通常与数据库表结构对应
│   │   │   ├── user.go
│   │   │   └── ...
│   │   ├── /repositories               # 实现数据访问逻辑，与数据库进行交互
│   │   │   ├── user_repository.go
│   │   │   └── ...
│   │   ├── /services                   # 实现业务逻辑，调用 repositories 中的方法来处理业务需求
│   │   │   ├── user_service.go
│   │   │   └── ...
│   ├── /loader                         # 配置文件加载，初始化、启动等
│   │   ├── /config
│   │   │   ├── config.go               # 加载配置文件
│   │   │   └── ...
│   │   ├── /initializer                # 初始化，如db、log、redis等
│   │   │   ├── db_init.go
│   │   │   └── ...
│   │   ├── loader.go                   # 加载工具、提供启动服务
│   │   └── ...
│   ├── /middlewares                    # 存放中间件代码，用于在请求处理流程中的特定阶段执行代码
│   │   ├── error_mid.go
│   │   └── ...
│   └── /utils                          # 包含通用的工具函数，这些函数可以被多个包所共享
├── /logs                               # 日志
├── /migrations                         # 管理数据库变更
├── /pkg                                # 存放第三方库，如第三方中间件、工具库等
├── /scripts                            # 存放各种脚本，如项目部署脚本、测试脚本等
├── /tests                              # 存放测试代码，包括单元测试、集成测试等
├── .env                                # 环境变量配置
├── go.mod                              # Go 模块管理
├── go.sum                          
├── Dockerfile                          # 容器化配置
└── README.md                           # 项目说明
```

### 项目启动
#### 1. 运行环境 
Go版本: 1.24
数据库: MySQL 8 +

#### 2. 依赖安装步骤 
下载依赖
```bash
go mod tidy
```
#### 3. 启动项目
* 使用air实现热重载 (推荐)
```bash
go install github.com/cosmtrek/air@latest
air
```
* 或者直接运行
```bash
go run cmd/main.go
```

### 项目测试
#### 数据 migrations 目录下
1. 使用 [blog.ddl](migrations/blog.ddl) 创建数据库和表
2. 使用 [export.sql](migrations/export.sql) 插入数据
#### 测试用例及 API 列表 在 docs 目录下
1. openapi spec form: [gin-web.openapi.json](docs/gin-web.openapi.json)
2. postman format: [gin-web.postman.json](docs/gin-web.postman.json)

### 项目概述和业务需求
- [x]  1. 项目初始化
    - 创建一个新的 Go 项目，使用 go mod init 初始化项目依赖管理。
    - 安装必要的库，如 Gin 框架、GORM 以及数据库驱动（如 MySQL 或 SQLite）。
    - 数据库设计与模型定义
- [x]  2. 设计数据库表结构，至少包含以下几个表：
    - users 表：存储用户信息，包括 id 、 username 、 password 、 email 等字段。
    - posts 表：存储博客文章信息，包括 id 、 title 、 content 、 user_id （关联 users 表的 id ）、 created_at 、 updated_at 等字段。
    - comments 表：存储文章评论信息，包括 id 、 content 、 user_id （关联 users 表的 id ）、 post_id （关联 posts 表的 id ）、 created_at 等字段。,
    - 使用 GORM 定义对应的 Go 模型结构体。
- [x]  3.用户认证与授权
    - 实现用户注册和登录功能，用户注册时需要对密码进行加密存储，登录时验证用户输入的用户名和密码。
    - 使用 JWT（JSON Web Token）实现用户认证和授权，用户登录成功后返回一个 JWT，后续的需要认证的接口需要验证该 JWT 的有效性。
- [x]  4.文章管理功能
    - 实现文章的创建功能，只有已认证的用户才能创建文章，创建文章时需要提供文章的标题和内容。
    - 实现文章的读取功能，支持获取所有文章列表和单个文章的详细信息。
    - 实现文章的更新功能，只有文章的作者才能更新自己的文章。
    - 实现文章的删除功能，只有文章的作者才能删除自己的文章。
- [x]  5.评论功能
    - 实现评论的创建功能，已认证的用户可以对文章发表评论。
    - 实现评论的读取功能，支持获取某篇文章的所有评论列表。
- [x]  6.错误处理与日志记录
    - 对可能出现的错误进行统一处理，如数据库连接错误、用户认证失败、文章或评论不存在等，返回合适的 HTTP 状态码和错误信息。
    - 使用日志库记录系统的运行信息和错误信息，方便后续的调试和维护。





