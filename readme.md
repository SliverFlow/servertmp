## 1、项目结构
```text
-api                 # 外部请求接口
-cmd                 # 程序运行
    -server 
        -main.go     # 程序入口
        -wire.go     # 依赖注入
        -wire_gen.go # 依赖注入生成文件
-config 
    -config.yaml     # 配置文件
-internal            # 内部程序
    -biz             # 业务逻辑
    -config          # 配置映射
    -core            # 服务核心程序
    -data            # 数据层
        -repo        # 数据操作接口
    -middleware      # 中间件
    -model           # 数据模型
        -request     # 请求模型
        -reply       # 响应模型
    -service         # 服务
-pkg                 # 公共包
    -zerror          # 自定义错误
    -util            # 工具
    -constant        # 常量
    -response        # 全局统一响应
```
## 2、依赖注入
```shell
# 保证项目安装了wire # github.com/google/wire
cd cmd/server
wire
```
### 3、ProviderSet
```go
api biz data middleware service 都有自己的ProviderSet，添加新的组件时，需要在对应的ProviderSet中添加构造函数
```