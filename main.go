package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-gomail/gomail"
	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered:", r)
		}
	}()
	fmt.Println("start monitor server...")
	viperConfig := viper.New()
	viperConfig.AddConfigPath(".")
	viperConfig.SetConfigFile("config.yaml")
	err := viperConfig.ReadInConfig()
	if err != nil {
		fmt.Println("Error on Reading Viper Config")
		panic(err)
	}
	server := createEmailServer(
		viperConfig.GetString("email.config.host"),
		viperConfig.GetInt("email.config.port"),
		viperConfig.GetString("email.config.username"),
		viperConfig.GetString("email.config.password"),
	)
	ips := viperConfig.GetStringSlice("email.ips")
	tos := viperConfig.GetStringSlice("email.tos")
	c := cron.New()
	m := gomail.NewMessage()
	m.SetHeader("From", viperConfig.GetString("email.config.from"))
	m.SetHeader("To", tos...)
	// */2 * * * *
	c.AddFunc("*/2 * * * *", func() {
		for _, ip := range ips {
			res := clientTest(ip)
			if res.StatusCode != 200 {
				subject := "服务器告警:ip-" + ip
				body := "服务器告警:ip-" + ip
				m.SetHeader("Subject", subject)
				m.SetBody("text/html", body)
				//m.Attach("../../go.mod")
				dial, err := server.server.Dial()
				if err != nil {
					fmt.Println("邮件服务器连接错误:" + err.Error())
				}
				server.send(&dial, m)
			} else {
				fmt.Println(ip + "：" + time.Now().Format("2006-01-02 15:04:05") + "正常")
			}
		}
	})

	// 每天  30 7,11 * * *
	c.AddFunc("30 7,11 * * *", func() {
		content := ""
		subject := "服务监控程序自检"

		for _, ip := range ips {
			resp := clientTestContent(ip)

			var data map[string]string
			json.Unmarshal([]byte(resp), &data)

			// Access and print the parsed data
			mysqlData := data["mysql"]
			redisData := data["redis"]

			content += "<h4>" + ip + "</h4>"
			content += "mysql:\n" + mysqlData + "<br><br>"
			content += "redis:\n" + redisData + "<br><br>------------------------------------------------------<br><br>"
		}
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", content)
		dial, err := server.server.Dial()
		if err != nil {
			fmt.Println("邮件服务器连接错误:" + err.Error())
		}
		server.send(&dial, m)
	})

	// 每天
	c.AddFunc("30 13 L * *", func() {
		clientClear()
	})

	c.Start()
	//关闭着计划任务, 但是不能关闭已经在执行中的任务.
	defer c.Stop()
	select {}
}

func clientTest(url string) *http.Response {
	// 创建一个新的HTTP客户端对象
	client := &http.Client{}

	// 构造一个GET请求
	req, _ := http.NewRequest("GET", url, nil)

	// 添加必要的头部信息（如果有）
	req.Header.Add("Content-Type", "application/json")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	defer resp.Body.Close()
	return resp
}

func clientTestContent(url string) string {
	// 创建一个新的HTTP客户端对象
	client := &http.Client{}

	// 构造一个GET请求
	req, _ := http.NewRequest("GET", url, nil)

	// 添加必要的头部信息（如果有）
	req.Header.Add("Content-Type", "application/json")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	defer resp.Body.Close()
	// 从响应体中读取数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	return string(body)
}

func clientClear() {
	// 创建一个新的HTTP客户端对象
	client := &http.Client{}

	// 构造一个GET请求
	req, _ := http.NewRequest("GET", "https://sdsh.scgsdsj.com/external/tool.tool/clear?sign=zenwell123456.", nil)

	// 添加必要的头部信息（如果有）
	req.Header.Add("Content-Type", "application/json")

	// 发送请求并获取响应
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	// 从响应体中读取数据
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(body))
}
