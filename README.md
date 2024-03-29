![my-flomo-server](https://socialify.git.ci/jerryshell/my-flomo-server/image?description=1&forks=1&issues=1&language=1&name=1&owner=1&pattern=Brick%20Wall&pulls=1&stargazers=1&theme=Dark)

## 功能清单

* Flomo 数据导入
* Flomo API 兼容
* 邮件每日回顾
* 注销账号，永久抹除数据
* CSV 数据导入导出
* 服务端支持 ARMv7 部署

## 体验 Demo

**⚠ 注意 ⚠ 推荐每个用户单独搭建自己的服务，体验 Demo 不保证数据安全性，所以请勿在体验 Demo 中使用真实用户名密码注册！请勿在体验
Demo 中录入敏感数据！**

~~[https://my-flomo.pages.dev](https://my-flomo.pages.dev/)~~ **服务器到期，体验 Demo 已经关闭**

## 如何运行

```bash
# 0. 首先提前准备好 PostgreSQL
# 1. 创建并进入 my-flomo-server 目录
mkdir my-flomo-server && cd my-flomo-server

# 2. 下载 config.json
# 国际网络连接顺畅
wget https://raw.githubusercontent.com/jerryshell/my-flomo-server/master/config.json
# 否则
wget https://raw.fastgit.org/jerryshell/my-flomo-server/master/config.json

# 3. 下载 docker-compose.yaml
# 国际网络连接顺畅
wget https://raw.githubusercontent.com/jerryshell/my-flomo-server/master/docker-compose.yaml
# 否则
wget https://raw.fastgit.org/jerryshell/my-flomo-server/master/docker-compose.yaml

# 4. 根据自己的环境修改配置
vim config.json

# 5. 启动服务
docker-compose up -d
```

## 相关项目

* [Web 端](https://github.com/jerryshell/my-flomo-web)
* [服务端](https://github.com/jerryshell/my-flomo-server)

## 开源许可证

[GNU Affero General Public License v3.0](https://choosealicense.com/licenses/agpl-3.0/)
