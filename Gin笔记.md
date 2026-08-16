# Gin 入门通俗笔记
## 1. 最通俗原理（外卖打电话类比）
整套访问流程：
浏览器(你本人) → 拨通电话发起请求 → Gin后端(饭店老板)接收 → 返回文字消息 → 浏览器页面展示消息

1. 浏览器 = 你
2. Gin程序 = 守在店里待命的饭店老板
3. 网址+端口 `127.0.0.1:8080` = 老板的电话号码
4. 访问不同后缀 `/`、`/news` = 拨打不同分机

## 2. 代码逐行拆解
```
package main
// 代表这是可以直接运行的主程序

import "github.com/gin-gonic/gin"
// 加载Gin网页工具，用来搭建网页后端

func main() {
	// r：创建服务器（开店，自带日志监控、崩溃保护）
	r := gin.Default()

	// 路由1：根路径  浏览器访问 127.0.0.1:8080
	r.GET("/", func(c *gin.Context) {
		// c=服务员，负责接收请求、返回文字；200代表访问成功
		c.String(200, "值:%v", "你好gin")
	})

	// 路由2：新闻页面 浏览器访问 127.0.0.1:8080/news
	r.GET("/news", func(c *gin.Context) {
		c.String(200, "我是新闻页面")
	})

	// r.Run() 开启店铺，8080端口开始等待浏览器访问
	r.Run()
}
```
|专业名词	|通俗理解|
|-------|--------|
|GET	|普通点开网页、敲门拜访，日常打开网页都是 GET 请求|
|路由	|不同网址，相当于店铺多个办事窗口|
|gin.Context（c）|	浏览器和服务器中间传话的服务员|
|c.String (200,"文字")	|服务员递给浏览器一段文字，200 = 请求成功暗号|
|127.0.0.1	|特指你自己本机电脑，仅本机能够访问|
|端口 8080	|服务器的门牌号 / 电话号码|
# Gin 框架全套学习笔记
## 一、前置必备基础
### 1. 需要掌握的Go基础语法
- 变量、基础数据类型
- 函数、函数传参与返回值
- 切片、map（后端高频容器）
- 结构体：封装用户、新闻等实体数据
- 指针：理解内存地址传递
- 匿名函数：Gin路由回调本质就是匿名函数

### 2. 闭包
> 函数内部嵌套内层函数，内层函数可以读取外层函数的变量
- Gin中 r.GET() 属于外层函数
- 回调 func(c *gin.Context){} 为内层匿名函数
- 内层函数可以拿到外层传递过来的上下文 c

### 3. HTTP 五种请求方式（通俗口诀：GET查、POST新增、PUT全改、PATCH小改、DELETE删除）
1. **GET**：获取、查询数据；参数展示在浏览器网址栏，适合搜索、查看页面
2. **POST**：提交新增数据；账号密码等数据藏在请求体内，不会暴露在地址栏
3. **PUT**：完整更新整条资源，一次性修改用户全部资料
4. **PATCH**：局部更新，仅修改单独一项字段（例如只修改昵称）
5. **DELETE**：删除评论、文件、用户等资源

### 4. IP 与 端口
1. `127.0.0.1`：本机回环地址，仅本机电脑能够访问服务
2. 端口：电脑程序的房门，Gin 默认端口 8080，依靠端口区分多个后端程序

### 5. 常见HTTP状态码
- 200：请求成功
- 404：接口地址不存在
- 500：后端代码报错、程序异常

### 6. 上下文‑gin.Context（c）
一次浏览器访问从头到尾所有数据的收纳工具箱，接收前端参数、存放返回数据。

---

## 二、入门阶段
### 1. 引擎初始化
```go
// 日常开发首选，自带日志、崩溃防护中间件
r := gin.Default()

// 纯净引擎，无内置工具，全部功能手动配置
r := gin.New()

// 设置上线正式模式，关闭冗余调试日志
gin.SetMode(gin.ReleaseMode)
```

### 2. 基础路由写法
```
r.GET("/see", func(c *gin.Context)  {c.String(200,"查看数据")})
r.POST("/add", func(c *gin.Context) {c.String(200,"提交新增")})
r.PUT("/allEdit", func(c *gin.Context){c.String(200,"完整修改")})
r.PATCH("/partEdit", func(c *gin.Context){c.String(200,"局部修改")})
r.DELETE("/del", func(c *gin.Context){c.String(200,"删除资源")})
```
### 3. 五种前端传参方式
路径参数 /user/:id
网址内嵌编号：127.0.0.1:8080/user/1001
```
r.GET("/user/:id",func(c *gin.Context){
    id := c.Param("id")
})
```
Query 问号参数？name = 小明
```
r.GET("/search",func(c *gin.Context){
    name := c.Query("name")
})
```
PostForm 表单参数（登录提交账号密码）
```
r.POST("/login",func(c *gin.Context){
    name := c.PostForm("username")
})
```
JSON 参数，前后端接口主流格式
结构体标签绑定数据，后面详细讲解
请求头、Cookie
```
// 获取浏览器标识
agent := c.GetHeader("User‑Agent")
// 获取浏览器Cookie
cookie,_ := c.Cookie("token")
```
### 1. 四种响应返回方式
```
// 1.返回纯文本
c.String(200,"hello gin")

// 2.返回JSON接口数据（项目最常用）
c.JSON(200,gin.H{"code":200,"msg":"成功"})

// 3.返回html网页模板
r.LoadHTMLFiles("index.html")
c.HTML(200,"index.html",gin.H{})

// 4.返回本地文件、图片
c.File("./test.jpg")
```
### 1. 路由分组
作用：归类管理同模块接口，统一接口前缀
```
userGroup := r.Group("/user")
{
    userGroup.GET("/info",func(c *gin.Context){})
    userGroup.POST("/login",func(c *gin.Context){})
}
```
三、中级核心知识点
1. 中间件 = 接口访问前的安检关卡
全局中间件：全部接口都需要经过安检

```r.Use(func(c *gin.Context){
    fmt.Println("请求进入")
    c.Next() // 放行执行业务代码
    fmt.Println("接口执行完毕")
})
```
路由组中间件：仅该分组下接口开启安检

userGroup.Use(LoginCheck)
单接口中间件：只对一条接口生效
```
r.GET("/pay",LoginCheck,func(c *gin.Context){})
c.Next ()：放行；业务代码结束后回头执行后续逻辑；常用于登录校验、接口计时
```
1. 结构体参数绑定（自动收纳前端参数）
结构体标签说明
```
type User struct{
    Name string `json:"name"` //接收json参数
    Pwd  string `form:"pwd"`  //接收表单参数
}
```
绑定 JSON 完整示例
```
r.POST("/login",func(c *gin.Context){
    var user User
    c.ShouldBindJSON(&user)
})
```
原理：结构体为收纳盒，标签对应前端字段名，自动装填数据，省去手动一行行接收参数
1. 文件上传
```
// 设置最大上传大小8MB
r.MaxMultipartMemory = 8 << 20

// 单文件上传
r.POST("/upload",func(c *gin.Context){
    file,err := c.FormFile("img")
    if err != nil{
        c.String(200,"未上传文件")
        return
    }
    c.SaveUploadedFile(file,"./upload/"+file.Filename)
})

// 多文件上传
files,_ := c.MultipartForm.File["img"]
for _,v := range files{
    c.SaveUploadedFile(v,v.Filename)
}
```
1. Cookie 和 Session
Cookie：身份小纸条存储在浏览器；数据存在客户端，容易被篡改，适合存放非敏感数据
Session：用户档案保存在服务器；浏览器只携带编号，安全性更高；占用服务器内存
现在主流方案：JWT‑Token，加密身份凭证，兼顾安全和便捷
1. 重定向、自定义 404 页面
```
// 网页跳转重定向
r.GET("/old",func(c *gin.Context){
    c.Redirect(302,"/new")
})

// 自定义404
r.NoRoute(func(c *gin.Context){
    c.String(404,"页面不存在")
})
```
# gin.Context(c) 方法分类笔记
## 核心概念
c *gin.Context = 单次请求的专属服务员、工具箱
所有 c.xxx() 只分为三类：读取浏览器数据、响应返回浏览器、特殊功能

## 一、读取浏览器传来的数据（接收参数）
| 方法 | 作用 |
|---|---|
| c.Param("id") | 获取路径参数 /user/10 |
| c.Query("name") | 获取url问号参数 ?name=张三 |
| c.PostForm("pwd") | 获取post表单提交参数 |
| c.ShouldBindJSON(&结构体) | json数据自动绑定结构体 |
| c.FormFile("img") | 获取单份上传文件 |
| c.MultipartForm.File["img"] | 获取多份上传文件 |
| c.GetHeader("User‑Agent") | 获取浏览器请求头信息 |
| c.Cookie("token") | 读取浏览器保存的cookie |

## 二、返回数据给浏览器（接口响应）
| 方法 | 作用 |
|---|---|
| c.String(200,"文本") | 返回普通字符串 |
| c.JSON(200,gin.H{}) | 返回json（项目最常用） |
| c.HTML(200,"xxx.html",data) | 返回网页模板 |
| c.File("图片地址") | 返回本地文件、图片 |
| c.Redirect(302,"/new") | 网页地址跳转 |

## 三、专属特殊方法
1. c.Next()
- 只用在中间件
- 含义：安检放行，执行业务路由代码
- 业务代码结束后，会回头执行Next后面代码

## 四、最简记忆口诀
1. 带读取、获取含义：拿浏览器的数据
2. 带返回、输出含义：给浏览器发送结果
3. c.Next：中间件放行开关

## 五、新手优先熟记高频API
### 获取参数
Param、Query、PostForm、ShouldBindJSON
### 返回响应
String、JSON
### 文件上传
FormFile、SaveUploadedFile
## 进阶
## 一、文件上传
```
//设置最大上传内存
r.MaxMultipartMemory = 8 << 20 //8MB

r.POST("/upload", func(c *gin.Context) {
	//前端表单name="file"
	file,err := c.FormFile("file")
	if err != nil {
		c.JSON(400,gin.H{"msg":"获取文件失败"})
		return
	}
	//保存到磁盘，生产环境不要直接用file.Filename，存在安全风险
	err = c.SaveUploadedFile(file,"./"+file.Filename)
	if err != nil {
		c.JSON(500,gin.H{"msg":"保存失败"})
		return
	}
	c.JSON(200,gin.H{"msg":"上传成功"})
})
```
## 二、Cookie & Session
Cookie
数据保存在用户浏览器本地；用户可以查看、修改、删除；不能存敏感数据。
```
//设置cookie
c.SetCookie("username","zhangsan",3600,"/","",false,true)
//读取cookie
val,err := c.Cookie("username")
//删除cookie：过期时间设为‑1
c.SetCookie("username","",‑1,"/","",false,true)
```
Session
真实业务数据保存在服务器；浏览器只存一个 session‑id（放在 cookie），安全。
Gin 原生无 session，需要第三方包 gin‑contrib/sessions
```
store,_ := sessions.NewCookieStore([]byte("密钥"))
r.Use(sessions.Sessions("mysession",store))

sess := sessions.Default(c)
sess.Set("userId",1001)
sess.Save() //必须save才生效

//读取
id := sess.Get("userId")
//清空
sess.Clear()
sess.Save()
```

| |Cookie|	Session|
|---|--|----|
|存储位置	|用户浏览器	|服务器|
|用户篡改	|可以篡改	|无法篡改真实业务数据|
|敏感数据	|禁止存放	|可以存放|
## 三、统一返回封装
所有接口返回格式统一，减少重复代码
```
func Success(c *gin.Context, data interface{}, msg string) {
	c.JSON(200, gin.H{
		"code":200,
		"msg":msg,
		"data":data,
	})
}

func Fail(c *gin.Context, code int, msg string) {
	c.JSON(200, gin.H{
		"code":code,
		"msg":msg,
		"data":nil,
	})
}

//接口中调用
Success(c, gin.H{"name":"张三"},"查询成功")
Fail(c,400,"参数错误")
```
## 四、GORM 基础（MySQL ORM）
ORM：写 Go 结构体，自动生成 SQL，不用手写大量 SQL 语句
1. 模型定义
```
type User struct {
	gorm.Model //内置ID、CreatedAt、UpdatedAt、DeletedAt(软删除)
	Name string `gorm:"size:32"`
	Age int
}
```
1. 连接数据库
```
dsn := "root:123456@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
db,err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
```
1. 自动建表
```
db.AutoMigrate(&User{})
```
4.CRUD
```
//新增
u := User{Name:"张三",Age:18}
db.Create(&u)

//id查询单条
var user User
db.First(&user,1)

//条件查询
db.Where("name = ?","张三").First(&user)

//更新
db.Model(&user).Update("age",20)

//删除：软删除，不会真正删除数据，给DeletedAt打时间标记
db.Delete(&user)
```
## 五、HTTP 常用状态码
200：成功
400：参数错误
401：未登录
403：权限不足
404：接口不存在
500：服务器内部错误
