package main

import (
	"log"
	"revocore/server/handlers"
	"revocore/server/pkg/config"
	"revocore/server/pkg/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig("./configs/local.yaml")
	if err != nil {
		log.Fatal("配置加载失败:", err)
	}

	// 连接数据库
	db, err := database.ConnectPostgres(cfg)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 初始化Gin
	r := gin.Default()

	// 注入数据库到上下文
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// 设置路由
	setupRoutes(r)

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}

func setupRoutes(r *gin.Engine) {
	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 用户路由
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", handlers.RegisterUser)
		userGroup.POST("/login", handlers.LoginUser)
	}
}
