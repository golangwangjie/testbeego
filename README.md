# testbeego
为面试搭建的beego测试平台
测试功能是在windows10环境下，mysql，
虚拟机搭建的redis-cluster集群（三主三从），zookeeper集群（一个leader，两个follower）

技术栈
golang

beego 
controllers包完成了测试的业务调用。
initial包完成了mysql，redis，zookeeper的初始化连接。
logs包完成日志的初始化。
models包完成mysql数据库的具体操作。
utils包完成redis，zookeeper的具体操作。

mysql
 使用了官网测试数据库sakila，完成actor表的数据查询。
 并给数据库sakila，新增users表，完成users表数据的增删改查。

redis
string类型
完成最为数据缓存的增删改查。
    bitmap
    完成了用户签到的例子功能。

待更新。。。。
hash类型
 功能例子：详情页，购物车
list类型
 功能例子：关注，论坛评论
set类型
 功能例子：共同好友，好友推荐
zset类型
 功能例子：排行榜

zookeeper
完成节点的增删改查。
实现分布式锁。
实现发布和订阅功能。

待更新。。。。
rabbitmq