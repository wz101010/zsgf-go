 # zsgf.client

![Version](https://img.shields.io/github/v/tag/wz101010/zsgf-go?label=version)

## 安装

安装以下依赖项：

```sh
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

将软件包放在项目文件夹下，并在import中添加以下内容：

```go
import zsgf.client "//"
```

要使用代理，请设置环境变量`HTTP_PROXY`：

```go
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## 服务器URL配置

默认配置带有`Servers`字段，其中包含OpenAPI规范中定义的服务器对象。

### 选择服务器配置

要使用索引0以外的其他服务器，请设置类型为`int`的上下文值`zsgf.client.ContextServerIndex`。

```go
ctx := context.WithValue(context.Background(), zsgf.client.ContextServerIndex, 1)
```

### 模板化服务器URL

模板化服务器URL使用配置中的默认变量或类型为`map[string]string`的上下文值`zsgf.client.ContextServerVariables`进行格式化。

```go
ctx := context.WithValue(context.Background(), zsgf.client.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

注意，枚举值始终会被验证，所有未使用的变量都会被静默忽略。

### 每个操作的URL配置

每个操作可以使用`Configuration`中的`OperationServers`映射定义不同的服务器URL。
操作由`"{classname}Service.{nickname}"`字符串唯一标识。
通过使用`zsgf.client.ContextOperationServerIndices`和`zsgf.client.ContextOperationServerVariables`上下文映射，可以应用类似的覆盖默认操作服务器索引和变量的规则。

```go
ctx := context.WithValue(context.Background(), zsgf.client.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), zsgf.client.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## API端点文档

类 | 方法 | HTTP请求 | 描述
------------ | ------------- | ------------- | -------------
*AccessTokenAPI* | [**AccessTokenDelete**](docs/AccessTokenAPI.md#accesstokendelete) | **Delete** /AccessToken/{appKey}/{id} | 删除令牌
*AccessTokenAPI* | [**AccessTokenPost**](docs/AccessTokenAPI.md#accesstokenpost) | **Post** /AccessToken/{appKey} | 创建令牌
*AccessTokenAPI* | [**AccessTokenPut**](docs/AccessTokenAPI.md#accesstokenput) | **Put** /AccessToken/{appKey}/{id} | 更新令牌
*AccessTokenAPI* | [**AccessTokens**](docs/AccessTokenAPI.md#accesstokens) | **Get** /AccessToken/{appKey} | 令牌列表
*AlipayAPI* | [**AlipayCreateOrder**](docs/AlipayAPI.md#alipaycreateorder) | **Post** /Alipay/{appKey}/CreateOrder | 创建订单 - 当面付
*AlipayAPI* | [**AlipayCreateOrderPagePay**](docs/AlipayAPI.md#alipaycreateorderpagepay) | **Post** /Alipay/{appKey}/CreateOrderPagePay | 创建订单 - PC支付
*AlipayAPI* | [**AlipayCreateOrderWapPay**](docs/AlipayAPI.md#alipaycreateorderwappay) | **Post** /Alipay/{appKey}/CreateOrderWapPay | 创建订单 - WAP支付
*AlipayAPI* | [**AlipayOrderDetail**](docs/AlipayAPI.md#alipayorderdetail) | **Get** /Alipay/{appKey}/OrderDetail | 订单详情
*AlipayAPI* | [**AlipayOrderRefund**](docs/AlipayAPI.md#alipayorderrefund) | **Post** /Alipay/{appKey}/OrderRefund | 订单退款
*AlipayAPI* | [**AlipayReturnPageNotify**](docs/AlipayAPI.md#alipayreturnpagenotify) | **Post** /Alipay/{appKey}/ReturnPageNotify | 支付成功同步通知
*AppAPI* | [**App**](docs/AppAPI.md#app) | **Get** /App/{appKey} | 应用详情
*AppAPI* | [**App2FA**](docs/AppAPI.md#app2fa) | **Get** /App/{appKey}/2FA | 双因素验证 获取
*AppAPI* | [**App2FACheck**](docs/AppAPI.md#app2facheck) | **Get** /App/{appKey}/2FACheck | 双因素验证 验证
*AppAPI* | [**AppCheckVersion**](docs/AppAPI.md#appcheckversion) | **Get** /App/{appKey}/CheckVersion | 更新应用数据库
*AppAPI* | [**AppDelete**](docs/AppAPI.md#appdelete) | **Delete** /App/{appKey} | 删除应用
*AppAPI* | [**AppInfo**](docs/AppAPI.md#appinfo) | **Get** /App/{appKey}/Info | 应用详情
*AppAPI* | [**AppPost**](docs/AppAPI.md#apppost) | **Post** /App | 创建应用
*AppAPI* | [**AppPut**](docs/AppAPI.md#appput) | **Put** /App/{appKey} | 更新应用
*AppAPI* | [**AppTransfer**](docs/AppAPI.md#apptransfer) | **Get** /App/{appKey}/Transfer | 转移应用
*AppAPI* | [**Apps**](docs/AppAPI.md#apps) | **Get** /App | 应用列表
*AppSettingAPI* | [**AppServiceSettingGroup**](docs/AppSettingAPI.md#appservicesettinggroup) | **Get** /AppSetting/{appKey}/Groups/{id} | 获取服务分组详情
*AppSettingAPI* | [**AppServiceSettingGroupDelete**](docs/AppSettingAPI.md#appservicesettinggroupdelete) | **Delete** /AppSetting/{appKey}/Groups/{id} | 删除服务分组
*AppSettingAPI* | [**AppServiceSettingGroupPost**](docs/AppSettingAPI.md#appservicesettinggrouppost) | **Post** /AppSetting/{appKey}/Groups | 添加服务分组
*AppSettingAPI* | [**AppServiceSettingGroupPut**](docs/AppSettingAPI.md#appservicesettinggroupput) | **Put** /AppSetting/{appKey}/Groups/{id} | 更新服务分组
*AppSettingAPI* | [**AppServiceSettingGroups**](docs/AppSettingAPI.md#appservicesettinggroups) | **Get** /AppSetting/{appKey}/Groups | 获取服务分组列表
*AppSettingAPI* | [**AppServiceSettingItem**](docs/AppSettingAPI.md#appservicesettingitem) | **Get** /AppSetting/{appKey}/Items/{id} | 服务配置项详情
*AppSettingAPI* | [**AppServiceSettingItemDelete**](docs/AppSettingAPI.md#appservicesettingitemdelete) | **Delete** /AppSetting/{appKey}/Items/{id} | 删除服务配置项
*AppSettingAPI* | [**AppServiceSettingItemPost**](docs/AppSettingAPI.md#appservicesettingitempost) | **Post** /AppSetting/{appKey}/Items | 添加服务配置项
*AppSettingAPI* | [**AppServiceSettingItemPut**](docs/AppSettingAPI.md#appservicesettingitemput) | **Put** /AppSetting/{appKey}/Items/{id} | 更新服务配置项
*AppSettingAPI* | [**AppServiceSettingItems**](docs/AppSettingAPI.md#appservicesettingitems) | **Get** /AppSetting/{appKey}/Items | 服务配置项列表
*AppSettingAPI* | [**AppServiceSettingProvider**](docs/AppSettingAPI.md#appservicesettingprovider) | **Get** /AppSetting/{appKey}/Providers/{id} | 获取服务商详情
*AppSettingAPI* | [**AppServiceSettingProviderDelete**](docs/AppSettingAPI.md#appservicesettingproviderdelete) | **Delete** /AppSetting/{appKey}/Providers/{id} | 删除服务商
*AppSettingAPI* | [**AppServiceSettingProviderPost**](docs/AppSettingAPI.md#appservicesettingproviderpost) | **Post** /AppSetting/{appKey}/Providers | 添加服务商
*AppSettingAPI* | [**AppServiceSettingProviderPut**](docs/AppSettingAPI.md#appservicesettingproviderput) | **Put** /AppSetting/{appKey}/Providers/{id} | 更新服务商
*AppSettingAPI* | [**AppServiceSettingProviders**](docs/AppSettingAPI.md#appservicesettingproviders) | **Get** /AppSetting/{appKey}/Providers | 获取服务商列表
*AppSettingAPI* | [**AppSetting**](docs/AppSettingAPI.md#appsetting) | **Get** /AppSetting/{appKey}/{id} | 配置详情
*AppSettingAPI* | [**AppSettingDelete**](docs/AppSettingAPI.md#appsettingdelete) | **Delete** /AppSetting/{appKey}/{id} | 删除配置
*AppSettingAPI* | [**AppSettingPost**](docs/AppSettingAPI.md#appsettingpost) | **Post** /AppSetting/{appKey} | 增加配置
*AppSettingAPI* | [**AppSettingPut**](docs/AppSettingAPI.md#appsettingput) | **Put** /AppSetting/{appKey}/{id} | 更新配置
*AppSettingAPI* | [**AppSettings**](docs/AppSettingAPI.md#appsettings) | **Get** /AppSetting/{appKey} | 配置列表
*AuthorizePolicyAPI* | [**AuthorizePolicies**](docs/AuthorizePolicyAPI.md#authorizepolicies) | **Get** /AuthorizePolicy/{appKey} | 获取鉴权策略列表
*AuthorizePolicyAPI* | [**AuthorizePolicy**](docs/AuthorizePolicyAPI.md#authorizepolicy) | **Get** /AuthorizePolicy/{appKey}/{id} | 获取鉴权策略详情
*AuthorizePolicyAPI* | [**AuthorizePolicyDelete**](docs/AuthorizePolicyAPI.md#authorizepolicydelete) | **Delete** /AuthorizePolicy/{appKey}/{id} | 删除鉴权策略
*AuthorizePolicyAPI* | [**AuthorizePolicyPost**](docs/AuthorizePolicyAPI.md#authorizepolicypost) | **Post** /AuthorizePolicy/{appKey} | 添加鉴权策略
*AuthorizePolicyAPI* | [**AuthorizePolicyPut**](docs/AuthorizePolicyAPI.md#authorizepolicyput) | **Put** /AuthorizePolicy/{appKey}/{id} | 更新鉴权策略
*CurrencyAPI* | [**Currencies**](docs/CurrencyAPI.md#currencies) | **Get** /Currency/{appKey} | 获取货币列表
*CurrencyAPI* | [**Currency**](docs/CurrencyAPI.md#currency) | **Get** /Currency/{appKey}/{id} | 获取货币详情
*CurrencyAPI* | [**CurrencyDelete**](docs/CurrencyAPI.md#currencydelete) | **Delete** /Currency/{appKey}/{id} | 删除货币
*CurrencyAPI* | [**CurrencyExchangeRateDelete**](docs/CurrencyAPI.md#currencyexchangeratedelete) | **Delete** /Currency/{appKey}/ExchangeRates/{id} | 删除汇率
*CurrencyAPI* | [**CurrencyExchangeRatePut**](docs/CurrencyAPI.md#currencyexchangerateput) | **Put** /Currency/{appKey}/ExchangeRates/{code} | 更新汇率
*CurrencyAPI* | [**CurrencyExchangeRates**](docs/CurrencyAPI.md#currencyexchangerates) | **Get** /Currency/{appKey}/ExchangeRates/{code} | 获取汇率列表
*CurrencyAPI* | [**CurrencyPost**](docs/CurrencyAPI.md#currencypost) | **Post** /Currency/{appKey} | 创建新货币
*CurrencyAPI* | [**CurrencyPut**](docs/CurrencyAPI.md#currencyput) | **Put** /Currency/{appKey}/{id} | 更新货币信息
*CurrencyAPI* | [**CurrencyTransactions**](docs/CurrencyAPI.md#currencytransactions) | **Get** /Currency/{appKey}/Transactions | 获取货币交易记录
*DingTalkAPI* | [**DingTalkUserInfo**](docs/DingTalkAPI.md#dingtalkuserinfo) | **Get** /DingTalk/{appKey}/UserInfo | 获取用户资料
*FileAPI* | [**FileCreateFolder**](docs/FileAPI.md#filecreatefolder) | **Post** /File/{appKey}/CreateFolder | 创建文件夹
*FileAPI* | [**FileDelete**](docs/FileAPI.md#filedelete) | **Delete** /File/{appKey} | 删除文件或文件夹
*FileAPI* | [**FileRename**](docs/FileAPI.md#filerename) | **Post** /File/{appKey}/Rename | 重命名文件或文件夹
*FileAPI* | [**FileUpload**](docs/FileAPI.md#fileupload) | **Post** /File/{appKey}/Upload | 上传文件
*FileAPI* | [**Files**](docs/FileAPI.md#files) | **Get** /File/{appKey} | 获取文件列表
*OAuthAPI* | [**OAuthAuthorize**](docs/OAuthAPI.md#oauthauthorize) | **Post** /OAuth/{appKey}/Authorize | 获取access_token
*OAuthAPI* | [**OAuthConsents**](docs/OAuthAPI.md#oauthconsents) | **Get** /OAuth/{appKey}/Consents | 授权记录
*OAuthAPI* | [**OAuthDeleteConsent**](docs/OAuthAPI.md#oauthdeleteconsent) | **Delete** /OAuth/{appKey}/Consents/{id} | 删除授权记录
*OAuthAPI* | [**OAuthGrantCode**](docs/OAuthAPI.md#oauthgrantcode) | **Post** /OAuth/{appKey}/GrantCode | 申请授权码
*OAuthAPI* | [**OAuthProfile**](docs/OAuthAPI.md#oauthprofile) | **Get** /OAuth/{appKey}/Profile | 获取个人资料
*OrderAPI* | [**Order**](docs/OrderAPI.md#order) | **Get** /Order/{appKey}/{id} | 获取订单详情
*OrderAPI* | [**OrderCreate**](docs/OrderAPI.md#ordercreate) | **Post** /Order/{appKey}/Create | 创建订单
*OrderAPI* | [**Orders**](docs/OrderAPI.md#orders) | **Get** /Order/{appKey} | 获取订单列表
*ProjectAPI* | [**Project**](docs/ProjectAPI.md#project) | **Get** /Project/{id} | 项目详情
*ProjectAPI* | [**ProjectDelete**](docs/ProjectAPI.md#projectdelete) | **Delete** /Project/{id} | 删除项目
*ProjectAPI* | [**ProjectPost**](docs/ProjectAPI.md#projectpost) | **Post** /Project | 创建项目
*ProjectAPI* | [**ProjectPut**](docs/ProjectAPI.md#projectput) | **Put** /Project/{id} | 更新项目
*ProjectAPI* | [**Projects**](docs/ProjectAPI.md#projects) | **Get** /Project | 项目列表
*ServiceSettingAPI* | [**ServiceSetting**](docs/ServiceSettingAPI.md#servicesetting) | **Get** /ServiceSetting/{id} | 获取配置详情
*ServiceSettingAPI* | [**ServiceSettingDelete**](docs/ServiceSettingAPI.md#servicesettingdelete) | **Delete** /ServiceSetting/{id} | 删除配置
*ServiceSettingAPI* | [**ServiceSettingGroup**](docs/ServiceSettingAPI.md#servicesettinggroup) | **Get** /ServiceSetting/Groups/{id} | 获取服务分组详情
*ServiceSettingAPI* | [**ServiceSettingGroupDelete**](docs/ServiceSettingAPI.md#servicesettinggroupdelete) | **Delete** /ServiceSetting/Groups/{id} | 删除服务分组
*ServiceSettingAPI* | [**ServiceSettingGroupPost**](docs/ServiceSettingAPI.md#servicesettinggrouppost) | **Post** /ServiceSetting/Groups | 添加服务分组
*ServiceSettingAPI* | [**ServiceSettingGroupPut**](docs/ServiceSettingAPI.md#servicesettinggroupput) | **Put** /ServiceSetting/Groups/{id} | 更新服务分组
*ServiceSettingAPI* | [**ServiceSettingGroups**](docs/ServiceSettingAPI.md#servicesettinggroups) | **Get** /ServiceSetting/Groups | 获取服务分组列表
*ServiceSettingAPI* | [**ServiceSettingItem**](docs/ServiceSettingAPI.md#servicesettingitem) | **Get** /ServiceSetting/Items/{id} | 服务配置详情
*ServiceSettingAPI* | [**ServiceSettingItemDelete**](docs/ServiceSettingAPI.md#servicesettingitemdelete) | **Delete** /ServiceSetting/Items/{id} | 删除服务配置
*ServiceSettingAPI* | [**ServiceSettingItemPost**](docs/ServiceSettingAPI.md#servicesettingitempost) | **Post** /ServiceSetting/Items | 添加服务配置
*ServiceSettingAPI* | [**ServiceSettingItemPut**](docs/ServiceSettingAPI.md#servicesettingitemput) | **Put** /ServiceSetting/Items/{id} | 更新服务配置
*ServiceSettingAPI* | [**ServiceSettingItems**](docs/ServiceSettingAPI.md#servicesettingitems) | **Get** /ServiceSetting/Items | 服务配置列表
*ServiceSettingAPI* | [**ServiceSettingPost**](docs/ServiceSettingAPI.md#servicesettingpost) | **Post** /ServiceSetting | 增加配置
*ServiceSettingAPI* | [**ServiceSettingProvider**](docs/ServiceSettingAPI.md#servicesettingprovider) | **Get** /ServiceSetting/Providers/{id} | 获取服务商详情
*ServiceSettingAPI* | [**ServiceSettingProviderDelete**](docs/ServiceSettingAPI.md#servicesettingproviderdelete) | **Delete** /ServiceSetting/Providers/{id} | 删除服务商
*ServiceSettingAPI* | [**ServiceSettingProviderPost**](docs/ServiceSettingAPI.md#servicesettingproviderpost) | **Post** /ServiceSetting/Providers | 添加服务商
*ServiceSettingAPI* | [**ServiceSettingProviderPut**](docs/ServiceSettingAPI.md#servicesettingproviderput) | **Put** /ServiceSetting/Providers/{id} | 更新服务商
*ServiceSettingAPI* | [**ServiceSettingProviders**](docs/ServiceSettingAPI.md#servicesettingproviders) | **Get** /ServiceSetting/Providers | 获取服务商列表
*ServiceSettingAPI* | [**ServiceSettingPut**](docs/ServiceSettingAPI.md#servicesettingput) | **Put** /ServiceSetting/{id} | 更新配置
*ServiceSettingAPI* | [**ServiceSettings**](docs/ServiceSettingAPI.md#servicesettings) | **Get** /ServiceSetting | 获取配置列表
*StorageAPI* | [**StorageAggregate**](docs/StorageAPI.md#storageaggregate) | **Get** /Storage/{appKey}/{table}/Aggregate | 聚合查询
*StorageAPI* | [**StorageClear**](docs/StorageAPI.md#storageclear) | **Delete** /Storage/{appKey}/{table}/Clear | 清空表数据
*StorageAPI* | [**StorageDelete**](docs/StorageAPI.md#storagedelete) | **Delete** /Storage/{appKey}/{table}/{id} | 删除数据
*StorageAPI* | [**StorageDeleteIndex**](docs/StorageAPI.md#storagedeleteindex) | **Delete** /Storage/{appKey}/{table}/indexes | 删除索引
*StorageAPI* | [**StorageDetail**](docs/StorageAPI.md#storagedetail) | **Get** /Storage/{appKey}/{table}/{id} | 数据详情
*StorageAPI* | [**StorageExecuteFunction**](docs/StorageAPI.md#storageexecutefunction) | **Get** /Storage/{appKey}/ExecuteFunction | 执行函数
*StorageAPI* | [**StorageExport**](docs/StorageAPI.md#storageexport) | **Get** /Storage/{appKey}/{table}/Export | 导出数据
*StorageAPI* | [**StorageImport**](docs/StorageAPI.md#storageimport) | **Post** /Storage/{appKey}/{table}/Import | 导入数据
*StorageAPI* | [**StorageIndexes**](docs/StorageAPI.md#storageindexes) | **Get** /Storage/{appKey}/{table}/Indexes | 获取索引
*StorageAPI* | [**StorageList**](docs/StorageAPI.md#storagelist) | **Get** /Storage/{appKey}/{table} | 查询数据
*StorageAPI* | [**StoragePost**](docs/StorageAPI.md#storagepost) | **Post** /Storage/{appKey}/{table} | 添加数据
*StorageAPI* | [**StoragePostIndex**](docs/StorageAPI.md#storagepostindex) | **Post** /Storage/{appKey}/{table}/indexes | 添加索引
*StorageAPI* | [**StoragePut**](docs/StorageAPI.md#storageput) | **Put** /Storage/{appKey}/{table}/{id} | 更新数据
*StorageAPI* | [**StorageRemove**](docs/StorageAPI.md#storageremove) | **Delete** /Storage/{appKey}/{table}/Remove | 删除表
*StorageAPI* | [**StorageStats**](docs/StorageAPI.md#storagestats) | **Get** /Storage/{appKey}/{table}/Stats | 数据表统计
*StorageAPI* | [**StorageTables**](docs/StorageAPI.md#storagetables) | **Get** /Storage/{appKey}/Tables | 获取数据表
*SystemFileAPI* | [**SystemFileCreateFolder**](docs/SystemFileAPI.md#systemfilecreatefolder) | **Post** /SystemFile/CreateFolder | 创建文件夹
*SystemFileAPI* | [**SystemFileDelete**](docs/SystemFileAPI.md#systemfiledelete) | **Delete** /SystemFile | 删除文件
*SystemFileAPI* | [**SystemFileRename**](docs/SystemFileAPI.md#systemfilerename) | **Post** /SystemFile/Rename | 重命名文件
*SystemFileAPI* | [**SystemFileUpload**](docs/SystemFileAPI.md#systemfileupload) | **Post** /SystemFile | 上传文件
*SystemFileAPI* | [**SystemFiles**](docs/SystemFileAPI.md#systemfiles) | **Get** /SystemFile | 获取文件列表
*TeamAPI* | [**TeamDelete**](docs/TeamAPI.md#teamdelete) | **Delete** /Team/{id} | 删除团队
*TeamAPI* | [**TeamPost**](docs/TeamAPI.md#teampost) | **Post** /Team | 创建团队
*TeamAPI* | [**TeamPut**](docs/TeamAPI.md#teamput) | **Put** /Team/{id} | 更新团队信息
*TeamAPI* | [**Teams**](docs/TeamAPI.md#teams) | **Get** /Team | 获取团队列表
*UserAPI* | [**User**](docs/UserAPI.md#user) | **Get** /User/{appKey}/{id} | 获取用户详情
*UserAPI* | [**UserClear**](docs/UserAPI.md#userclear) | **Delete** /User/{appKey}/Clear | 清空用户数据
*UserAPI* | [**UserCommonInterests**](docs/UserAPI.md#usercommoninterests) | **Get** /User/{appKey}/Friends/CommonInterests | 有共同爱好的用户推荐
*UserAPI* | [**UserDelete**](docs/UserAPI.md#userdelete) | **Delete** /User/{appKey}/{id} | 删除用户
*UserAPI* | [**UserEmailSignIn**](docs/UserAPI.md#useremailsignin) | **Post** /User/{appKey}/EmailSignIn | 邮箱登录
*UserAPI* | [**UserEmailSignUp**](docs/UserAPI.md#useremailsignup) | **Post** /User/{appKey}/EmailSignUp | 邮箱注册
*UserAPI* | [**UserExport**](docs/UserAPI.md#userexport) | **Get** /User/{appKey}/Export | 导出用户数据
*UserAPI* | [**UserFollowUser**](docs/UserAPI.md#userfollowuser) | **Post** /User/{appKey}/Follower/{userId} | 关注用户
*UserAPI* | [**UserFollowerPut**](docs/UserAPI.md#userfollowerput) | **Put** /User/{appKey}/Follower/{id} | 更新粉丝
*UserAPI* | [**UserFollowers**](docs/UserAPI.md#userfollowers) | **Get** /User/{appKey}/Followers | 获取我的粉丝列表
*UserAPI* | [**UserFollowing**](docs/UserAPI.md#userfollowing) | **Get** /User/{appKey}/Following | 获取我的关注列表
*UserAPI* | [**UserFriendsNearBy**](docs/UserAPI.md#userfriendsnearby) | **Get** /User/{appKey}/Friends/NearBy | 附近的用户推荐
*UserAPI* | [**UserImport**](docs/UserAPI.md#userimport) | **Post** /User/{appKey}/Import | 导入用户数据
*UserAPI* | [**UserLocation**](docs/UserAPI.md#userlocation) | **Get** /User/{appKey}/Location/{id} | 获取位置详情
*UserAPI* | [**UserLocationDelete**](docs/UserAPI.md#userlocationdelete) | **Delete** /User/{appKey}/Location/{id} | 删除位置
*UserAPI* | [**UserLocationPost**](docs/UserAPI.md#userlocationpost) | **Post** /User/{appKey}/Location | 添加位置
*UserAPI* | [**UserLocationPut**](docs/UserAPI.md#userlocationput) | **Put** /User/{appKey}/Location/{id} | 更新位置
*UserAPI* | [**UserLocations**](docs/UserAPI.md#userlocations) | **Get** /User/{appKey}/Locations | 获取位置列表
*UserAPI* | [**UserMutualFollowers**](docs/UserAPI.md#usermutualfollowers) | **Get** /User/{appKey}/Friends/MutualFollowers | 有共同粉丝的用户推荐
*UserAPI* | [**UserMutualFollowings**](docs/UserAPI.md#usermutualfollowings) | **Get** /User/{appKey}/Friends/MutualFollowings | 有共同关注的用户推荐
*UserAPI* | [**UserOAuthAccountBind**](docs/UserAPI.md#useroauthaccountbind) | **Post** /User/{appKey}/OAuthAccountBind | 外部账号 绑定
*UserAPI* | [**UserOAuthAccountSignIn**](docs/UserAPI.md#useroauthaccountsignin) | **Post** /User/{appKey}/OAuthAccountSignIn | 外部账号 登录
*UserAPI* | [**UserOAuthAccounts**](docs/UserAPI.md#useroauthaccounts) | **Get** /User/{appKey}/OAuthAccounts | 外部账号 绑定列表
*UserAPI* | [**UserOAuthAccountsPutBind**](docs/UserAPI.md#useroauthaccountsputbind) | **Put** /User/{appKey}/OAuthAccounts/{id} | 外部账号 更新绑定
*UserAPI* | [**UserOAuthAccountsUnBind**](docs/UserAPI.md#useroauthaccountsunbind) | **Delete** /User/{appKey}/OAuthAccounts/{id} | 外部账号 删除绑定
*UserAPI* | [**UserPhoneSignIn**](docs/UserAPI.md#userphonesignin) | **Post** /User/{appKey}/PhoneSignIn | 手机登录
*UserAPI* | [**UserPhoneSignUp**](docs/UserAPI.md#userphonesignup) | **Post** /User/{appKey}/PhoneSignUp | 手机注册
*UserAPI* | [**UserProfile**](docs/UserAPI.md#userprofile) | **Get** /User/{appKey}/Profile | 获取个人资料
*UserAPI* | [**UserPut**](docs/UserAPI.md#userput) | **Put** /User/{appKey}/{id} | 更新用户信息
*UserAPI* | [**UserQRCodePreSignIn**](docs/UserAPI.md#userqrcodepresignin) | **Post** /User/{appKey}/QRCodePreSignIn | 微信小程序 - 预登陆
*UserAPI* | [**UserQRCodeScan**](docs/UserAPI.md#userqrcodescan) | **Post** /User/{appKey}/QRCodeScan | 微信小程序 - 扫码
*UserAPI* | [**UserQRCodeSignIn**](docs/UserAPI.md#userqrcodesignin) | **Post** /User/{appKey}/QRCodeSignIn | 微信小程序 - 网页登录
*UserAPI* | [**UserQRCodeSignUp**](docs/UserAPI.md#userqrcodesignup) | **Post** /User/{appKey}/QRCodeSignUp | 微信小程序 - 注册
*UserAPI* | [**UserResetPwd**](docs/UserAPI.md#userresetpwd) | **Post** /User/{appKey}/ResetPwd | 重置密码
*UserAPI* | [**UserSendEmailCode**](docs/UserAPI.md#usersendemailcode) | **Post** /User/{appKey}/SendEmailCode | 发送邮箱验证码
*UserAPI* | [**UserSendSMSCode**](docs/UserAPI.md#usersendsmscode) | **Post** /User/{appKey}/SendSMSCode | 发送手机验证码
*UserAPI* | [**UserSignIn**](docs/UserAPI.md#usersignin) | **Post** /User/{appKey}/SignIn | 账号密码 登录
*UserAPI* | [**UserSignUp**](docs/UserAPI.md#usersignup) | **Post** /User/{appKey}/SignUp | 账号密码 注册
*UserAPI* | [**UserTwoFactorAuth**](docs/UserAPI.md#usertwofactorauth) | **Get** /User/{appKey}/TwoFactorAuth | 双因素验证
*UserAPI* | [**UserUnfollowUser**](docs/UserAPI.md#userunfollowuser) | **Delete** /User/{appKey}/Follower/{userId} | 取消关注
*UserAPI* | [**UserUnionIDSignIn**](docs/UserAPI.md#userunionidsignin) | **Post** /User/{appKey}/UnionIDSignIn | UnionID登录
*UserAPI* | [**UserUnionIDSignUp**](docs/UserAPI.md#userunionidsignup) | **Post** /User/{appKey}/UnionIDSignUp | UnionID注册
*UserAPI* | [**UserUpdateProfile**](docs/UserAPI.md#userupdateprofile) | **Put** /User/{appKey}/Profile | 更新个人资料
*UserAPI* | [**Users**](docs/UserAPI.md#users) | **Get** /User/{appKey} | 获取用户列表
*UserCurrencyAPI* | [**UserCurrencies**](docs/UserCurrencyAPI.md#usercurrencies) | **Get** /UserCurrency/{appKey}/{id} | 获取用户资产
*UserCurrencyAPI* | [**UserCurrencyConsume**](docs/UserCurrencyAPI.md#usercurrencyconsume) | **Post** /UserCurrency/{appKey}/CurrencyConsume | 消费虚拟币
*UserCurrencyAPI* | [**UserCurrencyExchange**](docs/UserCurrencyAPI.md#usercurrencyexchange) | **Post** /UserCurrency/{appKey}/CurrencyExchange | 兑换虚拟币
*UserCurrencyAPI* | [**UserCurrencyRecharge**](docs/UserCurrencyAPI.md#usercurrencyrecharge) | **Post** /UserCurrency/{appKey}/CurrencyRecharge | 充值虚拟币
*UserCurrencyAPI* | [**UserCurrencyTransactions**](docs/UserCurrencyAPI.md#usercurrencytransactions) | **Get** /UserCurrency/{appKey}/CurrencyTransactions | 虚拟币交易记录
*UserSettingAPI* | [**UserSetting**](docs/UserSettingAPI.md#usersetting) | **Get** /UserSetting/{appKey}/{id} | 获取用户配置项详情
*UserSettingAPI* | [**UserSettingDelete**](docs/UserSettingAPI.md#usersettingdelete) | **Delete** /UserSetting/{appKey}/{id} | 删除用户配置项
*UserSettingAPI* | [**UserSettingPost**](docs/UserSettingAPI.md#usersettingpost) | **Post** /UserSetting/{appKey} | 添加用户配置项
*UserSettingAPI* | [**UserSettingPut**](docs/UserSettingAPI.md#usersettingput) | **Put** /UserSetting/{appKey}/{id} | 更新用户配置项
*UserSettingAPI* | [**UserSettings**](docs/UserSettingAPI.md#usersettings) | **Get** /UserSetting/{appKey} | 获取用户配置列表
*WechatAPI* | [**WechatDecrypt**](docs/WechatAPI.md#wechatdecrypt) | **Get** /Wechat/{appKey}/Decrypt | 小程序-解密数据
*WechatAPI* | [**WechatGenerateScheme**](docs/WechatAPI.md#wechatgeneratescheme) | **Post** /Wechat/{appKey}/GenerateScheme | 小程序-生成scheme码(该接口用于获取小程序 scheme 码，适用于短信、邮件、外部网页、微信内等拉起小程序的业务场景)
*WechatAPI* | [**WechatJSCode2Session**](docs/WechatAPI.md#wechatjscode2session) | **Get** /Wechat/{appKey}/JSCode2Session | 小程序-登录凭证校验
*WechatAPI* | [**WechatJSConfig**](docs/WechatAPI.md#wechatjsconfig) | **Get** /Wechat/{appKey}/JSConfig | 公众号H5-JS SDK Config
*WechatAPI* | [**WechatSubscribeMSG**](docs/WechatAPI.md#wechatsubscribemsg) | **Post** /Wechat/{appKey}/SubscribeMSG | 公众号H5-发送一次性订阅消息
*WechatAPI* | [**WechatSubscribeSend**](docs/WechatAPI.md#wechatsubscribesend) | **Post** /Wechat/{appKey}/SubscribeSend | 小程序-发送订阅消息
*WechatAPI* | [**WechatUrlLinkGenerate**](docs/WechatAPI.md#wechaturllinkgenerate) | **Post** /Wechat/{appKey}/UrlLinkGenerate | 小程序-生成网页跳转地址(获取小程序 URL Link，适用于短信、邮件、网页、微信内等拉起小程序的业务场景)
*WechatAPI* | [**WechatUserInfo**](docs/WechatAPI.md#wechatuserinfo) | **Get** /Wechat/{appKey}/UserInfo | 公众号H5-获取用户UnionID
*WechatAPI* | [**WechatWXACodeGet**](docs/WechatAPI.md#wechatwxacodeget) | **Post** /Wechat/{appKey}/WXACodeGet | 小程序-获取小程序码
*WechatAPI* | [**WechatWXACodeGetUnlimited**](docs/WechatAPI.md#wechatwxacodegetunlimited) | **Post** /Wechat/{appKey}/WXACodeGetUnlimited | 小程序-获取小程序码(无限制)



## 模型文档

 - [AccessTokenListResult](docs/AccessTokenListResult.md)
 - [AccessTokenListResultApiResponse](docs/AccessTokenListResultApiResponse.md)
 - [AccessTokenPostRequest](docs/AccessTokenPostRequest.md)
 - [AccessTokenPutRequest](docs/AccessTokenPutRequest.md)
 - [AlipayCreateOrderPagePayRequest](docs/AlipayCreateOrderPagePayRequest.md)
 - [AlipayCreateOrderRequest](docs/AlipayCreateOrderRequest.md)
 - [AlipayCreateOrderWapPayRequest](docs/AlipayCreateOrderWapPayRequest.md)
 - [AlipayTradeQueryResponse](docs/AlipayTradeQueryResponse.md)
 - [AlipayTradeQueryResponseApiResponse](docs/AlipayTradeQueryResponseApiResponse.md)
 - [AlipayTradeRefundResponse](docs/AlipayTradeRefundResponse.md)
 - [AlipayTradeRefundResponseApiResponse](docs/AlipayTradeRefundResponseApiResponse.md)
 - [App](docs/App.md)
 - [AppApiResponse](docs/AppApiResponse.md)
 - [AppCheckVersionResult](docs/AppCheckVersionResult.md)
 - [AppCheckVersionResultApiResponse](docs/AppCheckVersionResultApiResponse.md)
 - [AppInfoItem](docs/AppInfoItem.md)
 - [AppInfoResult](docs/AppInfoResult.md)
 - [AppInfoResultApiResponse](docs/AppInfoResultApiResponse.md)
 - [AppListResult](docs/AppListResult.md)
 - [AppListResultApiResponse](docs/AppListResultApiResponse.md)
 - [AppPostResult](docs/AppPostResult.md)
 - [AppPostResultApiResponse](docs/AppPostResultApiResponse.md)
 - [AppProperty](docs/AppProperty.md)
 - [AppSetting](docs/AppSetting.md)
 - [AppSettingApiResponse](docs/AppSettingApiResponse.md)
 - [AppSettingGroupPostResult](docs/AppSettingGroupPostResult.md)
 - [AppSettingGroupPostResultApiResponse](docs/AppSettingGroupPostResultApiResponse.md)
 - [AppSettingItemPostResult](docs/AppSettingItemPostResult.md)
 - [AppSettingItemPostResultApiResponse](docs/AppSettingItemPostResultApiResponse.md)
 - [AppSettingListApiResponse](docs/AppSettingListApiResponse.md)
 - [AppSettingProviderPostResult](docs/AppSettingProviderPostResult.md)
 - [AppSettingProviderPostResultApiResponse](docs/AppSettingProviderPostResultApiResponse.md)
 - [AppSettingSettingPostResult](docs/AppSettingSettingPostResult.md)
 - [AppSettingSettingPostResultApiResponse](docs/AppSettingSettingPostResultApiResponse.md)
 - [AppUserConsentModel](docs/AppUserConsentModel.md)
 - [AppUserConsentModelListApiResponse](docs/AppUserConsentModelListApiResponse.md)
 - [AppUserListResponse](docs/AppUserListResponse.md)
 - [AppUserResetPwdRequest](docs/AppUserResetPwdRequest.md)
 - [AuthorizePolicy](docs/AuthorizePolicy.md)
 - [AuthorizePolicyApiResponse](docs/AuthorizePolicyApiResponse.md)
 - [AuthorizePolicyListApiResponse](docs/AuthorizePolicyListApiResponse.md)
 - [AuthorizeResult](docs/AuthorizeResult.md)
 - [AuthorizeResultApiResponse](docs/AuthorizeResultApiResponse.md)
 - [BkAgentRespInfo](docs/BkAgentRespInfo.md)
 - [BooleanApiResponse](docs/BooleanApiResponse.md)
 - [ChargeInfo](docs/ChargeInfo.md)
 - [CommonFriendModel](docs/CommonFriendModel.md)
 - [ContributeDetail](docs/ContributeDetail.md)
 - [CreateOrderRequest](docs/CreateOrderRequest.md)
 - [CreateOrderResult](docs/CreateOrderResult.md)
 - [CreateOrderResultApiResponse](docs/CreateOrderResultApiResponse.md)
 - [CreatePostResult](docs/CreatePostResult.md)
 - [CreatePostResultApiResponse](docs/CreatePostResultApiResponse.md)
 - [Currency](docs/Currency.md)
 - [CurrencyApiResponse](docs/CurrencyApiResponse.md)
 - [CurrencyConsumeRequest](docs/CurrencyConsumeRequest.md)
 - [CurrencyExchangeRate](docs/CurrencyExchangeRate.md)
 - [CurrencyExchangeRateApiResponse](docs/CurrencyExchangeRateApiResponse.md)
 - [CurrencyListApiResponse](docs/CurrencyListApiResponse.md)
 - [CurrencyTransaction](docs/CurrencyTransaction.md)
 - [CurrencyTransactionListApiResponse](docs/CurrencyTransactionListApiResponse.md)
 - [DirectoryItem](docs/DirectoryItem.md)
 - [EmailSignInRequest](docs/EmailSignInRequest.md)
 - [EmailSignUpRequest](docs/EmailSignUpRequest.md)
 - [EnterprisePayInfo](docs/EnterprisePayInfo.md)
 - [ExchangeCurrencyRequest](docs/ExchangeCurrencyRequest.md)
 - [ExchangeRatePutRequest](docs/ExchangeRatePutRequest.md)
 - [ExecuteFunctionRequest](docs/ExecuteFunctionRequest.md)
 - [FileItem](docs/FileItem.md)
 - [FileListResult](docs/FileListResult.md)
 - [FileListResultApiResponse](docs/FileListResultApiResponse.md)
 - [FollowerModel](docs/FollowerModel.md)
 - [FollowerPutModel](docs/FollowerPutModel.md)
 - [FulfillmentDetail](docs/FulfillmentDetail.md)
 - [GeoLocationModel](docs/GeoLocationModel.md)
 - [GeoLocationModelApiResponse](docs/GeoLocationModelApiResponse.md)
 - [GeoLocationResponseModel](docs/GeoLocationResponseModel.md)
 - [GoodsDetail](docs/GoodsDetail.md)
 - [GrantRequest](docs/GrantRequest.md)
 - [GrantResult](docs/GrantResult.md)
 - [GrantResultApiResponse](docs/GrantResultApiResponse.md)
 - [HbFqPayInfo](docs/HbFqPayInfo.md)
 - [Int64ApiResponse](docs/Int64ApiResponse.md)
 - [IntactChargeInfo](docs/IntactChargeInfo.md)
 - [ListResponseItem](docs/ListResponseItem.md)
 - [ListResponseItemListApiResponse](docs/ListResponseItemListApiResponse.md)
 - [OAuthAccountBindRequest](docs/OAuthAccountBindRequest.md)
 - [OAuthAccountPutBindRequest](docs/OAuthAccountPutBindRequest.md)
 - [OAuthAccountSignInRequest](docs/OAuthAccountSignInRequest.md)
 - [ObjectApiResponse](docs/ObjectApiResponse.md)
 - [ObjectListApiResponse](docs/ObjectListApiResponse.md)
 - [Order](docs/Order.md)
 - [OrderApiResponse](docs/OrderApiResponse.md)
 - [OrderListResult](docs/OrderListResult.md)
 - [OrderListResultApiResponse](docs/OrderListResultApiResponse.md)
 - [PaymentInfoWithId](docs/PaymentInfoWithId.md)
 - [PhoneSignInRequest](docs/PhoneSignInRequest.md)
 - [PhoneSignUpRequest](docs/PhoneSignUpRequest.md)
 - [PostIndexRequest](docs/PostIndexRequest.md)
 - [PostResult](docs/PostResult.md)
 - [PostResultApiResponse](docs/PostResultApiResponse.md)
 - [PresetPayToolInfo](docs/PresetPayToolInfo.md)
 - [ProfileResult](docs/ProfileResult.md)
 - [ProfileResultApiResponse](docs/ProfileResultApiResponse.md)
 - [Project](docs/Project.md)
 - [ProjectApiResponse](docs/ProjectApiResponse.md)
 - [ProjectListResult](docs/ProjectListResult.md)
 - [ProjectListResultApiResponse](docs/ProjectListResultApiResponse.md)
 - [QRCodePreSignInRequest](docs/QRCodePreSignInRequest.md)
 - [QRCodeScanRequest](docs/QRCodeScanRequest.md)
 - [QRCodeSignInRequest](docs/QRCodeSignInRequest.md)
 - [QRCodeSignUpRequest](docs/QRCodeSignUpRequest.md)
 - [RechargePointRequest](docs/RechargePointRequest.md)
 - [RecommendFriend](docs/RecommendFriend.md)
 - [RefundChargeInfo](docs/RefundChargeInfo.md)
 - [RefundSubFee](docs/RefundSubFee.md)
 - [ReturnPageNotifyRequest](docs/ReturnPageNotifyRequest.md)
 - [SendEmailCodeRequest](docs/SendEmailCodeRequest.md)
 - [SendSMSCodeRequest](docs/SendSMSCodeRequest.md)
 - [ServiceGroup](docs/ServiceGroup.md)
 - [ServiceGroupApiResponse](docs/ServiceGroupApiResponse.md)
 - [ServiceGroupListApiResponse](docs/ServiceGroupListApiResponse.md)
 - [ServiceItem](docs/ServiceItem.md)
 - [ServiceItemApiResponse](docs/ServiceItemApiResponse.md)
 - [ServiceItemListApiResponse](docs/ServiceItemListApiResponse.md)
 - [ServiceProvider](docs/ServiceProvider.md)
 - [ServiceProviderApiResponse](docs/ServiceProviderApiResponse.md)
 - [ServiceProviderListApiResponse](docs/ServiceProviderListApiResponse.md)
 - [ServiceSettingGroupPostResult](docs/ServiceSettingGroupPostResult.md)
 - [ServiceSettingGroupPostResultApiResponse](docs/ServiceSettingGroupPostResultApiResponse.md)
 - [ServiceSettingItemPostResult](docs/ServiceSettingItemPostResult.md)
 - [ServiceSettingItemPostResultApiResponse](docs/ServiceSettingItemPostResultApiResponse.md)
 - [ServiceSettingProviderPostResult](docs/ServiceSettingProviderPostResult.md)
 - [ServiceSettingProviderPostResultApiResponse](docs/ServiceSettingProviderPostResultApiResponse.md)
 - [ServiceSettingSettingPostResult](docs/ServiceSettingSettingPostResult.md)
 - [ServiceSettingSettingPostResultApiResponse](docs/ServiceSettingSettingPostResultApiResponse.md)
 - [Settings](docs/Settings.md)
 - [SettingsApiResponse](docs/SettingsApiResponse.md)
 - [SettingsListApiResponse](docs/SettingsListApiResponse.md)
 - [SetupCode](docs/SetupCode.md)
 - [SetupCodeApiResponse](docs/SetupCodeApiResponse.md)
 - [SignInRequest](docs/SignInRequest.md)
 - [SignUpRequest](docs/SignUpRequest.md)
 - [StorageListResult](docs/StorageListResult.md)
 - [StorageListResultApiResponse](docs/StorageListResultApiResponse.md)
 - [StringApiResponse](docs/StringApiResponse.md)
 - [StringListApiResponse](docs/StringListApiResponse.md)
 - [SubFee](docs/SubFee.md)
 - [SystemDirectoryItem](docs/SystemDirectoryItem.md)
 - [SystemFileItem](docs/SystemFileItem.md)
 - [SystemFileListResult](docs/SystemFileListResult.md)
 - [SystemFileListResultApiResponse](docs/SystemFileListResultApiResponse.md)
 - [TapPayInfo](docs/TapPayInfo.md)
 - [Team](docs/Team.md)
 - [TokenModel](docs/TokenModel.md)
 - [TokenModelApiResponse](docs/TokenModelApiResponse.md)
 - [TradeFundBill](docs/TradeFundBill.md)
 - [TradeSettleDetail](docs/TradeSettleDetail.md)
 - [TradeSettleInfo](docs/TradeSettleInfo.md)
 - [UnionIDSignInRequest](docs/UnionIDSignInRequest.md)
 - [UnionIDSignUpRequest](docs/UnionIDSignUpRequest.md)
 - [UpdateProfileRequest](docs/UpdateProfileRequest.md)
 - [User](docs/User.md)
 - [UserAccessToken](docs/UserAccessToken.md)
 - [UserApiResponse](docs/UserApiResponse.md)
 - [UserCommonInterestsResult](docs/UserCommonInterestsResult.md)
 - [UserCommonInterestsResultApiResponse](docs/UserCommonInterestsResultApiResponse.md)
 - [UserCurrency](docs/UserCurrency.md)
 - [UserCurrencyCurrencyTransResult](docs/UserCurrencyCurrencyTransResult.md)
 - [UserCurrencyCurrencyTransResultApiResponse](docs/UserCurrencyCurrencyTransResultApiResponse.md)
 - [UserCurrencyListApiResponse](docs/UserCurrencyListApiResponse.md)
 - [UserFollowersResult](docs/UserFollowersResult.md)
 - [UserFollowersResultApiResponse](docs/UserFollowersResultApiResponse.md)
 - [UserFollowingResult](docs/UserFollowingResult.md)
 - [UserFollowingResultApiResponse](docs/UserFollowingResultApiResponse.md)
 - [UserFriendsNearByResult](docs/UserFriendsNearByResult.md)
 - [UserFriendsNearByResultApiResponse](docs/UserFriendsNearByResultApiResponse.md)
 - [UserListResult](docs/UserListResult.md)
 - [UserListResultApiResponse](docs/UserListResultApiResponse.md)
 - [UserLocationPostResult](docs/UserLocationPostResult.md)
 - [UserLocationPostResultApiResponse](docs/UserLocationPostResultApiResponse.md)
 - [UserLocationsResult](docs/UserLocationsResult.md)
 - [UserLocationsResultApiResponse](docs/UserLocationsResultApiResponse.md)
 - [UserLogins](docs/UserLogins.md)
 - [UserLoginsListApiResponse](docs/UserLoginsListApiResponse.md)
 - [UserMutualFollowersResult](docs/UserMutualFollowersResult.md)
 - [UserMutualFollowersResultApiResponse](docs/UserMutualFollowersResultApiResponse.md)
 - [UserMutualFollowingsResult](docs/UserMutualFollowingsResult.md)
 - [UserMutualFollowingsResultApiResponse](docs/UserMutualFollowingsResultApiResponse.md)
 - [UserProfileResult](docs/UserProfileResult.md)
 - [UserProfileResultApiResponse](docs/UserProfileResultApiResponse.md)
 - [UserQRCodeScanResult](docs/UserQRCodeScanResult.md)
 - [UserQRCodeScanResultApiResponse](docs/UserQRCodeScanResultApiResponse.md)
 - [UserSetting](docs/UserSetting.md)
 - [UserSettingApiResponse](docs/UserSettingApiResponse.md)
 - [UserSettingListApiResponse](docs/UserSettingListApiResponse.md)
 - [UserSettingPostResult](docs/UserSettingPostResult.md)
 - [UserSettingPostResultApiResponse](docs/UserSettingPostResultApiResponse.md)
 - [VoucherDetail](docs/VoucherDetail.md)
 - [WechatJSConfigResult](docs/WechatJSConfigResult.md)
 - [WechatJSConfigResultApiResponse](docs/WechatJSConfigResultApiResponse.md)


## 授权文档


为API定义的认证方案：
### Bearer

- **类型**：HTTP Bearer令牌认证

示例

```go
auth := context.WithValue(context.Background(), zsgf.client.ContextAccessToken, "BEARER_TOKEN_STRING")
r, err := client.Service.Operation(auth, args)
```


## 实用方法文档

由于模型结构成员都是指针，此包包含许多实用函数，可以轻松获取基本类型值的指针。
这些函数中的每一个都接受给定基本类型的值并返回指向它的指针：

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`