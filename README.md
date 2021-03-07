# http-proxy

简单混淆http代理服务器

```shell 

protocol:

-----------------------------------------------------------
    local <-- confuse ---> vps  <-- transparent -->  dst 
-----------------------------------------------------------
```

### Install

1. 安装服务端

* 下载最新的release
* go 命令`

```
    go get -u github.com/clearcodecn/http-proxy/server
```

2. 安装客户端

* 下载最新的release
* go 命令

```
    go get -u github.com/clearcodecn/http-proxy/client
```

### Use it

1. 运行服务端

```
server -a :9000
```

2. 运行客户端

```
client -s 192.168.199.1:9000 -l :9001  (替换ip地址为vps)
```

3. 设置浏览器代理

```shell
export http_proxy=http://localhost:9001
export https_proxy=http://localhost:9001
curl https://www.google.com
```

> 参照scripts脚本


[LICENSE](LICENSE)