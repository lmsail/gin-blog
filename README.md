## 简单说两句  
机缘巧合之下，接触了`Go`语言，先后看过`Beego`、`Gin`，也了解过`IRIS`，最后被`Gin`所吸引，跟之前所接触的`PHP`框架不同，`Gin`的项目目录自己组建，自由度很高，文档齐全，上手快，所以决定写点东西，既然是`Web`出身，当然后选择从`Web`入手，熟悉它的基础语法与实际应用先。。  

## 项目介绍  
> 项目结构有点像`PHP`框架，个人习惯，可随意调整，作为入门Go的练手项目，适合新手入门，使用`Gin+Gorm`搭建而成       

- 项目结构  
结构相对简单直观，不做过多赘述  
```text
gin-lmsail
├── app
│   ├── Helpers
│   ├── Http
│   ├── Middleware
│   ├── Models
│   ├── Service
│   ├── Task
├── config
├── resource
├── routers
├── storage
├── views
└── .env        ------ 配置文件
└── main.go
```  

- 实现功能，暂时只实现了博客的基础功能
	- 登录
	- 注册
	- 文章列表
	- 个人中心

## 三方依赖
- Gin [中文文档](https://www.kancloud.cn/shuangdeyu/gin_book/949438)
- Gorm [中文文档](http://gorm.book.jasperxu.com/)
- Sessions [Github地址](github.com/gin-contrib/sessions)
- Goconfig [Github地址](github.com/Unknwon/goconfig)
- Semantic UI [Github地址](https://www.semantic-ui.com/)
- ....

## 安装
- 更改 .env 文件中数据库链接相关配置 [mysql]
- 注意本地使用时的路径加载问题，可根据情况切换使用`Helpers` -> `GetGoRunPath、GetCurrentDirectory`方法

## 关于迭代  
这个再说吧，可能后面会转战`微服务`相关的学习，毕竟`Go-Web`只是用来过渡/入门学习  