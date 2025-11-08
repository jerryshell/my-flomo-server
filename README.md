# My Flomo

My Flomo 是一个自托管的备忘录管理应用，提供与 Flomo 兼容的功能实现。该项目采用前后端分离架构，后端使用 Go 语言开发，前端使用 React 和 TypeScript 构建，支持数据完全自主控制。

## 功能特性

- **备忘录管理**: 支持备忘录的创建、编辑、删除和查看操作
- **数据导入导出**: 提供 Flomo 数据导入和 CSV 格式的数据交换功能
- **API 兼容**: 实现与 Flomo 兼容的 API 接口规范
- **邮件回顾**: 配置定时任务发送每日备忘录回顾邮件
- **插件支持**: 通过 API 插件令牌支持第三方应用集成
- **账户管理**: 用户注册、登录、密码修改和账户注销功能
- **数据安全**: 支持用户账户及相关数据的永久删除

## 技术架构

### 后端实现

- **编程语言**: Go
- **Web 框架**: Gin
- **数据存储**: SQLite 数据库，通过 GORM ORM 进行数据操作
- **身份认证**: JWT (JSON Web Token) 机制
- **定时任务**: Cron 表达式调度
- **邮件服务**: SMTP 协议支持

### 前端实现

- **编程语言**: TypeScript
- **前端框架**: React
- **状态管理**: Recoil
- **路由管理**: React Router DOM
- **构建工具**: Vite
- **HTTP 客户端**: Axios

## 项目结构

```
My Flomo/
├── api/                    # 后端服务
│   ├── config/            # 配置管理
│   ├── db/                # 数据库连接
│   ├── form/              # 表单结构定义
│   ├── handler/           # HTTP 请求处理器
│   ├── middleware/        # 中间件
│   ├── model/             # 数据模型
│   ├── result/            # 响应结果结构
│   ├── route/             # 路由定义
│   ├── service/           # 业务逻辑层
│   ├── store/             # 数据存储层
│   └── util/              # 工具函数
└── web/                   # 前端应用
    ├── src/
    │   ├── api/           # API 接口定义
    │   ├── atoms/         # Recoil 状态原子
    │   ├── components/    # React 组件
    │   ├── interfaces/    # TypeScript 接口定义
    │   └── pages/         # 页面组件
    └── public/            # 静态资源
```

## 安装与部署

### 后端服务部署

1. 进入后端目录

```bash
cd api
```

2. 安装项目依赖

```bash
go mod download
```

3. 配置环境变量（可选配置）

```bash
export HOST=0.0.0.0        # 服务监听地址，默认 localhost
export PORT=8060           # 服务监听端口，默认 8060
export DSN=My Flomo.db     # SQLite 数据库文件路径
export JWT_KEY=YOUR_JWT_KEY # JWT 签名密钥
export CRON_SPEC="0 20 * * *" # 定时任务表达式，默认每天 20:00
export SMTP_HOST=smtp.example.com # SMTP 服务器地址
export SMTP_PORT=587       # SMTP 服务端口
export SMTP_USERNAME=your_email@example.com # SMTP 认证用户名
export SMTP_PASSWORD=your_password # SMTP 认证密码
export SMTP_SUBJECT="My Flomo 每日回顾" # 邮件主题
```

4. 启动后端服务

```bash
go run main.go
```

### 前端应用部署

1. 进入前端目录

```bash
cd web
```

2. 安装项目依赖

```bash
npm install
```

3. 配置环境变量
   编辑 `.env` 文件，配置后端 API 地址：

```
VITE_API_BASE_URL=http://localhost:8060
```

4. 开发环境运行

```bash
npm run dev
```

5. 构建生产版本

```bash
npm run build
```

6. 预览构建结果

```bash
npm run preview
```

## API 接口文档

### 认证相关接口

- `POST /auth/login` - 用户登录或注册
- `POST /auth/verifyToken/token/:token` - 验证令牌有效性

### 备忘录管理接口

- `GET /memo/list` - 获取备忘录列表（需认证）
- `POST /memo/create` - 创建备忘录（需认证）
- `POST /memo/update` - 更新备忘录（需认证）
- `POST /memo/delete/id/:id` - 删除备忘录（需认证）
- `GET /memo/dailyReview` - 触发每日回顾邮件

### 插件功能接口

- `GET /plugin/getToken` - 获取插件令牌（需认证）
- `POST /plugin/createToken` - 创建插件令牌（需认证）
- `POST /plugin/createMemo/:pluginToken` - 通过插件令牌创建备忘录
- `GET /plugin/randomMemo/:pluginToken` - 通过插件令牌获取随机备忘录

### 用户管理接口

- `POST /user/updatePassword` - 更新用户密码（需认证）

### 数据管理接口

- `POST /upload` - 文件上传（需认证）
- `POST /deleteMyAccount` - 删除用户账户（需认证）
- `GET /csvExport/token/:token` - 导出 CSV 格式数据
- `POST /csvImport` - 导入 CSV 格式数据（需认证）

## 开源协议

本项目采用 GNU Affero General Public License v3.0 开源协议。详细条款请参阅 [LICENSE](LICENSE) 文件。

## 贡献指南

欢迎提交 Issue 报告和 Pull Request 贡献代码。在提交代码前，请确保：

1. 代码符合项目编码规范要求
2. 添加必要的测试用例覆盖
3. 更新相关文档内容

## 联系方式

如有问题或建议，请通过 GitHub Issues 页面提交反馈。
