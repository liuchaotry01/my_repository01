package main


import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

// Configure the upgrader
// 配置升级程序
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Define our message object
// 定义我们的信息对象
type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	// Create a simple file server
	//  创建一个简单的文件服务器
	fs := http.FileServer(http.Dir("../public"))
	http.Handle("/", fs)

	// Configure websocket route
	//  配置websocket路线
	http.HandleFunc("/ws", handleConnections)

	// Start listening for incoming chat messages
	// 开始监听传入的聊天消息
	go handleMessages()

	// Start the server on localhost port 8000 and log any errors
	//  在本地主机端口8000启动服务器并记录任何错误
	log.Println("http server started on :7845")
	err := http.ListenAndServe(":7845", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade initial GET request to a websocket
	//将初始GET请求升级到websocket
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}
	// Make sure we close the connection when the function returns
	// 当函数返回时，请确保关闭连接
	defer ws.Close()

	// Register our new client
	// 注册我们的新客户
	clients[ws] = true

	for {
		var msg Message
		// Read in a new message as JSON and map it to a Message object
		// 将新消息读入JSON并将其映射到消息对象
		err := ws.ReadJSON(&msg)
		if err != nil {
			log.Printf("error: %v", err)
			delete(clients, ws)
			break
		}
		// Send the newly received message to the broadcast channel
		// 将新收到的消息发送到广播频道
		broadcast <- msg
	}
}

func handleMessages() {
	for {
		// Grab the next message from the broadcast channel
		// 从广播频道获取下一条消息
		msg := <-broadcast
		// Send it out to every client that is currently connected
		// 把它发送给当前连接的每个客户端
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}