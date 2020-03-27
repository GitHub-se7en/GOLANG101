# GOLANG101
go语言，rss，goroutine，channel
使用手册：    

第一个项目：

在main方法里面传入key word，会查询rss里面所有包含这个关键字的文章    

重点是run方法里面根据几十个rss启动几十个goroutine    

我没有试过java里面是怎么实现的，但是go明显简洁    
 
第二个项目：

rpc服务器的实现，conn连接了客户端和服务器，

http包下面的handler也是基于这个conn升级的