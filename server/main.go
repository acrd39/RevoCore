package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
)

// 配置结构体
type Config struct {
	DBHost     string `yaml:"db_host"`
	DBPort     int    `yaml:"db_port"`
	DBUser     string `yaml:"db_user"`
	DBPassword string `yaml:"db_password"`
	DBName     string `yaml:"db_name"`
}

func main() {
	// 初始化配置
	cfg := loadConfig()

	// 连接数据库
	db, err := connectPostgres(cfg)
	if err != nil {
		log.Fatal("数据库连接失败:", err)
	}

	// 初始化Gin引擎
	r := gin.Default()
	setupRoutes(r, db)

	// 启动服务
	if err := r.Run(":8080"); err != nil {
		log.Fatal("服务启动失败:", err)
	}
}

func setupRoutes(r *gin.Engine, db *gorm.DB) {
	// 健康检查接口
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// 用户模块路由组
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", registerUser)
		userGroup.POST("/login", loginUser)
	}
}