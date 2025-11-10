# zsgf.client Go SDK

![Version](https://img.shields.io/github/v/tag/wz101010/zsgf-go?label=version)

一个功能完整的Go SDK，用于快速集成ZSGF平台的各种服务，包括用户认证、支付、存储、文件管理等功能。

## 📋 目录

- [快速开始](#快速开始)
- [安装配置](#安装配置)
- [基础用法](#基础用法)
- [主要功能](#主要功能)
- [配置说明](#配置说明)
- [API文档](#api文档)
- [常见问题](#常见问题)

## 🚀 快速开始

### 1. 安装依赖

```bash
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

### 2. 导入SDK

```go
import zsgf.client "//"
```

### 3. 基本使用示例

```go
package main

import (
    "context"
    zsgf.client "//"
)

func main() {
    // 创建配置
    cfg := zsgf.client.NewConfiguration()
    
    // 创建客户端
    client := zsgf.client.NewAPIClient(cfg)
    
    // 设置认证令牌
    ctx := context.WithValue(context.Background(), zsgf.client.ContextAccessToken, "your_token_here")
    
    // 调用API
    appInfo, _, err := client.AppAPI.AppInfo(ctx, "your_app_key")
    if err != nil {
        log.Fatal(err)
    }
    
    fmt.Printf("App Info: %+v\n", appInfo)
}
```

## 📦 安装配置

### 环境要求

- Go 1.16 或更高版本
- 有效的ZSGF应用密钥（AppKey）

### 代理配置

如果需要使用代理，可以设置环境变量：

```go
import "os"

os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## 💡 基础用法

### 客户端初始化

```go
// 基本配置
cfg := zsgf.client.NewConfiguration()

// 自定义服务器地址（可选）
cfg.Servers = []zsgf.client.ServerConfiguration{
    {
        URL: "https://your-custom-server.com",
    },
}

client := zsgf.client.NewAPIClient(cfg)
```

### 认证设置

```go
// 设置Bearer令牌
ctx := context.WithValue(context.Background(), zsgf.client.ContextAccessToken, "your_token")

// 使用带认证的上下文调用API
result, response, err := client.UserAPI.UserProfile(ctx, "your_app_key")
```

## 🎯 主要功能

### 👤 用户管理
- 用户注册/登录（邮箱、手机、密码）
- 个人资料管理
- 二次验证
- 外部账号绑定

```go
// 用户邮箱登录示例
loginReq := zsgf.client.EmailSignInRequest{
    Email:    zsgf.client.PtrString("user@example.com"),
    Password: zsgf.client.PtrString("password123"),
}

result, _, err := client.UserAPI.UserEmailSignIn(ctx, "app_key", loginReq)
```

### 💰 支付服务
- 支付宝支付（当面付、PC支付、WAP支付）
- 订单管理
- 退款处理

```go
// 创建支付宝订单示例
orderReq := zsgf.client.AlipayCreateOrderRequest{
    Subject:     zsgf.client.PtrString("商品名称"),
    TotalAmount: zsgf.client.PtrString("99.00"),
    OutTradeNo:  zsgf.client.PtrString("ORDER_" + time.Now().Format("20060102150405")),
}

order, _, err := client.AlipayAPI.AlipayCreateOrder(ctx, "app_key", orderReq)
```

### 📁 文件管理
- 文件上传/下载
- 文件夹操作
- 文件列表查询

```go
// 获取文件列表示例
files, _, err := client.FileAPI.Files(ctx, "app_key")
```

### 🗄️ 数据存储
- 数据增删改查
- 聚合查询
- 灵活的数据表操作

```go
// 存储数据示例
data := map[string]interface{}{
    "name": "张三",
    "age":  25,
}

result, _, err := client.StorageAPI.StoragePost(ctx, "app_key", "users", data)
```

## ⚙️ 配置说明

### 服务器配置

#### 选择不同服务器

```go
// 使用索引1的服务器
ctx := context.WithValue(context.Background(), zsgf.client.ContextServerIndex, 1)
```

#### 服务器变量配置

```go
// 配置服务器模板变量
ctx := context.WithValue(context.Background(), zsgf.client.ContextServerVariables, map[string]string{
    "basePath": "v2",
})
```

#### 针对特定操作的服务器配置

```go
// 为特定操作配置服务器
ctx := context.WithValue(context.Background(), zsgf.client.ContextOperationServerIndices, map[string]int{
    "UserAPIService.UserProfile": 2,
})
```

## 📚 API文档

### 核心API分类

| 分类 | 说明 | 主要功能 |
|------|------|----------|
| **用户管理** | UserAPI | 注册、登录、资料管理 |
| **认证授权** | OAuthAPI, AccessTokenAPI | OAuth认证、令牌管理 |
| **支付服务** | AlipayAPI, OrderAPI | 支付宝支付、订单管理 |
| **文件服务** | FileAPI | 文件上传、管理 |
| **数据存储** | StorageAPI | 数据增删改查 |
| **社交功能** | UserFriendsAPI | 关注、推荐、社交 |
| **微信服务** | WechatAPI | 小程序、公众号功能 |
| **位置服务** | UserLocationAPI | 地理位置管理 |
| **虚拟货币** | UserCurrencyAPI | 积分、虚拟币管理 |

### 常用API快速索引

<details>
<summary>👤 用户相关API</summary>

- `UserEmailSignIn` - 邮箱登录
- `UserPhoneSignIn` - 手机登录  
- `UserProfile` - 获取个人资料
- `UserUpdateProfile` - 更新个人资料
- `UserSendEmailCode` - 发送邮箱验证码
- `UserSendSMSCode` - 发送短信验证码

</details>

<details>
<summary>💰 支付相关API</summary>

- `AlipayCreateOrder` - 创建当面付订单
- `AlipayCreateOrderPagePay` - 创建PC支付订单
- `AlipayOrderDetail` - 获取订单详情
- `AlipayOrderRefund` - 发起订单退款

</details>

<details>
<summary>📁 文件相关API</summary>

- `FileUpload` - 上传文件
- `Files` - 获取文件列表
- `FileCreateFolder` - 创建文件夹
- `FileDelete` - 删除文件/文件夹

</details>

<details>
<summary>🗄️ 存储相关API</summary>

- `StoragePost` - 添加数据
- `StorageList` - 查询数据列表
- `StorageDetail` - 获取数据详情
- `StoragePut` - 更新数据
- `StorageDelete` - 删除数据

</details>

## ❓ 常见问题

### Q: 如何处理API错误？

```go
result, response, err := client.UserAPI.UserProfile(ctx, "app_key")
if err != nil {
    // 检查是否是API错误
    if response != nil {
        fmt.Printf("HTTP Status: %d\n", response.StatusCode)
        fmt.Printf("Response Body: %s\n", response.Body)
    }
    return err
}
```

### Q: 如何使用指针辅助函数？

SDK提供了便利的指针辅助函数：

```go
// 使用辅助函数创建指针
req := SomeRequest{
    Name:    zsgf.client.PtrString("张三"),
    Age:     zsgf.client.PtrInt(25),
    Active:  zsgf.client.PtrBool(true),
    Score:   zsgf.client.PtrFloat64(95.5),
}
```

### Q: 如何处理分页数据？

```go
// 获取存储数据（带分页）
ctx := context.WithValue(ctx, "page", 1)
ctx = context.WithValue(ctx, "limit", 20)

result, _, err := client.StorageAPI.StorageList(ctx, "app_key", "table_name")
```

### Q: 如何设置超时？

```go
import "time"

cfg := zsgf.client.NewConfiguration()
cfg.HTTPClient = &http.Client{
    Timeout: 30 * time.Second,
}
```

## 🔧 实用工具函数

SDK提供了便利的指针工具函数，用于处理API参数：

| 函数 | 说明 | 示例 |
|------|------|------|
| `PtrString(v)` | 字符串指针 | `PtrString("hello")` |
| `PtrInt(v)` | 整数指针 | `PtrInt(123)` |
| `PtrBool(v)` | 布尔值指针 | `PtrBool(true)` |
| `PtrFloat64(v)` | 浮点数指针 | `PtrFloat64(99.99)` |
| `PtrTime(v)` | 时间指针 | `PtrTime(time.Now())` |

## 📖 更多资源

- [完整API文档](docs/) - 查看所有API的详细文档
- [模型文档](docs/) - 了解所有数据模型结构
- [示例代码](examples/) - 更多使用示例
- [更新日志](CHANGELOG.md) - 版本更新记录

---

## 📋 附录

<details>
<summary>查看完整API列表</summary>

### 完整API端点列表

| 类 | 方法 | HTTP请求 | 描述 |
|---|---|---|---|
| *AccessTokenAPI* | [AccessTokenDelete](docs/AccessTokenAPI.md#accesstokendelete) | **Delete** /AccessToken/{appKey}/{id} | 删除令牌 |
| *AccessTokenAPI* | [AccessTokenPost](docs/AccessTokenAPI.md#accesstokenpost) | **Post** /AccessToken/{appKey} | 创建令牌 |
| *AccessTokenAPI* | [AccessTokenPut](docs/AccessTokenAPI.md#accesstokenput) | **Put** /AccessToken/{appKey}/{id} | 更新令牌 |
| *AccessTokenAPI* | [AccessTokens](docs/AccessTokenAPI.md#accesstokens) | **Get** /AccessToken/{appKey} | 令牌列表 |
| *AlipayAPI* | [AlipayCreateOrder](docs/AlipayAPI.md#alipaycreateorder) | **Post** /Alipay/{appKey}/CreateOrder | 创建当面付订单 |
| *AlipayAPI* | [AlipayCreateOrderPagePay](docs/AlipayAPI.md#alipaycreateorderpagepay) | **Post** /Alipay/{appKey}/CreateOrderPagePay | 创建PC支付订单 |
| *AlipayAPI* | [AlipayCreateOrderWapPay](docs/AlipayAPI.md#alipaycreateorderwappay) | **Post** /Alipay/{appKey}/CreateOrderWapPay | 创建WAP支付订单 |
| *AlipayAPI* | [AlipayOrderDetail](docs/AlipayAPI.md#alipayorderdetail) | **Get** /Alipay/{appKey}/OrderDetail | 获取订单详情 |
| *AlipayAPI* | [AlipayOrderRefund](docs/AlipayAPI.md#alipayorderrefund) | **Post** /Alipay/{appKey}/OrderRefund | 发起订单退款 |
| *AlipayAPI* | [AlipayReturnPageNotify](docs/AlipayAPI.md#alipayreturnpagenotify) | **Post** /Alipay/{appKey}/ReturnPageNotify | 支付成功回调通知 |
| *AppAPI* | [AppInfo](docs/AppAPI.md#appinfo) | **Get** /App/{appKey}/Info | 应用详情 |

*[其他API请参考原始列表...]*

</details>

<details>
<summary>查看完整模型列表</summary>

### 数据模型文档

主要模型包括：
- 用户相关：`User`, `UserProfile`, `GetUserProfileResult`
- 认证相关：`TokenModel`, `AuthorizeResult`, `GrantResult`
- 支付相关：`Order`, `AlipayTradeQueryResponse`, `CreateOrderResult`
- 文件相关：`FileItem`, `DirectoryItem`, `FileListResult`
- 存储相关：`StorageListResult`

*[完整模型列表请参考原始文档...]*

</details>