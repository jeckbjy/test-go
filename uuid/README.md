# 常用UUID库

## 介绍
- [shortid](https://github.com/teris-io/shortid)  
源自[nodejs](https://github.com/dylang/shortid)版本  
可以生成非常短的id,但是也有一些限制  
在分布式情况下，需要明确的指定workid,才能保证唯一,且不能超过31  
DefaultABC中包含了0,1,l,o,这些易混淆字符，以及-,_特殊字符,自定义Alphabet要求必须是64个字符,很难满足
- [go-shortid](https://github.com/skahack/go-shortid)  
和shortid类似，仅仅实现不同,代码相对更容易理解一些，都是源自nodejs的，也要求64个字符,也就是base64编码
- [xid](https://github.com/rs/xid)  
生成20个字符的uuid,可应用于分布式下，保证唯一且有序
- [shortuuid](https://github.com/lithammer/shortuuid)  
依赖uuid库,使用base57对标准uuid进行编码
- [uuid](https://github.com/gofrs/uuid)  
标准uuid库实现
- [uuid](github.com/google/uuid)  
官方uuid库

## 生成结果
```text
teris id   :L6HINt7ZR
skahack id :H98ltZUDAf
xid        :bk1l7lrc1osksttjnbrg
shortuuid  :Hk5Q7dYXLSTR86zaiycmkU
gofrs_uuid :da225b33-23e4-42fe-bb2c-8f8a99f6aa45
google_uuid:d1b6aa19-9714-400d-97ea-c6d15259e212
```

## 总结
uuid的生成算法大体相似,都是时间戳+(机器标识)+(进程ID)+(uniqueID)+自增数，
然后将字符串进行base64或者base32，base64，base57，base58编码，字符集越长，最终结果越短.
有一些易混淆的字符,比如I,l,1,o,0，通常需要将其去除,但base64则很难选出这么多可用字符
一个疑问，有没有更加通用的编码方式不依赖字符集长度呢，即字符集长度是N，就是baseN编码

实际使用中有几种方案  
- 直接使用xid，缺点是有点长
- 只使用shortid，缺点是不好指定workid，需要一个中央服务器生成workid
- 优先使用shortid,当发生碰撞时再使用xid
- 在ID的使用方面,可以内部使用uint64，客户端展示使用baseN编码,这样既可以保证内部效率，又可以避免对外直接暴露数字ID

## 参考
- [base58 wiki](https://zh.wikipedia.org/wiki/Base58)
- [base64 wiki](https://zh.wikipedia.org/wiki/Base64)
- [base62x](https://github.com/wadelau/Base62x)
- [base62](// https://www.jianshu.com/p/3156cc5d6ae3)
