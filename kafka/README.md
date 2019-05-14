# kafka实践

- 例子: https://medium.com/@yusufs/getting-started-with-kafka-in-golang-14ccab5fa26
- 源码: https://github.com/yusufsyaifudin/go-kafka-example/tree/accadecfca65c956bf03a6199ff1f3944bc6d7dc
- 启动kafka：MY_IP=192.168.18.34 docker-compose -f docker-compose-kafka.yml up  
- api进程负责写入kafka,worker负责监听kafka,读取消息
