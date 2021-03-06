# Raft总结

## Raft论文概述

### 1 Raft 将一致性算法分解成了几个关键模块:
* 领导人选举、日志复制、安全性
* Raft算法还包括一个新的机制来允许集群成员的动态改变，它利用重叠的大多数来保证安全性。

Raft 算法在许多方面和现有的一致性算法都很相似（主要是 Oki 和 Liskov 的 Viewstamped Replication），但是它也有一些独特的特性：

* **强领导者**：和其他一致性算法相比，Raft 使用一种更强的领导能力形式。比如，日志条目只从领导者发送给其他的服务器。这种方式简化了对复制日志的管理并且使得 Raft 算法更加易于理解。
* **领导选举**：Raft 算法使用一个随机计时器来选举领导者。这种方式只是在任何一致性算法都必须实现的心跳机制上增加了一点机制。在解决冲突的时候会更加简单快捷。
* **关系调整**：Raft 使用一种共同一致的方法来处理集群成员变换的问题，在这种方法中，两种不同的配置都要求的大多数机器会重叠。这就使得集群在成员变换的时候依然可以继续工作。
