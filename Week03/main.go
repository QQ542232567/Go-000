// tcp/server/main.go

// TCP server端
package main

import "fmt"

type Iface interface {
	Hello(i int) error
}

type Gface interface {
	HelloB(i int) error
}

type Istruct struct {
}

//隐式
var _ Iface = (*Istruct)(nil)

//显式
var _ Iface = &Istruct{}
var _ Gface = (*Istruct)(nil)

func (c *Istruct) Hello(i int) error {
	return nil
}
func (c *Istruct) HelloB(i int) error {
	return nil
}
func main() {

	fmt.Println("go")

}

// import (
// 	"bufio"
// 	"fmt"
// 	"net"
// )

// // 处理函数
// func process(conn net.Conn) {
// 	defer conn.Close() // 关闭连接
// 	for {
// 		reader := bufio.NewReader(conn)
// 		var buf [128]byte
// 		n, err := reader.Read(buf[:]) // 读取数据
// 		if err != nil {
// 			fmt.Println("read from client failed, err:", err)
// 			break
// 		}
// 		recvStr := string(buf[:n])
// 		fmt.Println("收到client端发来的数据：", recvStr)
// 		conn.Write([]byte(recvStr)) // 发送数据
// 	}
// }

// func main() {
// 	listen, err := net.Listen("tcp", ":0")
// 	if err != nil {
// 		fmt.Println("listen failed, err:", err)
// 		return
// 	}
// 	fmt.Println("Server [tcp] Listening on %s", listen.Addr().String())

// 	for {
// 		conn, err := listen.Accept() // 建立连接
// 		if err != nil {
// 			fmt.Println("accept failed, err:", err)
// 			continue
// 		}
// 		go process(conn) // 启动一个goroutine处理连接
// 	}
// }
