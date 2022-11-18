## blog

> gin + vue

npm config set registry https://registry.npm.taobao.org

> go 编译相关文档

https://www.cnblogs.com/binHome/p/14845617.html

## 启动

npm install --legacy-peer-deps  // 忽略一些版本问题

## 目录结构分析

```

-- server   后端目录
---- pkg 配置一些包
---- dto 数据传输层
---- app    应用程序
------ conf 配置文件
------ controllers 控制器
------ dao 连接 mysql 配置
------ models 初始化模型
------ setting 数据库配置
---- routes 路由
------ user 路由组

```


> 创建一个 vite + vue 项目

后端返回的 data 结构

> 登录

```json
{
  "code": 200, // 状态码
  "msg": "success", // 返回信息
  "data": {   // 数据
    "accessToken": "admin-accessToken" // token
  }
}

```

> 用户信息

```json

{
  "code": 200,
  "msg": "success",
  "data": {
    "roles": [
      "admin"
    ],
    "ability": [
      "READ",
      "WRITE",
      "DELETE"
    ],
    "username": "admin",
    "avatar": "https://i.gtimg.cn/club/item/face/img/2/15922_100.gif"
  }
}

```

> 退出登录

```json

{
  "code": 200,
  "msg": "success"
}

```

> 获取表格信息

```json

{
  "code": 200,
  "msg": "success",
  "total": 50,
  "data": [
    {
      "uuid": "0bCCA780-DD72-AbbA-6C77-74dEF5Ab93D0",
      "id": "350000200301256314",
      "title": "Rwvbpnjs Mtnpsevsb",
      "description": "家用己斯门率千目育也然机验现取但号。",
      "status": "draft",
      "author": "潘勇",
      "datetime": "1990-06-19 18:58:28",
      "pageViews": 3229,
      "img": "https://picsum.photos/228/228?random=F11A5a2A-1Ad3-AbAf-FE1c-eFFfFC8CdBBF",
      "switch": true,
      "percent": 94,
      "rate": 2
    }
  ]
}

```