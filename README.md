## BI-Activity

### 项目结构说明
- `dao` 数据库访问层, 封装了数据库操作
- `service` 业务层, 业务逻辑处理，调用数据库访问层
- `controller` 控制层，处理请求参数，返回数据，调用业务层
- `util` 工具包，常用方法
- `models` 数据表模型
- `configs` 配置文件及加载
- `router` 路由映射
- `cmd` 启动入口
- `response` 响应封装
    - `errors` 错误封装
- `middleware` 中间件

### 环境构建
1. MySQL 8.0版本
  - 数据库配置见 `configs/configs.yaml`
  - 运行`init.sql`创建数据库
  - 运行`cmd/flag/main.go`迁移数据表