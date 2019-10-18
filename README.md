# Streamnet-Gateway

Streamnet-Gateway 是一套支持http/https协议的网关系统，由gateway服务、manager服务构成。

### 环境要求
etcd安装

### 开发初衷
提供针对streamnet服务器请求的负载均衡，做到可视化配置，热生效

### 功能特性
+ 反向代理 
+ 黑名单
+ 访问日志
+ 限流
+ 熔断 【错误熔断、超时熔断、强制熔断】
+ 负载均衡 【轮询、随机】
+ 请求路径过滤、路径级私有负载均衡
+ 访问统计 【访问量、访问时长】
+ 访问拷贝 【复制请求数据、返回数据】

#### 域名配置
![hgw](https://github.com/triasteam/streamnet-gateway/blob/master/img/hgw.gif)

#### 路径配置 请求拦截
![hgw-path](https://github.com/triasteam/streamnet-gateway/blob/master/img/hgw-path.gif)

#### 强制熔断
![hgw-breaker](https://github.com/triasteam/streamnet-gateway/blob/master/img/hgw-breaker.gif)

#### 访问统计
![domain-metrics](https://github.com/triasteam/streamnet-gateway/blob/master/img/metrics.png)

## 安装
#### 1. 获取hgw代码
```
git clone https://github.com/triasteam/streamnet-gateway.git
```

gateway文件夹是网关服务的核心代码

manager文件夹是控制服务的核心代码


#### 2. 编译gateway
```
go build gateway.go
```

#### 3. 编译manager
```
go build manager.go
```

#### 4. 运行gateway
```
[root@localhost gateway]# ./gateway -h
usage: gateway --ser-name=SER-NAME --addr=ADDR --etcd=ETCD [<flags>]

Flags:
  -h, --help               Show context-sensitive help (also try --help-long and --help-man).
      --ser-name=SER-NAME  SerName: gateway listen addr
      --addr=ADDR          Addr: gateway listen addr
      --tls-addr=""        Tls-Addr: gateway tls listen addr
      --etcd=ETCD          Addr: etcd server addr
      --u=""               Username: etcd username
      --p=""               Password: etcd password
      --version            Show application version.
```

##### 参数
--ser-name 【gateway服务的识别名称】

--addr 【gateway http服务的监听地址】

--tls-addr 【gateway https服务的监听地址】

--etcd 【连接etcd服务的地址】

--u 【连接etcd服务的账户】

--p 【连接etcd服务的密码】

```
./gateway --ser-name=gateway-1 --addr=0.0.0.0:80 --etcd=127.0.0.1:2379
```

#### 5. 运行manager
```
[root@localhost manager]# ./manager -h
usage: manager --addr=ADDR --etcd=ETCD [<flags>]

Flags:
  -h, --help       Show context-sensitive help (also try --help-long and --help-man).
      --addr=ADDR  gateway listen addr
      --etcd=ETCD  etcd server addr
      --u=""       Username: etcd username
      --p=""       Password: etcd password
```
##### 参数
--addr 【manager 服务的监听地址】

--etcd 【连接etcd服务的地址】

--u 【连接etcd服务的账户】

--p 【连接etcd服务的密码】

```
./manager --addr=0.0.0.0:7000 --etcd=127.0.0.1:2379
```

#### 6. 使用
访问 manager监听的服务地址+/admin/ 管理gateway服务。

<font color="red">**注**</font>： 自己搭建服务，第一次访问请先访问 /admin/init.html 初始化管理账号密码。


## 感谢
[hwg](https://github.com/dmhao/hgw) 源码项目