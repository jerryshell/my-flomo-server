![my-flomo-server](https://socialify.git.ci/jerryshell/my-flomo-server/image?description=1&forks=1&issues=1&language=1&name=1&owner=1&pattern=Brick%20Wall&pulls=1&stargazers=1&theme=Dark)

## 功能清单

- Flomo 数据导入
- Flomo API 兼容
- 邮件每日回顾
- 注销账号，永久抹除数据
- CSV 数据导入导出

## 体验 Demo

~~[https://my-flomo.d8s.fun](https://my-flomo.d8s.fun)~~

**服务器到期，体验 Demo 的后端已经关闭**

## 如何运行

### 本地运行

my-flomo-server 使用 SQLite 数据库，无需额外准备数据库

无需建表，my-flomo-server 第一次启动会自动创建 SQLite 数据库文件

```bash
# 1. 克隆项目
git clone https://github.com/jerryshell/my-flomo-server.git
cd my-flomo-server

# 2. 安装依赖
go mod download

# 3. 运行服务
go run main.go
```

服务默认运行在 8080 端口，你可以通过修改配置文件来更改端口和其他设置。

## 相关项目

- [Web 端](https://github.com/jerryshell/my-flomo-web)
- [服务端](https://github.com/jerryshell/my-flomo-server)

## 开源协议

[GNU Affero General Public License v3.0](https://choosealicense.com/licenses/agpl-3.0/)
