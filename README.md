![my-flomo-server](https://socialify.git.ci/jerryshell/my-flomo-server/image?description=1&descriptionEditable=%E5%8F%AF%E4%BB%A5%E8%87%AA%E5%B7%B1%E6%90%AD%E5%BB%BA%E7%9A%84%E6%83%B3%E6%B3%95%E8%AE%B0%E5%BD%95%E5%8F%8A%E5%9B%9E%E9%A1%BE%E6%9C%8D%E5%8A%A1%EF%BC%8C%E5%BC%80%E6%BA%90%E3%80%81%E5%85%8D%E8%B4%B9%E3%80%81%E7%AE%80%E5%8D%95%E3%80%81%E4%B8%93%E6%B3%A8%E4%BA%8E%E6%A0%B8%E5%BF%83%E5%8A%9F%E8%83%BD&font=Raleway&forks=1&issues=1&language=1&owner=1&pattern=Brick%20Wall&pulls=1&stargazers=1&theme=Dark)

## 功能清单

* Flomo 数据导入
* Flomo API 兼容
* 邮件每日回顾
* Memo 增删改查
* 用户登录注册

## 体验 Demo

**⚠️ 注意：推荐每个用户单独搭建自己的服务，体验 Demo 不保证数据安全性，所以请勿在体验 Demo 中使用真实用户名密码注册！请勿在体验 Demo 中录入敏感数据！**

[https://my-flomo.pages.dev/](https://my-flomo.pages.dev/)

## 如何运行

### 1 下载可执行文件

https://github.com/jerryshell/my-flomo-server/releases

### 2 在可执行文件同目录下创建 config.json（可选）

#### 2.1 如果没有 config.json，则使用以下默认配置

```json
{
  "port": 8060,
  "dsn": "host=localhost user=my_flomo password=my_flomo dbname=my_flomo port=5432 sslmode=disable TimeZone=Asia/Shanghai",
  "jwtKey": "jwT_p@sSw0rd",
  "cronSpec": "0 20 * * *",
  "smtpHost": "smtp-mail.outlook.com",
  "smtpPort": 587,
  "smtpSubject": "My Flomo 每日回顾",
  "smtpUsername": "",
  "smtpPassword": ""
}
```

#### 2.2 根据你的 PostgreSQL 数据库配置 dsn（可选）

#### 2.3 根据你的邮箱配置 smtp（可选）

### 3 运行可执行文件

## 相关项目

* [Web 端](https://github.com/jerryshell/my-flomo-web)
* [服务端](https://github.com/jerryshell/my-flomo-server)

## 开源许可证

[GNU Affero General Public License v3.0](https://choosealicense.com/licenses/agpl-3.0/)
