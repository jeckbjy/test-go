# minikube在mac下使用

- [安装文档](https://kubernetes.io/docs/tasks/tools/install-minikube/)
- mac下快速安装:brew cask install minikube
- 启动: minikube start --vm-driver hyperkit
- 删除: minikube delete
- 查看ip:minikube ip
- 设置默认driver：minikube config set vm-driver hyperkit
- 查看dashboard: minikube dashboard
- [可选的driver](https://github.com/kubernetes/minikube/blob/master/docs/drivers.md#hyperkit-driver)
- [Medium教程](https://medium.com/@wisegain/minikube-cheat-sheet-a273385e66c9)
