# Logtransfer
结构：

* 从kafka获取日志信息
* 写入Elasticsearch
* 使用kibana展示日志结果


从kafka里面把日志取出来，写入ES，使用Kibana做可视化的展示
## Elasticsearch
特点：

* 分布式文档数据库
* 每个字段、数据可被搜索

应用场景：

* 搜索网上商品
* 处理日志数据、交易数据
* GitHub搜索

基本概念

* 几乎实时
* 基于集群，多个节点共同保存数据
* 索引里有一个或多个类型，类型是所有的分区
* 文档是索引的基本单位
* 分片和副本

对比常规数据库：

* 索引：数据库
* 类型：表
* 文档：数据行
* field：数据列
* mapping：模式

使用方法：

* restful API```curl -X GET 127.0.0.1:9200/_cat/health?v```

## 安装es
下载地址：https://www.elastic.co/cn/downloads/elasticsearch

### 运行
```
cd /Users/zhouyixiang/Documents/softwares/Elasticsearch/elasticsearch-7.8.1
bin/elasticsearch
# 默认端口号9200
# 运行起来后，在浏览器访问127.0.0.1:9200
```
## 使用

* 查询：get
* 创建索引：```curl -x put 192.0.0.1:9200/www```

解决Elasticsearch报内存不够的错误：

```
在配置文件末尾加
cluster.routing.allocation.disk.threshold_enabled: false

# 使用以下restful指令
put  http://127.0.0.1:9200/_all/_settings
headers: key Content-Type value application/json
body: 
{
  "index.blocks.read_only_allow_delete": null
}
```

插入

```
post 127.0.0.1:9200/student/go
body:
{
	"name":"zhangxiaohua",
	"age":18,
	"married":false
} 
```
精确查询

```
get 127.0.0.1:9200/student/go/_search
body: 
{
	"query":{
		"match":{
			"name":"ouyan"
		}
	}
}
```
分词后匹配任何字段

```
get 127.0.0.1:9200/student/go/_search
body: 
{
	"query":{
		"fuzzy":{
			"name":"o"
		}
	}
}
```

查询匹配某一条件

```
get 127.0.0.1:9200/student/go/_search
body: 
{
	"query":{
		"match":{
			"married":false
		}
	}
}
```
使用golang操作es

```
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Married bool   `json:"married"`
}
p1 := &Person{Name: "Rion", Age: 22, Married: false}
put1,err :=client.Index().Index("student").Type("go").BodyJson(p1)
```

## Kibana
# 系统监控
gopsutil做系统监控信息的采集，写入influxDB，使用grafana展示

prometheus监控，采集性能指标数据，保存起来，使用grafana展示
