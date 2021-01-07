# mock

## 流程
- 定义接口 DB interface{}
- 生成mock文件
    - mockgen -source=db.go -destination=db_mock.go -package=main
    - 或者使用go generate
- 编写test测试，TestGetFromDB

## 资料
- 官网: https://github.com/golang/mock
- https://geektutu.com/post/quick-gomock.html
- https://www.jianshu.com/p/598a11bbdafb