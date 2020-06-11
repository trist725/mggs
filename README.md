# mggs
My Game Gate Server

### 说明

- 配合[mlgs](http://https://github.com/trist725/mlgs "mlgs")使用的游戏网关,底层基于[leaf框架](https://github.com/name5566/leaf "leaf框架")；
- client模块监听对外客户端连接；
- server模块监听对内游服连接；
- client和server模块的消息会路由给gate模块做转发;
- 通过重写processor,gate和server之间的通信会在包头加上4字节的客户端ID用于区分不同的客户端；
- 默认使用最小连接数分配游服;
- 使用该网关需要对[leaf框架](https://github.com/name5566/leaf "leaf框架")有一定理解。

#### 预期未实现的功能
- 超时主动关闭连接;
- 限制同一客户端IP的连接数用于防止恶意连接;
- 服务端握手认证;
- **分别提取出前后端filter模块,在两端收发包前后实现自定义业务;**
- 转发游服间的消息;
- 各种负载均衡算法,如虚拟一致性hash;
- 配置中心-服务注册与发现;
- 开启reuseport；
- 内存池。
