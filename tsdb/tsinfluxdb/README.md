# influxdb 2.0
从2.0开始将不再支持InfluxQL,取而代之的是使用Flux

安装： brew install influx
启动服务:influxd  
可视化界面: http://localhost:9999/  
执行程序生成或查询数据(需要修改代码)  
go run main.go   
influxdb写入使用Line protocol协议,一条记录需要包括,名字,时间戳,tags,fields  
influxdb使用[flux](https://v2.docs.influxdata.com/v2.0/query-data/get-started/)查询语言  
influxdb数据输出格式为csv  
输出数据    
```text
2019/09/10 14:37:28 ,result,table,_start,_stop,_field,_measurement,cluster,host,_value
,_result,0,2019-09-10T04:30:00Z,2019-09-10T04:40:00Z,cpu_usage,node_status,private,192.168.1.0,14.412640129345162 
,_result,1,2019-09-10T04:30:00Z,2019-09-10T04:40:00Z,cpu_usage,node_status,private,192.168.1.1,8.965080851863629
,_result,2,2019-09-10T04:30:00Z,2019-09-10T04:40:00Z,cpu_usage,node_status,private,192.168.1.10,3.6845035812246345
,_result,3,2019-09-10T04:30:00Z,2019-09-10T04:40:00Z,cpu_usage,node_status,private,192.168.1.11,74.81227383775688
```

在服务器调试中的使用：  
可以作为tag标识的:
core: project,environment(dev,alpha,beta,prod),service,host,error  
k8s:  namespace,pod name,image name,container name,

## FAB: 
1:如何查询某个tag下的所有可选字段
```text
from(bucket: "my-bucket")
  |> range(start: -12h)
  |> keep(columns: ["host"])
  |> distinct(column: "host")
```
返回数据结果
```csv
2019/09/10 15:53:12 ,result,table,host,_value
,_result,0,192.168.0.0,192.168.0.0
,_result,1,192.168.0.1,192.168.0.1
,_result,2,192.168.0.10,192.168.0.10
,_result,3,192.168.0.11,192.168.0.11
,_result,4,192.168.0.12,192.168.0.12
```
2:log内容如何做全文检索  
时序数据库并不能作为日志全文索引  
InfluxDB is not designed to satisfy full-text search or log management use cases and therefore would be out of scope. For these use cases, we recommend sticking with Elasticsearch or similar full-text search engines
