# 测试学习go-micro和micro

## 简介
- 本例整理一下官网[booking](https://github.com/micro/examples/tree/master/booking)的例子
- 结构:  
src/api:frontend,用于前端web调用
src/srv目录下:backend，用于给frontend提供功能服务
- 使用[Docker Compose](https://docs.docker.com/compose/)运行方法
```
$ make build
$ make run
```
## 测试

非法的token
```bash
curl -H 'Content-Type: application/json' \
       -H "Authorization: Bearer INVALID_TOKEN" \
       -d '{"inDate": "2015-04-09"}' \
        http://localhost:8080/hotel/rates
        
{"id":"api.hotel.rates","code":401,"detail":"Unauthorized","status":"Unauthorized"}
```
非法的数据,没有inDate或者outDate
```bash
curl -H 'Content-Type: application/json' \
       -H "Authorization: Bearer VALID_TOKEN" \
       -d '{"inDate": "2015-04-09"}' \
        http://localhost:8080/hotel/rates

{"id":"api.hotel.rates","code":400,"detail":"Please specify inDate/outDate params","status":"Bad Request"}
```
合法的请求
```bash
curl -H 'Content-Type: application/json' \
       -H "Authorization: Bearer VALID_TOKEN" \
       -d '{"inDate": "2015-04-09", "outDate": "2015-04-10"}' \
        http://localhost:8080/hotel/rates
```    
返回结果
```json
{
    "hotels": [
        {
            "id": "3",
            "name": "Hotel Zetta",
            "phoneNumber": "(415) 543-8555",
            "description": "A 3-minute walk from the Powell Street cable-car turnaround and BART rail station, this hip hotel 9 minutes from Union Square combines high-tech lodging with artsy touches.",
            "address": {
                "streetNumber": "55",
                "streetName": "5th St",
                "city": "San Francisco",
                "state": "CA",
                "country": "United States",
                "postalCode": "94103"
            }
        },
        {
            "id": "4",
            "name": "Hotel Vitale",
            "phoneNumber": "(415) 278-3700",
            "description": "This waterfront hotel with Bay Bridge views is 3 blocks from the Financial District and a 4-minute walk from the Ferry Building.",
            "address": {
                "streetNumber": "8",
                "streetName": "Mission St",
                "city": "San Francisco",
                "state": "CA",
                "country": "United States",
                "postalCode": "94105"
            }
        },
        {
            "id": "5",
            "name": "Phoenix Hotel",
            "phoneNumber": "(415) 776-1380",
            "description": "Located in the Tenderloin neighborhood, a 10-minute walk from a BART rail station, this retro motor lodge has hosted many rock musicians and other celebrities since the 1950s. It’s a 4-minute walk from the historic Great American Music Hall nightclub.",
            "address": {
                "streetNumber": "601",
                "streetName": "Eddy St",
                "city": "San Francisco",
                "state": "CA",
                "country": "United States",
                "postalCode": "94109"
            }
        },
        {
            "id": "6",
            "name": "The St. Regis San Francisco",
            "phoneNumber": "(415) 284-4000",
            "description": "St. Regis Museum Tower is a 42-story, 484 ft skyscraper in the South of Market district of San Francisco, California, adjacent to Yerba Buena Gardens, Moscone Center, PacBell Building and the San Francisco Museum of Modern Art.",
            "address": {
                "streetNumber": "125",
                "streetName": "3rd",
                "city": "San Francisco",
                "state": "CA",
                "country": "United States",
                "postalCode": "94109"
            }
        },
        {
            "id": "1",
            "name": "Clift Hotel",
            "phoneNumber": "(415) 775-4700",
            "description": "A 6-minute walk from Union Square and 4 minutes from a Muni Metro station, this luxury hotel designed by Philippe Starck features an artsy furniture collection in the lobby, including work by Salvador Dali.",
            "address": {
                "streetNumber": "495",
                "streetName": "Geary St",
                "city": "San Francisco",
                "state": "CA",
                "country": "United States",
                "postalCode": "94102"
            }
        }
    ],
    "ratePlans": [
        {
            "hotelId": "1",
            "code": "RACK",
            "inDate": "2015-04-09",
            "outDate": "2015-04-10",
            "roomType": {
                "bookableRate": 109,
                "totalRate": 109,
                "totalRateInclusive": 123.17,
                "code": "KNG"
            }
        }
    ]
}
```

## 依赖库
```bash
$ go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
$ go get -u github.com/micro/protoc-gen-micro
$ go get -u github.com/micro/go-micro
$ go get -u github.com/hailocab/go-geoindex
```

- 依赖工具
    - [consul](https://www.consul.io/docs/index.html)  
    web ui地址: http://localhost:8500   
    
    - [micro](https://micro.mu/docs/)
    查看服务: micro --registry=consul --registry_address=localhost:8500 list services  
    启动web: micro --registry=consul --registry_address=localhost:8500 web  
    web ui地址: http://localhost:8082  

## 其他文章
- [再见，micro](https://songrgg.github.io/microservice/goodbye-micro/)
- [架构介绍](https://www.cnblogs.com/li-peng/p/9558421.html)

## 一些想法
- 感觉目前go-micro更适合于类似http的rest请求，并不支持服务器主动推送
- go-micro的handler中并不保证线程安全，需要自己加锁