
#### 项目框架结构
```html
├── /cmd                            # 项目入口 (main.go)
│   └── main.go                     # 整个应用的入口，在这里启动应用 
├── /configs                        # 存放应用的配置文件和配置加载逻辑
│   └── config.yaml                 # 应用的配置文件，通常包含数据库连接信息、服务器设置等
├── /docs                           # 存放应用的文档，如 API 文档、用户手册等
├── /internal                       # 私有代码 (禁止外部导入)
│   ├── /api                        # 路由定义
│   │   ├── /v1
│   │   │   └── routes.go
│   ├── /app                        # 核心应用逻辑，初始化、启动等
│   │   ├── /config
│   │   │   ├── config.go
│   │   │   └── ...
│   │   ├── /initializer            # 初始化，如db、log、redis等
│   │   │   ├── db_init.go
│   │   │   └── ...
│   │   ├── loader.go               # 加载工具、提供启动服务
│   │   └── ...
│   ├── /controllers                # 包含控制器逻辑，处理请求并返回响应
│   │   ├── user_controller.go
│   │   └── ...
│   ├── /middlewares                # 存放中间件代码，用于在请求处理流程中的特定阶段执行代码
│   │   ├── error_mid.go
│   │   └── ...
│   ├── /models                     # 定义应用的数据模型，通常与数据库表结构对应
│   │   ├── user.go
│   │   └── ...  
│   ├── /repositories               # 实现数据访问逻辑，与数据库进行交互
│   │   ├── user_repository.go
│   │   └── ...  
│   ├── /services                   # 实现业务逻辑，调用 repositories 中的方法来处理业务需求
│   │   ├── user_service.go
│   │   └── ...
│   └── /utils                      # 包含通用的工具函数，这些函数可以被多个包所共享
├── /logs                           # 日志
├── /migrations                     # 管理数据库变更
├── /pkg                            # 存放第三方库，如第三方中间件、工具库等
├── /scripts                        # 存放各种脚本，如项目部署脚本、测试脚本等
├── /tests                          # 存放测试代码，包括单元测试、集成测试等
├── .env                            # 环境变量配置
├── go.mod                          # Go 模块管理
├── go.sum                          
├── Dockerfile                      # 容器化配置
└── README.md                       # 项目说明
```