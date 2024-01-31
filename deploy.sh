set GOARCH=amd64
go env -w GOARCH=amd64
set GOOS=linux
go env -w GOOS=linux

go build -o monitor_server

sudo chmod +x monitor_server
# nohup ./tool > monitor.log 2>&1 &
nohup ./monitor_server > monitor_server.log 2>&1 &
monitor_server
ps -aux|grep monitor_server
kill -9 进程号
sudo pkill monitor_server