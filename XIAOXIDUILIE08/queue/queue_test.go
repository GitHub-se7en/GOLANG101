package queue

import (
	"log"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestQueue(t *testing.T) {
	wg.Add(2)
	pool := NewTopicPool()
	message := "hello"
	topic := "打招呼"
	//queueName := "zhaoweiguodequeue"
	queueName1 := "zhaoweiguodequeue"
	queueName2 := "zhaoweiguodequeue"
	//pool.Bind(topic,queueName)
	pool.Bind(topic, queueName1)
	go func() {
		defer wg.Done()
		pool.CreateTopic(topic)
		result := pool.PubPool(&message, topic)
		//由于使用的是无缓冲通道，把数据放到channel里面之后就停止了，导致程序停留了10秒之后，才继续打印这行日志
		log.Println("resultPub", result)
	}()
	//如果没问题的话，这行代码应该会一直停留着在里
	go func() {
		defer wg.Done()
		//subPool := pool.SubPool(topic, queueName)
		subPool := pool.SubPool(topic, queueName2)
		log.Println("subPool", *subPool)
	}()
	wg.Wait()

}
