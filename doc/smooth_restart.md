### sharding服务重启步骤：

1. 重启前，先更新sharding服务状态
2. client端心跳检测到服务端状态后:
   1) 关闭空闲连接中该服务相关连接
   2) 已有连接

##
client端维护sharding server的地址？并进行心跳检测与路由？



