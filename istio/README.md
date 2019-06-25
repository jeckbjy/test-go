# istio学习

## 文档

- [official](https://istio.io/docs/setup/kubernetes/install/)
- [aws](https://eksworkshop.com/servicemesh_with_istio/install/)
  
## 安装
- istio安装分为三种方式,官方推荐使用helm安装
  1. 直接使用kubectl安装 
  2. 使用helm templat生成配置，然后调用kubectl安装,好处是不需要helm tiller
  3. 使用helm install直接安装
- 安装流程,虽然分为三种安装方式,但安装流程大体一致
  1. 下载istio,curl -L https://git.io/getLatestIstio | ISTIO_VERSION=1.2.0 sh -  
  2. 安装[Custom Resource Definitions](https://kubernetes.io/docs/concepts/extend-kubernetes/api-extension/custom-resources/#customresourcedefinitions),简单介绍一下CRDs功能,k8s给我提供了一些常用的资源,比如Pod，Deployment，service等，而CRDs使k8s拥有了扩展资源的能力，第三方可以自己定义新的资源类型
  3. 安装istio,install/kubernetes/helm/istio目录下提供了helm的charts,可以通过配置不同values安装相应需要的功能,预提供的[Configuration Profiles](https://istio.io/docs/setup/kubernetes/additional-setup/config-profiles/),至此istio就已经安装完了
  4. 安装自己应用,比如安装[Bookinfo sample](https://istio.io/docs/examples/bookinfo/)

## 进阶 [task](https://istio.io/docs/tasks/)
- [Traffic Management](https://istio.io/docs/tasks/traffic-management/)
  - Gateway: 定义了HTTP/TCP等连接协议
  - DestinationRule:定义如何分发流量，如ROUND_ROBIN，LEAST_CONN等
  - VirtualService:定义了路由规则，根据不同url转发到不同服务器
- [Security]
- [Policies]
- [Telemetry]
