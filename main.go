package main

func main() {
	serverHost := "smtp.qq.com"
	serverPort := 465
	fromEmail := "820363773@qq.com"  //发送者邮箱
	fromPasswd := "azpcihwgxomibfgj" //  授权码

	//myToers :="hongery@yeah.net"// "li@latelee.org, latelee@163.com" 逗号隔开 接收者邮箱
	myCCers := "" //"readchy@163.com" 抄送

	subject := "服务器告警"
	body := `这是正文<br>
            <h3>这是标题</h3>
             Hello <a href = "www.baidu.com">主页</a><br>`

	TimeSettle(subject, body, serverHost, fromEmail, fromPasswd, myCCers, serverPort)
}

func TimeSettle(subject, body, serverHost, fromEmail, fromPasswd, myCCers string, serverPort int) {
	// 结构体赋值
	myEmail := &EmailParam{
		ServerHost: serverHost,
		ServerPort: serverPort,
		FromEmail:  fromEmail,
		FromPasswd: fromPasswd,
		Toers:      "765276145@qq.com",
		CCers:      myCCers,
	}
	//    发送邮件
	InitEmail(myEmail)
	SendEmail(subject, body)
}
