package test

import (
	"net/http"
	"net/http/httptest"
	"revocore/server/handlers"
	"revocore/server/pkg/database"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func SetupTestRouter() *gin.Engine {
	r := gin.Default()

	// 初始化 SQLite 内存数据库
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})

	// 自动迁移你的模型
	db.AutoMigrate(&database.User{}) // 替换成你实际的模型

	// 把数据库绑定到 gin context
	r.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	// 注册路由
	userGroup := r.Group("/users")
	{
		userGroup.POST("/register", handlers.RegisterUser)
		userGroup.POST("/login", handlers.LoginUser)
	}

	return r
}

func TestRegisterSuccess(t *testing.T) {
	router := SetupTestRouter()

	body := `{"username": "testuser", "password": "123456"}`
	req, _ := http.NewRequest("POST", "/users/register", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Contains(t, w.Body.String(), "testuser")
}

func TestRegisterDuplicate(t *testing.T) {
	router := SetupTestRouter()

	// 先注册一次
	body := `{"username": "duplicate", "password": "123"}`
	req1, _ := http.NewRequest("POST", "/users/register", strings.NewReader(body))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	// 再注册一次
	req2, _ := http.NewRequest("POST", "/users/register", strings.NewReader(body))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 409, w2.Code)
	assert.Contains(t, w2.Body.String(), "用户名已存在")
}

func TestLoginSuccess(t *testing.T) {
	router := SetupTestRouter()

	// 先注册
	registerBody := `{"username": "loginuser", "password": "mypwd"}`
	req1, _ := http.NewRequest("POST", "/users/register", strings.NewReader(registerBody))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	// 再登录
	loginBody := `{"username": "loginuser", "password": "mypwd"}`
	req2, _ := http.NewRequest("POST", "/users/login", strings.NewReader(loginBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 200, w2.Code)
	assert.Contains(t, w2.Body.String(), "token")
}

func TestLoginWrongPassword(t *testing.T) {
	router := SetupTestRouter()

	// 注册
	body := `{"username": "wrongpwd", "password": "goodpwd"}`
	req1, _ := http.NewRequest("POST", "/users/register", strings.NewReader(body))
	req1.Header.Set("Content-Type", "application/json")
	w1 := httptest.NewRecorder()
	router.ServeHTTP(w1, req1)

	// 登录失败
	loginBody := `{"username": "wrongpwd", "password": "badpwd"}`
	req2, _ := http.NewRequest("POST", "/users/login", strings.NewReader(loginBody))
	req2.Header.Set("Content-Type", "application/json")
	w2 := httptest.NewRecorder()
	router.ServeHTTP(w2, req2)

	assert.Equal(t, 401, w2.Code)
}
