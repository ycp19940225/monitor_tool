go build -o monitor_server

sudo chmod +x monitor_server
# nohup ./tool > monitor.log 2>&1 &
nohup ./monitor_server > monitor_server.log 2>&1 &
monitor_server
ps -aux|grep monitor_server
kill -9 进程号