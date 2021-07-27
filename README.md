# 基于gRPC的微服务项目

## 项目简介

   项目采用微服务架构，涉及到的主要技术有:Golang、gRPC、go-micro、Gin、Consul、Mysql、Redis等

   项目分为客户端和服务端，客户端使用Gin框架向网页端提供HTTP请求的接口，服务端采用go-micro框架实现微服务，使用Consul做服务发现。客户端与服务端使用protobuf交换数据。

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
   - view                 ## 视图资源
   - main.go



