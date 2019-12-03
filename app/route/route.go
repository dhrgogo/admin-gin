package route

import (
	"admin-gin/app/controller/admin"
	"admin-gin/app/controller/jaeger_conn"
	"admin-gin/app/controller/product"
	"admin-gin/app/controller/test"
	"admin-gin/app/route/middleware/exception"
	"admin-gin/app/route/middleware/jaeger"
	"admin-gin/app/route/middleware/logger"
	"admin-gin/app/util/response"
	"github.com/gin-gonic/gin"
)

func SetupRouter(engine *gin.Engine) {

	//设置路由中间件
	engine.Use(logger.SetUp(), exception.SetUp(), jaeger.SetUp())

	//404
	engine.NoRoute(func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(404, "请求方法不存在", nil)
	})

	engine.GET("/ping", func(c *gin.Context) {
		utilGin := response.Gin{Ctx: c}
		utilGin.Response(1, "pong", nil)
	})

	// 测试链路追踪
	engine.GET("/jaeger_test", jaeger_conn.JaegerTest)

	//@todo 记录请求超时的路由

	ProductRouter := engine.Group("/product")
	{
		// 新增产品
		ProductRouter.POST("", product.Add)

		// 更新产品
		ProductRouter.PUT("/:id", product.Edit)

		// 删除产品
		ProductRouter.DELETE("/:id", product.Delete)

		// 获取产品详情
		ProductRouter.GET("/:id", product.Detail)
	}

	// 测试加密性能
	TestRouter := engine.Group("/test")
	{
		// 测试 MD5 组合 的性能
		TestRouter.GET("/md5", test.Md5Test)

		// 测试 AES 的性能
		TestRouter.GET("/aes", test.AesTest)

		// 测试 RSA 的性能
		TestRouter.GET("/rsa", test.RsaTest)
	}

	// 管理端 数据库操作
	AdminRouter := engine.Group("/admin")
	{
		// 创建用户
		AdminRouter.POST("/user", admin.UserCreat)
		// 更新用户
		AdminRouter.PUT("/user/:unique_id", admin.UserUpdate)
		//获取单个用户
		AdminRouter.GET("/user/:id", admin.UserObj)
		// 用户列表
		AdminRouter.GET("/user", admin.UserList)

		// 创建角色
		AdminRouter.POST("/role", admin.RoleCreat)
		// 更新角色
		AdminRouter.PUT("/role/:unique_id", admin.RoleUpdate)
		// 根绝唯一id获取角色
		AdminRouter.GET("/role/:unique_id", admin.RoleObj)
		// 获取角色列表
		AdminRouter.GET("/role", admin.RoleList)

	}
}
