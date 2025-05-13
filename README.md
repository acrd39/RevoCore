# RevoCore - 跨平台产品评测社区

[![Go Version](https://img.shields.io/badge/go-1.21+-blue.svg)](https://golang.org/)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

**用Go+Vue3构建的现代评测社交平台** | [在线演示](https://your-demo-link.com) | [API文档](https://api-docs.revocore.dev)

## 🚀 核心功能
### 用户系统
- 手机号/第三方登录（微信、Google）
- 个人主页与收藏夹
- 用户等级体系（后续开发）

### 内容生态
- 📝 多格式评测（图文/视频/语音）
- 🗣️ 多级嵌套评论（支持@用户）
- 🏷️ 智能标签系统（自动分类）

### 产品库
- 🏭 品牌-品类-产品三级结构
- 🔍 产品对比矩阵（参数可视化）
- 🎯 智能推荐（基于用户行为）

## 🛠️ 技术栈
### 后端服务
| 组件               | 技术选型                 |
|--------------------|-------------------------|
| 开发框架           | Gin                     |
| 数据库             | PostgreSQL + Redis      |
| 对象存储           | MinIO                   |
| 消息队列           | Kafka                   |
| 服务通信           | gRPC + REST             |
| 分布式追踪         | Jaeger                  |

### 前端架构
| 端                 | 技术方案                |
|--------------------|-------------------------|
| 网页端             | Vue3 + Vant             |
| 移动端             | Uni-app（跨端编译）     |
| 状态管理           | Pinia                   |
| 构建工具           | Vite                    |

### 基础设施
```mermaid
graph TD
  A[Docker] --> B{Kubernetes}
  B --> C[Prometheus]
  B --> D[Grafana]
  B --> E[Cert-Manager]