# 基于gRPC的微服务项目

## 代码结构

 - service                     ## 服务端
   - getCaptcha       ##   验证码相关微服务
     - handler       ##  处理逻辑
     - model         ## 数据库操作
     - proto          ## proto文件
     - main.go
     - Makefile
   - user                  ##  登录及个人信息微服务(结构同上)
 - web                         ## 客户端
   - conf                  ## 相关文件
   - controller          ## 控制器
   - model               ## 部分数据库操作
   - proto                ## proto文件
   - test                  ## 测试代码
   - utils                  ## 工具方法
   - view                 ## 前端资源
   - main.go
