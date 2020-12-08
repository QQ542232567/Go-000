package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"golang.org/x/sync/errgroup"
)

// func main() {

// 	http.HandleFunc("/", myHandler)
// 	http.HandleFunc("/close", myClose)
// 	http.ListenAndServe("127.0.0.1:8000", nil)

// }

// // handler函数
// func myHandler(w http.ResponseWriter, r *http.Request) {
// 	// 回复
// 	w.Write([]byte("gogogo"))
// }

// // handler函数
// func myClose(w http.ResponseWriter, r *http.Request) {
// 	// 回复
// 	w.Write([]byte("close"))
// }

var (
	addr1 string = ":8001"
	addr2 string = ":8002"
)

type HomeHandler struct {
	keyId string
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(h.keyId))
}

type ShutDownHandler struct {
	Server *http.Server
}

func (h *ShutDownHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("模拟关闭"))

	err := h.Server.Shutdown(context.Background()) //主动退出
	if err != nil {
		fmt.Printf("service close err %+v \n", err)
	}
}

func main() {

	fmt.Println("week 03 start")

	group := new(errgroup.Group)

	//服务1
	chan1 := make(chan error)
	shutdown1 := &ShutDownHandler{}
	mux1 := http.NewServeMux()
	mux1.Handle("/", &HomeHandler{keyId: "gogogo1"})
	mux1.Handle("/shutdown", shutdown1)

	server1 := &http.Server{Addr: addr1, Handler: mux1}
	shutdown1.Server = server1

	group.Go(func() error {
		err := server1.ListenAndServe() // 阻塞

		close(chan1)
		fmt.Printf("server1 closed %+v \n", err)
		return errors.New("服务1返回的失败")
	})

	//服务2
	chan2 := make(chan error)
	shutdown2 := &ShutDownHandler{}
	mux2 := http.NewServeMux()
	mux2.Handle("/", &HomeHandler{keyId: "gogogo2"})
	mux2.Handle("/shutdown", shutdown2)

	server2 := &http.Server{Addr: addr2, Handler: mux2}
	shutdown2.Server = server2

	group.Go(func() error {
		err := server2.ListenAndServe() // 阻塞

		close(chan2)
		fmt.Printf("server2 closed %+v \n", err)
		return errors.New("服务2返回的失败")
	})

	// ctrl+c kill
	ch := make(chan os.Signal, 10)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)

	group.Go(func() (err error) {
		// 收到任何一个信号就关闭服务
		select {
		case <-ch:
			fmt.Println("receive close signal!")
			err = errors.New("信号错误")
		case <-chan1:
			fmt.Println("receive server1 close!")
			err = errors.New("服务1错误")
		case <-chan2:
			fmt.Println("receive server2 close!")
			err = errors.New("服务2错误")
		}

		fmt.Println("----------------------")

		signal.Stop(ch)

		//
		err2 := server2.Close()
		fmt.Printf("#############server2 close %+v \n", err2)

		err1 := server1.Close()
		fmt.Printf("#############server1 close %+v \n", err1)

		return err
	})

	err := group.Wait()

	fmt.Printf("group err %+v \n", err)
}
