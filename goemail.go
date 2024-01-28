package main

import (
	"fmt"
	"github.com/go-gomail/gomail"
	"io"
	"sync"
)

var emailServer *EmailServer
var once = &sync.Once{}

type EmailServer struct {
	server *gomail.Dialer
}

type MySender interface {
	Send(from string, to []string, msg io.WriterTo) error
	Close() error
}

func createEmailServer(host string, port int, username string, password string) *EmailServer {
	if emailServer == nil {
		once.Do(func() {
			dialer := gomail.NewDialer(host, port, username, password)
			emailServer = &EmailServer{
				server: dialer,
			}
		})
	}
	return emailServer
}

func (email *EmailServer) send(dial *gomail.SendCloser, m *gomail.Message) {
	// 将 *gomail.SendCloser 转换为 MySender 接口类型
	mySenderInstance := MySender(*dial)
	err2 := gomail.Send(mySenderInstance, m)
	if err2 != nil {
		fmt.Println("发送报错:" + err2.Error())
	}
}
