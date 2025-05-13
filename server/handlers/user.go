package handlers

import (
	"net/http"
	"revocore/server/pkg/database"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func RegisterUser(c *gin.Context) {
	db := c.MustGet("db").(*gorm.DB)

	var req UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 密码需要加密存储（示例未加密，实际需使用bcrypt）
	user := database.User{
		Username: req.Username,
		Password: req.Password,
	}

	if err := db.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "用户创建失败"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"id": user.ID})
}

func LoginUser(c *gin.Context) {
	// 登录逻辑实现（JWT生成等）
}
