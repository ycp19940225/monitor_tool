package main

import (
	"fmt"
	"github.com/robfig/cron/v3" //插件
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	//serverHost := "smtp.qq.com"
	//serverPort := 465
	//fromEmail := "820363773@qq.com"  //发送者邮箱
	//fromPasswd := "" //  授权码
	//subject := "服务器告警:ip-"
	//body := ``
	//TimeSettle(subject, body, serverHost, fromEmail, fromPasswd, "curry.yang@zenwell.cn", serverPort)
	//TimeSettle(subject, body, serverHost, fromEmail, fromPasswd, "820363773@qq.com", serverPort)

	c := cron.New()
	c.AddFunc("*/3 * * * *", func() {
		serverHost := "smtp.qq.com"
		serverPort := 465
		fromEmail := "820363773@qq.com" //发送者邮箱
		fromPasswd := ""                //  授权码
		//myToers :="hongery@yeah.net"// "li@latelee.org, latelee@163.com" 逗号隔开 接收者邮箱
		myCCers := "" //"readchy@163.com" 抄送
		ips := []string{"http://abc178.zenwell.cn/", "http://abc179.zenwell.cn/", "http://abc180.zenwell.cn/"}

		for _, ip := range ips {
			res := clientTest(ip)
			expectedString := `{"msg":"缺少必要的参数：token","code":-1}`
			if res != expectedString {
				subject := "服务器告警:ip-" + ip
				body := "服务器告警:ip-" + ip
				TimeSettle(subject, body, serverHost, fromEmail, fromPasswd, myCCers, serverPort)
			} else {
				fmt.Println(ip + "：" + time.Now().Format("2006-01-02 15:04:05") + "正常")
			}
		}
	})

	// 每天
	c.AddFunc("0 9 */1 * *", func() {
		serverHost := "smtp.qq.com"
		serverPort := 465
		fromEmail := "820363773@qq.com"  //发送者邮箱
		fromPasswd := "azpcihwgxomibfgj" //  授权码

		//myToers :="hongery@yeah.net"// "li@latelee.org, latelee@163.com" 逗号隔开 接收者邮箱
		myCCers := "" //"readchy@163.com" 抄送
		ips := []string{"http://abc178.zenwell.cn/", "http://abc179.zenwell.cn/", "http://abc180.zenwell.cn/"}

		for _, ip := range ips {
			res := clientTest(ip)
			expectedString := `{"msg":"缺少必要的参数：token","code":-1}`
			if res == expectedString {
				subject := "服务器正常:ip-" + ip
				body := "服务器正常:ip-" + ip
				TimeSettle(subject, body, serverHost, fromEmail, fromPasswd, myCCers, serverPort)
			}
		}
	})

	c.Start()

	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()

	select {}
}

func clientTest(url string) string {
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

func TimeSettle(subject, body, serverHost, fromEmail, fromPasswd, myCCers string, serverPort int) {

	// "eric.qin@zenwell.cn", "bob.shang@zenwell.cn",
	tos := []string{"765276145@qq.com", "curry.yang@zenwell.cn", "eric.qin@zenwell.cn", "bob.shang@zenwell.cn"}

	for _, to := range tos {
		time.Sleep(5)
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

}
