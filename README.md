# kubeconfig_merge


用于合并kubeconfig文件的工具

## 前提条件
1. $HOME/.kube 必需目录
2. $HOME/.kube/kubeconfig 必需目录，这个目录下存放kubeconfig文件，文件名必需以.conf结尾，context的名称将会使用你的文件名

## 使用方法

```shell
## 拉去项目
git clone https://github.com/lucs-t/merge_conf.git
## 如果mac是apple芯片
make build_apple
## 如果mac是intel芯片
make build_intel
## 其他系统自己写吧，就是go语言的交叉编译

## 运行
./merge_conf

## 查看contexts
kubectl config get-contexts
CURRENT   NAME       CLUSTER    AUTHINFO   NAMESPACE
          fz-idc-0   fz-idc-0   fz-idc-0   
*         local-0    local-0    local-0   

## 选择一个context
kubectl config use-context local-0
```

## 希望可以帮到你!!!