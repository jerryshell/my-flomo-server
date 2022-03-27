![my-flomo-server](https://socialify.git.ci/jerryshell/my-flomo-server/image?description=1&descriptionEditable=%E5%8F%AF%E4%BB%A5%E8%87%AA%E5%B7%B1%E6%90%AD%E5%BB%BA%E7%9A%84%E6%83%B3%E6%B3%95%E8%AE%B0%E5%BD%95%E5%8F%8A%E5%9B%9E%E9%A1%BE%E6%9C%8D%E5%8A%A1%EF%BC%8C%E5%BC%80%E6%BA%90%E3%80%81%E5%85%8D%E8%B4%B9%E3%80%81%E7%AE%80%E5%8D%95%E3%80%81%E4%B8%93%E6%B3%A8%E4%BA%8E%E6%A0%B8%E5%BF%83%E5%8A%9F%E8%83%BD&font=Raleway&forks=1&issues=1&language=1&owner=1&pattern=Brick%20Wall&pulls=1&stargazers=1&theme=Dark)

## 功能清单

* Flomo 数据导入
* Flomo API 兼容
* 邮件每日回顾
* 注销账号，永久抹除数据
* CSV 数据导入导出

## 体验 Demo

**⚠️ 注意：推荐每个用户单独搭建自己的服务，体验 Demo 不保证数据安全性，所以请勿在体验 Demo 中使用真实用户名密码注册！请勿在体验 Demo 中录入敏感数据！**

[https://my-flomo.pages.dev/](https://my-flomo.pages.dev/)

## 如何运行

0. 首先提前准备好 PostgreSQL

```bash
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
