package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var broadcast = make(chan Message)

var clients = make(map[*websocket.Conn]bool)

//为什么要配置一个upgrader？？？？
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	//创建一个简单的文件服务器
	//访问服务器上的静态文件，java中可以使用tomcat，nginx
	fs := http.FileServer(http.Dir("./public"))
	http.Handle("/", fs)

	//创建websocket路由
	http.HandleFunc("/ws", handleConnection)

	//监听所有的message
	go handleMessage()

	log.Println("监听8080端口")
	e := http.ListenAndServe(":8080", nil)
	if e != nil {
		log.Fatal("启动服务器出错", e)
	}

}
func handleConnection(w http.ResponseWriter, r *http.Request) {
	//upgrade 把GET请求变成一个websocket
	ws, e := upgrader.Upgrade(w, r, nil)
	if e != nil {
		log.Fatal("把request和response升级为ws的时候出现异常", e)
	}

	defer ws.Close()

	clients[ws] = true

	for true {
		var msg Message
		//按照正常的逻辑来说应该是从request中取出来message
		//可是这里只是把一个变量给了
		e := ws.ReadJSON(&msg)
		if e != nil {
			log.Println("websocket读取message出现异常", e)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		//把消息推送到channel中
		broadcast <- msg
	}
}
func handleMessage() {
	for {
		msg := <-broadcast
		for client := range clients {
			e := client.WriteJSON(msg)
			if e != nil {
				log.Fatal("", e)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
