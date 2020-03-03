# 学习MIT6.824分布式系统课程

- 课程1:实现MapReduce
  - 实现目标
    - common_map.go:实现doMap函数
    - common_reduce.go:实现doReduce函数
    - schedule.go:实现schedule函数
    - main/wc.go:实现mapF和reduceF函数,注意英文分隔,不能只用空格分隔,测试文档包含标点符号
  - 理论基础
    - 函数原型:map(String key, String value),reduce(String key, Iterator values),map需要处理kv结构,reduce需要处理的是key->string list结构
    - M个Map任务,R个Reduce任务,R由用户指定,M可以是将一个大文件分隔成M个小文件
    - Map处理的中间结果,保存到文件中,再由Reduce处理
    - M个Map,R个Reduce将会有M*R个中间结果,map需要将key hash到对应的reduce中。那么为什么这样?因为只有这样才能保证reduce处理时,每个key只会在一个reduce中被处理
    - Shuffle:主要用于将Map产生的中间数据转移到对应的Reduce节点,并进行合并,相同的key的value合并成一个数组,再根据key排序
  - 实现细节:
    - map和reduce的实现比较简单,主要是中间文件的生成与读取,格式的定义,shuffle的实现
    - schedule的实现稍微复杂一些，schedule用于rpc调度map和reduce,保证所有的任务都能完成后才能进行下一步,同时还需要监控异常节点,失败后自动转移到其他节点重新执行
  - 注:
    - 源代码master_rpc.go:48行是有bug的,format格式不对,会导致test跑不起来

## 参考

- [6.824: Distributed Systems](https://pdos.csail.mit.edu/6.824/)
- [实现](https://github.com/yixuaz/6.824-2018)
- [知乎](https://www.zhihu.com/question/29597104)
- [gitee收藏的相关资源](https://gitee.com/tantexian/MIT6.824)
- [blog](https://thesquareplanet.com/blog/students-guide-to-raft/)
- [bilibili视频](https://www.bilibili.com/video/av87684880/)
- [MapReduce Tutorial](https://www.tutorialscampus.com/tutorials/map-reduce/index.htm)