package main

import (
	"fmt"
	"github.com/robfig/cron/v3" //插件
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	c := cron.New()
	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("222")
		serverHost := "smtp.qq.com"
		serverPort := 465
		fromEmail := "820363773@qq.com"  //发送者邮箱
		fromPasswd := "azpcihwgxomibfgj" //  授权码

		//myToers :="hongery@yeah.net"// "li@latelee.org, latelee@163.com" 逗号隔开 接收者邮箱
		myCCers := "" //"readchy@163.com" 抄送
		body := ``
		ips := []string{"http://abc178.zenwell.cn/", "http://abc179.zenwell.cn/", "http://abc180.zenwell.cn/"}

		for _, ip := range ips {
			res := clientTest(ip)
			fmt.Println(res)

			fmt.Println(res != `{"msg":"缺少必要的参数：token","code":-1}`)
			if res != `{"msg":"缺少必要的参数：token","code":-1}` {
				subject := "服务器告警:ip-" + ip
				TimeSettle(subject, body, serverHost, fromEmail, fromPasswd, myCCers, serverPort, "765276145@qq.com")
			}
		}
	})
	c.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select {}
}

func clientTest(url string) string {
	fmt.Println("222")
	// 创建一个新的HTTP客户端对象
	client := &http.Client{}

	// 构造一个GET请求
	req, _ := http.NewRequest("GET", url, nil)

	// 添加必要的头部信息（如果有）
	req.Header.Add("Content-Type", "application/json")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	// 从响应体中读取数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}

func TimeSettle(subject, body, serverHost, fromEmail, fromPasswd, myCCers string, serverPort int, to string) {
	return
	// 结构体赋值
	myEmail := &EmailParam{
		ServerHost: serverHost,
		ServerPort: serverPort,
		FromEmail:  fromEmail,
		FromPasswd: fromPasswd,
		Toers:      to,
		CCers:      myCCers,
	}
	//    发送邮件
	InitEmail(myEmail)
	SendEmail(subject, body)
}
