# 树莓派搭建k3s


# 树莓派安装k3s
## 1.安装k3s
### 控制节点  

    curl -sfL https://get.k3s.io | sh -
    cat /var/lib/rancher/k3s/server/node-token
  
### 工作节点  

    curl -sfL https://get.k3s.io | K3S_URL=https://myserver:6443 K3S_TOKEN=mynodetoken sh -
  
树莓派特别要注意一个坑，就是关于内存的问题这个之后再讲  

    k3s kubectl get nodes
    #显示正确的节点表示完成
  
### 卸载 k3s  

    #server 节点
    /usr/local/bin/k3s-uninstall.sh
    #agent 节点
    /usr/local/bin/k3s-agent-uninstall.sh

## 2.安装dashboard k3s面板
### 部署 Kubernetes 仪表盘  

    GITHUB_URL=https://github.com/kubernetes/dashboard/releases
    VERSION_KUBE_DASHBOARD=$(curl -w '%{url_effective}' -I -L -s -S ${GITHUB_URL}/latest -o /dev/null | sed -e 's|.*/||')
    sudo k3s kubectl create -f https://raw.githubusercontent.com/kubernetes/dashboard/${VERSION_KUBE_DASHBOARD}/aio/deploy/recommended.yaml
  
### 仪表盘 RBAC 配置
创建以下资源清单文件：  
dashboard.admin-user.yml  

    apiVersion: v1
    kind: ServiceAccount
    metadata:
        name: admin-user
        namespace: kubernetes-dashboard
  
dashboard.admin-user-role.yml  
  
    apiVersion: rbac.authorization.k8s.io/v1
    kind: ClusterRoleBinding
    metadata:
        name: admin-user
    roleRef:
        apiGroup: rbac.authorization.k8s.io
        kind: ClusterRole
        name: cluster-admin
    subjects:
        - kind: ServiceAccount
          name: admin-user
          namespace: kubernetes-dashboard
  
部署admin-user 配置：  
  
    sudo k3s kubectl create -f dashboard.admin-user.yml -f dashboard.admin-user-role.yml

获得 Bearer Token  
  
    sudo k3s kubectl -n kubernetes-dashboard describe secret admin-user-token | grep '^token'

现在可以通过以下网址访问仪表盘： 

    sudo k3s kubectl proxy
  
http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

### 连接lens 
    cat /etc/rancher/k3s/k3s.yaml
    更改本地host
    穿透服务器IP    local
     
    

## 3.安装 kube—prometheus 监控
### 一键安装
    wget https://github.com/prometheus-operator/kube-prometheus/archive/refs/tags/v0.9.0.tar.gz
    tar -zxvf v0.9.0.tar.gz
    cd kube-prometheus-0.9.0/manifests
    k3s kubectl apply -f setup/
    k3s kubectl get pod -n monitoring
    k3s kubectl apply -f .

  
### 一键卸载
    cd kube-prometheus/manifests
    k3s kubectl delete -f .
    k3s kubectl delete -f setup/

## 4.安装 nfs外部驱动挂载storageclass

## 5.创建有状态pods（mysql）
