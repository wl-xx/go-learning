# TodoList 接口清单

前端开发服务器会将以 `/api` 开头的请求代理到 `http://localhost:3000`。因此，浏览器中的实际请求路径为 `/api/...`，后端实现的路由路径不包含 `/api` 前缀。

## 通用约定

- 数据交换格式：`application/json; charset=utf-8`。
- 所有接口均返回 JSON。
- 成功响应统一使用 `success: true`，失败响应统一使用 `success: false` 和 `msg` 字段。
- `id` 由服务端在创建待办时生成，必须是十进制正整数字符串，且新创建的 ID 必须大于已有 ID。例如当前最大 ID 为 `"3"` 时，下一条应为 `"4"`。
- 前端会始终按 `id` 数值降序展示待办；因此 ID 越大，待办越靠上。建议后端列表接口也按同一顺序返回。
- `content` 去除首尾空白后不能为空，建议最长 200 个字符。

### 待办数据结构

```json
{
  "id": "1",
  "content": "学习 Go 语言",
  "completed": false
}
```

### 成功响应结构

```json
{
  "success": true,
  "data": {}
}
```

### 失败响应结构

```json
{
  "success": false,
  "msg": "错误说明"
}
```

## 1. 获取待办列表

```http
GET /list
```

成功时返回 `200 OK`：

`data` 中的待办应按 `id` 数值降序排列。

```json
{
  "success": true,
  "data": [
    {
      "id": "1",
      "content": "学习 Go 语言",
      "completed": false
    }
  ]
}
```

## 2. 新建待办

```http
POST /todo
Content-Type: application/json
```

请求体：

```json
{
  "content": "完成前端页面"
}
```

成功时返回 `201 Created`，`data` 为服务端保存后的完整待办对象：

```json
{
  "success": true,
  "data": {
    "id": "4",
    "content": "完成前端页面",
    "completed": false
  }
}
```

参数不合法时返回 `400 Bad Request`：

```json
{
  "success": false,
  "msg": "待办内容不能为空"
}
```

## 3. 更新完成状态

```http
PATCH /todo/:id
Content-Type: application/json
```

请求体：

```json
{
  "completed": true
}
```

成功时返回 `200 OK`，`data` 为更新后的完整待办对象：

```json
{
  "success": true,
  "data": {
    "id": "4",
    "content": "完成前端页面",
    "completed": true
  }
}
```

待办不存在时返回 `404 Not Found`：

```json
{
  "success": false,
  "msg": "待办不存在"
}
```

## 4. 删除单个待办

```http
DELETE /todo/:id
```

成功时返回 `200 OK`：

```json
{
  "success": true,
  "data": null
}
```

待办不存在时返回 `404 Not Found`。

## 5. 清除全部已完成待办

```http
DELETE /completed
```

成功时返回 `200 OK`，`data` 为清除后剩余的待办列表：

```json
{
  "success": true,
  "data": [
    {
      "id": "1",
      "content": "学习 Go 语言",
      "completed": false
    }
  ]
}
```

即使当前没有已完成待办，也应返回成功和当前列表。

## 前端调用映射

| 页面操作 | 方法 | 后端路径 | 前端更新方式 |
| --- | --- | --- | --- |
| 首次加载、重新加载 | `GET` | `/list` | 以响应列表替换当前列表后，按 ID 数值降序展示 |
| 添加待办 | `POST` | `/todo` | 将响应中的待办插入列表顶部，并按 ID 数值降序展示 |
| 勾选或取消勾选 | `PATCH` | `/todo/:id` | 以响应中的待办替换对应项，保持 ID 数值降序 |
| 删除待办 | `DELETE` | `/todo/:id` | 请求成功后移除对应项 |
| 清除已完成 | `DELETE` | `/completed` | 以响应中的剩余列表替换当前列表后，按 ID 数值降序展示 |
