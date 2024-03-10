# 项目简介

本项目是对原个人 Java 项目 【qyblog-java】([github](https://github.com/iwinder/qyblog-java) | [gitee](https://gitee.com/windcoder/qyblog-java)) 的 GoLang 版改造， 由于时间有限，这一版直接使用了B站的开源微服务框架 [Kratos](https://go-kratos.dev/) 。

配套前端项目 【qyblog-web】采用 [vue3](https://cn.vuejs.org/) + [Ant Design Vue](https://www.antdv.com/)重构，使用 [nuxt3](https://nuxt.com/) 实现 ssr，小程序部分使用 [uni-app](https://en.uniapp.dcloud.io/) 重构。

前端项目地址：
- github ：[https://github.com/iwinder/qyblog-web](https://github.com/iwinder/qyblog-web)
- gitee ：[https://gitee.com/windcoder/qyblog-web](https://gitee.com/windcoder/qyblog-web)

### 当前功能
项目目前以个人博客自用为主，加上微服务化会增加实际运行成本，短期内不会有微服务化的考虑。目前实现了如下的基础功能。
- 文章系统
  - 文章：基本的文章增删改查管理
  - 页面：单独的特殊页面，不会展示在文章列表中，需要手动配置在菜单。
  - 分类：文章所属分类
  - 标签：文章所属标签
  - 评论：小型的评论管理，当有新评论时提供邮件提醒。
- 媒体库管理
  - 媒体库配置(本地、七牛、阿里)
  - 媒体库文件展示：目前只实现了本地和七牛云的。
- 链接管理
  - 导航：默认的只有顶部和页脚菜单
  - 友联：目前仅支持后端手动添加
  - 短链接：方便将网址隐藏成 `http://xxx/go/xxx` 的格式
- 系统管理
  - 用户管理：默认只有管理员，暂未设计对外开放注册的功能
  - 角色管理：目前权限粒度控制在角色层，本身使用 [casbin](https://casbin.org/) 实现 RBAC权限模型，方便后期扩展。
  - 菜单管理：管理端页面的菜单管理，与权限管理配套使用。
  - API管理：管理端页面的涉及的接口api管理，因为前后端分离，需要与权限管理配套使用。
  - 站点管理：涉及站点名称、页头页脚自定义代码添加等定制化配置管理。
- 定时任务
  - 更新网站地图：网站的sitemap文件定时更新，有需要时可手动调用更新。
  - 更新评论统计：评论系统设计多表，设计相对复杂，前端页面的评论总数处于延迟定时更新，有需要时可手动调用更新。
  - 更新文章浏览量统计：浏览量本身记录在 redis 缓存中定时同步，有需要时可手动调用更新。
  - 推送回复邮件消息：博主的回复消息会定时发送给留言者，有需要时可手动调用更新，但这个的手动推送一般用不到。
### 未来功能
部分待做功能
- 细粒度权限控制：目前只支持到了菜单，暂时未做按钮级别的前端控制逻辑。
- 专题功能
- 微语功能：类似说说那些简短的一句话
- 电子书单/书架
- MetaWebLog API：一种文章发布思路，实现后可用 MWeb 直接发文，同时 CSDN 等多数博客支持这一协议，可以扩展相关同步发布等功能。

## 单体服务架构
项目目录及作用如下，属于标准的 Kratos 的单体应用模板，基本上是遵循 [project-layout](https://github.com/golang-standards/project-layout) 的。
- api: 所有所用接口均在这里生成，由于基于 grpc，所以也是 dao 所在位置
- cmd：用于启动命令等生成
- configs：配置文件所在位置
- doc：文档
  - api：自动生成 API JSON 文件
  - sql：数据库文件
  - templates：邮件等模板
  - configs: 一些配置文件，目前将 casbin 规则文件放在了这里
- internal 服务实现所在位置
    - biz：定义 do,具体业务层，DO 与 PO 互转层
    - conf: 定义读取配置文件的对象。
    - data：定义 PO，数据库连接创建以及增删改查等基本服务， PO服务执行层。
    - job：定时任务
    - server: http与grpc协议服务实际注册层，如果需要多个api文件中请求注册到项目中，则需要调整这里面的http或者 grpc 文件，增加类似 `v1.RegisterUserHTTPServer(srv, userService)`的语句。
    - service：dao 与 do 互转层

## API 生成
```shell
# 创建api
kratos proto add api/qycms_bff/admin/v1/qy_admin_api_group.proto
kratos proto add api/qycms_bff/web/v1/qy_web_article.proto
# 生成api
kratos proto client api/
kratos proto client  internal/qycms_blog/conf/conf.proto
```
## API JSON文件生成

暂时需要手动在 Makefile 文件中手动配置需要生成 API 的文件，然后执行 `make swagger` 会在对应的文件夹下面生成json文件

暂时将目前配置的部分接口文档复制到了 doc/api 文件夹下面一份。

# Kratos Project Template

## Install Kratos
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```
## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/store/conf <your-docker-image-name>
```

