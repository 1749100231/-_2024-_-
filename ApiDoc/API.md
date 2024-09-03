---
title: 个人项目
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# 个人项目

Base URLs:

# Authentication

# 用户界面

## POST 登录

POST /api/user/login

用户登录接口，根据账号不同，登录响应结果区分管理员和学生登录
前端要求登录成功后能跳转到不同的界面
后端要求可以返回不同的type码来区分

> Body 请求参数

```json
{
  "username": "123",
  "password": "12345678"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» username|body|string| 是 | 用户名|使用的是学号|
|» password|body|string| 是 | 密码|none|

> 返回示例

```json
{
  "code": 200,
  "data": {
    "user_id": 2,
    "user_type": 1
  },
  "msg": "success"
}
```

```json
{
  "code": 200506,
  "data": null,
  "msg": "用户不存在"
}
```

```json
{
  "code": 200507,
  "data": null,
  "msg": "密码错误"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object¦null|true|none||none|
|»» user_id|integer|true|none|用户编号|none|
|»» user_type|integer|true|none|用户类型|1代表学生，2代表管理员|

## POST 注册

POST /api/user/reg

生成学生和管理员账户
前端不需要实现这个接口
后端要求实现

> Body 请求参数

```json
{
  "username": "123",
  "name": "ximo",
  "password": "12345678",
  "user_type": 1
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» username|body|string| 是 | 账号|输入学号或管理工号，只由数字组成|
|» name|body|string| 是 | 姓名|none|
|» password|body|string| 是 | 密码|md5哈希，32位，小写字母|
|» user_type|body|integer| 是 | 用户类型|1代表学生，2代表管理员|
|» admin_pwd|body|string¦null| 否 | 管理员密码|仅注册管理员时所需的密码|

> 返回示例

```json
{
  "code": 200,
  "data": null,
  "msg": "success"
}
```

```json
{
  "code": 200502,
  "data": null,
  "msg": "用户名必须为纯数字"
}
```

```json
{
  "code": 200503,
  "data": null,
  "msg": "密码长度必须在8-16位"
}
```

```json
{
  "code": 200504,
  "data": null,
  "msg": "用户类型错误"
}
```

```json
{
  "code": 200505,
  "data": null,
  "msg": "用户名已存在"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|null|true|none||none|

# 学生界面

## POST 发布帖子

POST /api/student/post

学生发帖接口
前端要求发帖成功后清空输入框，失败要有用户友好提醒
后端要求发帖时记录当前时间

> Body 请求参数

```json
{
  "content": "string",
  "user_id": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» content|body|string| 是 | 帖子内容|none|
|» user_id|body|integer| 是 | 学生编号|none|

> 返回示例

```json
{
  "code": 200,
  "data": null,
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|null|true|none||none|

## GET 获取所有发布的帖子

GET /api/student/post

学生获取所有发布的帖子
前端要求展示出已经发出的帖子
后端要求返回帖子列表

> 返回示例

```json
{
  "code": 200,
  "data": {
    "post_list": [
      {
        "id": 1,
        "content": "123",
        "user_id": 2,
        "time": "2024-07-19T05:13:34.779+08:00"
      },
      {
        "id": 2,
        "content": "1233",
        "user_id": 3,
        "time": "2024-07-19T05:13:50.038+08:00"
      }
    ]
  },
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» post_list|[object]|true|none||none|
|»»» article_id|integer|true|none|帖子id|none|
|»»» user_id|integer|true|none|用户编号|none|
|»»» content|string|true|none|帖子内容|none|
|»»» create_time|string|true|none|创建时间|none|
|»»» update_time|string|true|none|更新时间|none|

## DELETE 删除帖子

DELETE /api/student/post

学生删除自己发布的帖子

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|user_id|query|integer| 是 ||用户编号|
|post_id|query|integer| 是 ||帖子编号|

> 返回示例

```json
{
  "code": 200,
  "data": null,
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|null|true|none||none|

## PUT 修改帖子

PUT /api/student/post

学生修改自己发布的帖子

> Body 请求参数

```json
{
  "user_id": 0,
  "post_id": 0,
  "content": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» user_id|body|integer| 是 | 用户编号|none|
|» post_id|body|integer| 是 | 帖子编号|none|
|» content|body|string| 是 | 帖子内容|none|

> 返回示例

```json
{
  "code": 200,
  "data": null,
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|null|true|none||none|

## POST 举报帖子

POST /api/student/report-post

学生举报违禁帖子

> Body 请求参数

```json
{
  "user_id": 0,
  "post_id": 0,
  "reason": "string"
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» user_id|body|integer| 是 | 用户编号|none|
|» post_id|body|integer| 是 | 帖子编号|none|
|» reason|body|string| 是 | 举报原因|none|

> 返回示例

```json
{
  "code": 200,
  "data": null,
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|null|true|none||none|

## GET 查看举报审批

GET /api/student/report-post

学生可以查看举报给管理员的帖子的审核结果

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|user_id|query|integer| 否 ||none|

> 返回示例

```json
{
  "code": 200,
  "data": {
    "report_list": [
      {
        "post_id": 2,
        "content": "1233",
        "reason": "123",
        "status": 0
      },
      {
        "post_id": 3,
        "content": "12333",
        "reason": "1233",
        "status": 1
      }
    ]
  },
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» report_list|[object]|true|none||none|
|»»» post_id|integer|true|none|帖子编号|none|
|»»» content|string|true|none|帖子内容|none|
|»»» reason|string|true|none|举报原因|none|
|»»» status|integer|true|none|审批状态|0代表未审批1代表已通过2代表不通过|

# 管理员界面

## GET 获取所有未审批的被举报帖子

GET /api/admin/report

管理员可以获取所有学生发来未审批的举报帖子列表

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|user_id|query|integer| 是 ||用户编号|

> 返回示例

```json
{
  "code": 200,
  "data": {
    "report_list": [
      {
        "username": "110",
        "content": "会长好帅",
        "reason": "尊嘟假嘟",
        "post_id": 1
      },
      {
        "username": "120",
        "content": "会长真帅",
        "reason": "尊嘟",
        "post_id": 2
      }
    ]
  },
  "msg": "success"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|object|true|none||none|
|»» report_list|[object]|true|none||举报帖子列表|
|»»» username|string|true|none||用户名（学号）|
|»»» post_id|integer|true|none||帖子编号|
|»»» content|string|true|none||帖子内容|
|»»» reason|string|true|none||举报原因|

## POST 审核被举报的帖子

POST /api/admin/report

管理员审批被举报的帖子，审批通过会删除被举报帖子

> Body 请求参数

```json
{
  "user_id": 0,
  "post_id": 0,
  "approval": 0
}
```

### 请求参数

|名称|位置|类型|必选|中文名|说明|
|---|---|---|---|---|---|
|body|body|object| 否 ||none|
|» user_id|body|integer| 是 | 用户编号|none|
|» post_id|body|integer| 是 | 帖子编号|none|
|» approval|body|integer| 是 | 状态|1代表同意，2代表拒绝|

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": null
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» code|integer|true|none||none|
|» msg|string|true|none||none|
|» data|null|true|none||none|

# 数据模型

