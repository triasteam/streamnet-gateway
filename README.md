[中文说明](https://github.com/triasteam/streamnet-gateway/blob/master/README_cn.md)
# Streamnet-Gateway

Streamnet-Gateway is a gateway system supporting HTTP/HTTPS protocol, which consists of gateway service and manager service.

### Environment
etcd service is necessary

### Why Designed
Provide the load balance,visual configuration and hot effective for streamnet server.

### Functional Characteristics
+ Reverse proxy 
+ Blacklist
+ Access log
+ Current limiting
+ fuses [Wrong Fusing, Overtime Fusing, Forced Fusing]
+ load balance [Polling, Random]
+ Request path filtering, path level private load balancing
+ Access statistics [Visit Volume, Visit Duration]
+ Access copy [copy request data, response data]

#### Domain Configuration
![hgw](https://github.com/triasteam/streamnet-gateway/blob/master/img/hgw.gif)

#### Path Configuration & Request Interception
![hgw-path](https://github.com/triasteam/streamnet-gateway/blob/master/img/hgw-path.gif)

#### Forced Fusing
![hgw-breaker](https://github.com/triasteam/streamnet-gateway/blob/master/img/hgw-breaker.gif)

#### Access statistics
![domain-metrics](https://github.com/triasteam/streamnet-gateway/blob/master/img/metrics.png)

## Install
#### 1. 获取streamnet-gateway代码
```
git clone https://github.com/triasteam/streamnet-gateway.git
```

gateway Folder is the code of gateway service

manager Folder is the code for management service


#### 2. Compile gateway service
```
go build gateway.go
```

#### 3. Compile management service
```
go build manager.go
```

#### 4. Run Gateway
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

##### Paramater
--ser-name [gateway server name(address/ser-name)]

--addr [gateway http listen address]

--tls-addr [gateway https listen address]

--etcd [etcd service address]

--u [etcd account]

--p [etcd password]

```
./gateway --ser-name=gateway-1 --addr=0.0.0.0:80 --etcd=127.0.0.1:2379
```

#### 5. Run Manager
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
##### Paramater
--addr [manager service listen address]

--etcd [etcd service address]

--u [etcd account]

--p [etcd password]

```
./manager --addr=0.0.0.0:7000 --etcd=127.0.0.1:2379
```

#### 6. Use Guide
Use browser to view
http://service address/admin/ 

<font color="red">**注**</font>： First time, please visit /admin/init.html to init account and password。


## Thanks to
[hwg](https://github.com/dmhao/hgw) 