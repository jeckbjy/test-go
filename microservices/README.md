# 微服务架构实践

## 参考

- [Microservices in Go](https://medium.com/seek-blog/microservices-in-go-2fc1570f6800)在这篇文章里详细对比了Go Micro，Go Kit，Gizmo，Kite这几个被熟知的库,[例子代码仓库](https://github.com/antklim/go-microservices)
- [TarsGo](https://medium.com/@sandyskieschan/a-high-performance-microservice-framework-in-golang-a-linux-foundation-project-ec7ba899173)这里介绍了腾讯开源微服务框架TarsGo，据说性能比gRPC高5倍

## 总结
- go-micro功能比较齐全，使用比较简单,但如果想同时支持grpc和web则比较繁琐，需要手动写很多代码
- go-kit和gizmo比较类似，gizmo封装了一下go-kit，使用上比较繁琐