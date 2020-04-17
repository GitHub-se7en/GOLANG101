package queue

import (
	"log"
	"time"
)

/**
map的key是topic，value是channel
pub的时候向channel里面放数据，这样另外一个线程监听这个channel的时候就会打印出信息
*/
type Queue struct {
	QueueMap map[string]chan *string
}

func NewQueue() *Queue {
	//这个只是初始化了map，并没有初始化channel，这个时候如果直接使用的话，会空指针
	//初始化map，只不过是零值，可以向里面放channel数据
	return &Queue{QueueMap: map[string]chan *string{}}
}

type TopicPool struct {
	TopicMap map[string]*Queue
}

func NewTopicPool() *TopicPool {
	return &TopicPool{TopicMap: make(map[string]*Queue)}
}
func (t *TopicPool) CreateTopic(topic string) {
	if _, ok := t.TopicMap[topic]; !ok {
		t.TopicMap[topic] = NewQueue()
	}
}

func (t *TopicPool) PubPool(message *string, topicName string) bool {
	queue, ok := t.TopicMap[topicName]
	if !ok {
		return false
	}
	//订阅这个topic的所有queue里面的map数据
	//我不管这个小map的key是多少
	for _, value := range queue.QueueMap {
		log.Println("value", value)
		value <- message
	}
	return true
}

func (t *TopicPool) SubPool(topic string, queueName string) *string {
	queue, ok := t.TopicMap[topic]
	if !ok {
		return nil
	}
	q, ok := queue.QueueMap[queueName]
	if !ok {
		queue.QueueMap[queueName] = make(chan *string)
		q = queue.QueueMap[queueName]
	}
	time.Sleep(time.Second * 10)
	return <-q
}
func (t *TopicPool) Bind(topic string, queueName string) {
	if _, ok := t.TopicMap[topic]; !ok {
		t.TopicMap[topic] = NewQueue()
	}
	t.TopicMap[topic].QueueMap[queueName] = make(chan *string)
}
func (queue *Queue) Pub(message *string, topicName string) {
	//map的key是topicName，value是一个channel，里面是message
	//我一开始也有这个疑问，会只是初始化第一层
	//queueMap := make(map[string]chan *string)
	//queueMap[topicName] = make(chan *string)
	//queueMap[topicName] <- message
	queue.QueueMap[topicName] <- message
	log.Println("queue是", queue)
}
func (queue *Queue) Sub(topicName string) *string {
	chanMessage := queue.QueueMap[topicName]
	//返回值，我是使用指针类型还是复制？我返回去别人还要不要修改同一个数据，应该是复制
	//but 人家用的是复制，可惜可惜
	return <-chanMessage
}
