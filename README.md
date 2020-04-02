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

第三个项目：

handler接口，同一个url，使用多个处理器处理

这个抽象出来了一个一个函数类型，但凡是使用这种形式的函数

都是这种类型，和第一个项目很相似，抽象接口，利用接口写程序

而这个却是抽象函数，忽略函数的具体实现，使用抽象函数完成逻辑


第四个项目：

面试题目，对go的语法有用处，这个是原文链接，有些格式不对，我调整了
https://gist.github.com/stamhes/6b82a695bfac83b7448b4ca3dfa380f3

添加第二个面试题目的链接

https://github.com/lifei6671/interview-go/blob/master/question/q014.md