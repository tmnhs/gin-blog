## 这是一个用gin框架和vue框架搭建一个简易博客前台展示和后台管理系统

### 博客前台展示页面
<div align=center>
<img src="http://tmnhs.top/Fp7b_S-Awx1_VXXKy1diFWQ_H-Qz" width=“100%” />
</div>

### 博客后台管理页面
<div align=center>
<img src="http://tmnhs.top/FvtitX9vqMyYNDHaYpgH7dq8vg46" width=“100%” />
</div>

### 一.技术选型
-前端：用基于vue的``ant design vue``构建后台管理页面和基于vue的``vuetify``构建前台展示页面。

-后端：用``Gin``快速搭建基础restful风格API，Gin是一个go语言编写的Web框架。

-数据库：采用``MySql``，使用gorm实现对数据库的基本操作。

-缓存：使用``Redis``实现记录缓存。

-日志：使用``logrus``实现日志记录。

### 二.项目结构

```shell
├── api
│       └── v1
├── config
├── middleware
├── model
├── proto
├── routes
├── static
│       └── admin
│       └── front
└── utils
|       └── errmsg
|       └── validator
├── web
│       └── admin
│       └── front
```
| 文件夹           | 说明        | 描述               |
| ------------- | --------- | ---------------- |
| `api`         | api层      | api层             |
| `--v1`        | v1版本接口    | v1版本接口           |
| `config`      | 配置包       | 配置文件             |
| `middleware`  | 中间件层      | 用于存放 `gin` 中间件代码 |
| `model`       | 模型层       | 模型对应数据表和数据库查询    |
| `routes`      | 路由层       | 路由层              |
| `resource`    | 静态资源文件夹   | 负责存放静态文件         |
| `--admin`     | admin     | 后台管理dist文件打包     |
| `--front`     | front     | 前台展示dist文件打包     |
| `utils`       | 工具包       | 工具函数封装           |
| `--errmsg`    | errmsg    | 错误信息的封装          |
| `--validator` | validator | 后端数据校验           |
| `web`         | web       | 前端代码             |
| `--admin`     | admin     | 后台管理的前端代码        |
| `--front`     | front     | 前台展示的前端代码        |

###  三.使用说明

1-需要把config/config.ini里的文件配置(特别是数据库mysql的配置)修改成自己需要的配置，本项目使用七牛云对象存储上传的文件，你可以自己在七牛云注册一个账号，可以免费获赠10G的存储空间

2-本项目可以直接在windows上运行，建议使用goland

3-[前台展示]（http://localhost:8080/front）和 [后台管理] (http://localhost:8080/admin) 的切换需要修改routes/router.go,打开相关注释即可

4-如果需要修改前端部分，可以修改web下的文件，修改完后运行`npm run build`，把dist文件覆盖掉static里的文件

5-本项目是前后端分离项目，可以注释掉routes/router.go中加载静态资源的代码，把后端代码运行起来，然后在web/admin(或者web/front)目录下，运行`npm run serve`即可

###  四.部署项目（Linux系统下）

方式一：使用dockerfile部署(确保服务器上有下载docker)

```shell
#在MyBlog目录下
cd MyBlog
#docker编译镜像
docker build -t myblog .
#运行docker
docker run -p 8080:8080 -t --name myblog myblog
#注意此时的数据库配置应该改为服务器上的配置
```

方式二：使用脚本部署

```shell
#在MyBlog目录下
cd MyBlog
#必须在linux系统下
#启动项目
./serve.sh start
#暂停项目
./serve.sh stop
#重启服务
./serve.sh restart
#注意，脚本运行的端口号必须大于1024，不然可能会没有权限
```

