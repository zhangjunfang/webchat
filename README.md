# webchat
web chat tool
##  测试
### 测试说明
使用tcpudp_2.1.1.exe  
使用tcp与服务器通信  
发送数据量16字节 
每毫秒发送10条数据
### 运行环境
服务器配置centos7虚拟机  
配置2g内存  
cpu 1核  E3-1231 v3 3.4Ghz
mysql 5.7.17  innodb   单机
redis 3.2.8    单机
### 测试结果
每毫秒可以处理10条数据  当大于10条之后 提示连接拒绝  【没有做任何的系统优化】
