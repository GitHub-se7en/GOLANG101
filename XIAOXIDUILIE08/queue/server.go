package queue

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"runtime"
	"strconv"
	"strings"
)

func StartServer() {
	listener, e := net.Listen("tcp", "127.0.0.1:8080")
	if e != nil {
		log.Fatal("启动端口失败", e)
	}
	for true {
		//这个conn就是最重要的部分
		conn, e := listener.Accept()
		if e != nil {
			//如果accept失败了之后，直接跳过
			continue
		}
		go serveConn(conn)
	}
}

func serveConn(conn net.Conn) {
	defer func() {
		clientAddr := conn.RemoteAddr().String()
		r := recover()
		if err, ok := r.(error); ok {
			//这个是创建切片啊
			buf := make([]byte, 4096)
			buf = buf[:runtime.Stack(buf, false)]
			log.Println("出错的远程client ip是", clientAddr,
				"stack日志记录为", string(buf),
				"出错记录为", err)
			conn.Write([]byte("-ERROR " + err.Error() + "\r\n"))
		}
		conn.Close()
	}()
	for true {
		//这就有点难受了，根据reply的不同，write的方式也是不一样的
		reply, err := ConvertConnToRequest(conn)
		//少了一步逻辑，应该是得到request，然后根据request做逻辑处理
		if err != nil {
			conn.Write([]byte("-ERROR " + err.Error() + "\r\n"))
			//conn.Close()
			//log.Fatalln(err)
		}
		//conn write之后就结束了吗？
		conn.Write([]byte("$" + strconv.Itoa(len(reply)) + "\r\n" + reply + "\r\n"))
		//conn.Close()
		//log.Fatalln("我要结束这次请求，不能死循环")
	}
}

func ConvertConnToRequest(conn net.Conn) (string, error) {
	reader := bufio.NewReader(conn)
	log.Println("reader", reader)
	line, err := reader.ReadString('\n')
	log.Println("line", line)
	if err != nil {
		return "", err
	}
	var argCount int
	if line[0] == '*' {
		if _, err := fmt.Sscanf(line, "*%d\r\n", &argCount); err != nil {
			return "", err
		}
		log.Println("argCount", argCount)
		command, err := readArgument(reader)
		log.Println("command", command)
		if err != nil {
			return "", err
		}
		arguments := make([][]byte, argCount-1)
		for i := 0; i < argCount-1; i++ {
			if arguments[i], err = readArgument(reader); err != nil {
				log.Println("argument", arguments)
				return "", err
			}
		}
		return strings.ToUpper(string(command)), nil
	}
	return "", errors.New("最终还是出错了")
}

func readArgument(reader *bufio.Reader) ([]byte, error) {
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, err
	}
	var argLength int
	if _, err := fmt.Sscanf(line, "$%d\r\n", &argLength); err != nil {
		return nil, err
	}
	data, err := ioutil.ReadAll(io.LimitReader(reader, int64(argLength)))
	log.Println("io utils read all-- data is", data)
	log.Println("io utils read all-- data to string is ", string(data))
	if err != nil {
		return nil, err
	}
	if len(data) != argLength {
		return nil, errors.New("长度不符合")
	}
	//我的目的自然是如果是b是\r或者是\n的时候放过去
	//但是这一行判断出来，不管b是什么都会阻断
	//\r\n之外的阻断，怎么实现？？
	//把b不等于的判断条件翻译一下，变成b等于，就容易理解多了
	//比如说，b不等于并且b不等于就相当于，b等于两者之外的数值呗
	//我在想高数，线代，概率论里面还有别的技巧能走出逻辑迷宫吗？
	if b, err := reader.ReadByte(); err != nil || (b != '\r' && b != '\n') {
		log.Println("what is b why miss crlf", b, '\n', "is", byte('\n'), byte('\r'))
		return nil, errors.New("缺少CRLF")
	}
	return data, nil
}
