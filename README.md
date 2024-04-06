![my-flomo-server](https://socialify.git.ci/jerryshell/my-flomo-server/image?description=1&forks=1&issues=1&language=1&name=1&owner=1&pattern=Brick%20Wall&pulls=1&stargazers=1&theme=Dark)

## 功能清单

* Flomo 数据导入
* Flomo API 兼容
* 邮件每日回顾
* 注销账号，永久抹除数据
* CSV 数据导入导出
* Docker 镜像支持 ARMv7（树莓派）

## 体验 Demo

~~[https://my-flomo.d8s.fun](https://my-flomo.d8s.fun)~~

**服务器到期，体验 Demo 的后端已经关闭**

## 如何运行

### Docker

```bash
# 0. 首先提前准备好 PostgreSQL

# 1. 创建并进入 my-flomo-server 目录
mkdir my-flomo-server && cd my-flomo-server

# 2. 下载 docker-compose.yaml
wget https://raw.githubusercontent.com/jerryshell/my-flomo-server/master/docker-compose.yaml

# 3. 根据自己的环境修改配置
vim docker-compose.yaml

# 4. 启动服务
docker-compose up -d
```

### K8s

具体请参考 [k8s/*.yaml](k8s)，同样也要注意修改 `deployment.yaml` 的环境配置、`ingress.yaml` 的域名

局域网的 TLS Ingress 可以参考 [K8s Traefik cert-manager DNS01 TLS](https://github.com/jerryshell/k8s-traefik-cert-manager-dns01-tls)

## 相关项目

* [Web 端](https://github.com/jerryshell/my-flomo-web)
* [服务端](https://github.com/jerryshell/my-flomo-server)

## 开源协议

[GNU Affero General Public License v3.0](https://choosealicense.com/licenses/agpl-3.0/)
