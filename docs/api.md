# RevoCore 全栈开发项目文档

## 📚 项目概述
**RevoCore** 是一个基于 Go + Vue3 的跨平台产品评测社区，支持 Web/小程序/Android/iOS 多端访问。项目采用云原生架构设计，核心功能包含用户评测、产品对比、内容社交互动等模块。

## 🛠️ 技术栈总览

### 后端基础设施
| 组件             | 技术选型                  | 用途                          |
|------------------|--------------------------|------------------------------|
| 开发框架         | Gin                      | HTTP 路由/中间件              |
| 数据库           | PostgreSQL + Redis       | 主数据存储 + 缓存/会话         |
| 消息队列         | Kafka                    | 异步任务处理                  |
| 对象存储         | MinIO                    | 图片/视频存储                 |
| API 网关         | Nginx                    | 负载均衡/反向代理             |
| 服务通信         | gRPC + REST              | 内部服务高性能通信             |
| 容器编排         | Docker + Kubernetes      | 服务部署与扩缩容              |
| 监控告警         | Prometheus + Grafana     | 系统指标监控                  |

### 前端架构
| 端               | 技术方案                 |
|------------------|--------------------------|
| 网页端           | Vue3 + Vant              |
| 移动端           | Uni-app（跨端编译）      |
| 状态管理         | Pinia                    |
| 构建工具         | Vite                     |

## 🗺️ 开发路线图

### 阶段 1：核心功能实现（4-6 周）
- 用户注册登录（JWT 鉴权）
- 产品库管理（品牌/品类/产品）
- 评测内容发布（图文混排）
- 基础评论互动

### 阶段 2：进阶功能（2-3 周）
- 产品对比矩阵
- 图片上传与管理
- 内容标签系统
- 简易推荐算法

### 阶段 3：多端适配（2 周）
- 微信小程序编译
- Android/iOS 打包
- 响应式网页优化

## 🖥️ 开发环境搭建

### 必备工具清单
```bash
# Go 语言环境
go version ≥ 1.21

# Node.js 环境
nvm install 18

# Docker Desktop
docker --version ≥ 20.10

# 数据库客户端
brew install libpq  # PostgreSQL CLI
```

# API

## 产品模块

### 创建品牌
```
POST /api/v1/brands
Content-Type: application/json

{
  "name": "Apple",
  "logo": "https://cdn.example.com/logo.png"
}
```

### 产品对比
```
GET /api/v1/products/compare?a=1&b=2
```

## 内容模块

### 发布评测
```
POST /api/v1/posts
Authorization: Bearer <token>
Content-Type: application/json

{
  "product_id": 1, 
  "content": "实测表现优异...",
  "images": ["img1.jpg"],
  "tags": ["数码"]
}
```

## 数据库设计

```sql
-- 产品表
CREATE TABLE products (
  id SERIAL PRIMARY KEY,
  brand_id INT REFERENCES brands(id),
  name VARCHAR(255) NOT NULL,
  specs JSONB NOT NULL
);

-- 评测内容表
CREATE TABLE posts (
  id SERIAL PRIMARY KEY,
  user_id INT REFERENCES users(id),
  content TEXT NOT NULL,
  images TEXT[],
  tags VARCHAR(50)[]
);

-- 评论表（支持嵌套）
CREATE TABLE comments (
  id SERIAL PRIMARY KEY,
  post_id INT REFERENCES posts(id),
  parent_id INT REFERENCES comments(id),
  content TEXT NOT NULL
);
```