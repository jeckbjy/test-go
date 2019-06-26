# study CustomResourceDefinition

## [Demo](https://medium.com/velotio-perspectives/extending-kubernetes-apis-with-custom-resource-definitions-crds-139c99ed3477)

### kubectl命令,在kubernetes目录中
- 创建CRD: kubectl create -f sslconfig-crd.yaml 
- 查看CRD: kubectl get crd
- 应用CRD: kubectl create -f sslconfig-obj.yaml
- 查看CRD-obj:kubectl get sslconfig 
  
### demo代码
- 下载依赖(测试环境):go get k8s.io/client-go@master
- [k8s client-go](https://github.com/kubernetes/client-go)
- [原始repo](https://github.com/velotio-tech/crd-example)
- 新的客户端需要实现DeepCopyObject()接口函数,可以使用deepcopy-gen工具自动生成
- TypeMeta已经实现了GetObjectKind()接口函数
- deepcopy.sh用于自动生成deepcopy代码
- build.sh 用于编译生成bin文件

## 官方demo
- [sample-controller](https://github.com/kubernetes/sample-controller)

## 其他Demo
- [Extend Kubernetes 1.7 with Custom Resources](https://thenewstack.io/extend-kubernetes-1-7-custom-resources/)
- [kube-crd](https://github.com/yaronha/kube-crd)
- [k8s代码自动生成过程的解析](http://blog.xbblfz.site/2018/09/19/k8s%E4%BB%A3%E7%A0%81%E8%87%AA%E5%8A%A8%E7%94%9F%E6%88%90%E8%BF%87%E7%A8%8B%E7%9A%84%E8%A7%A3%E6%9E%90/)