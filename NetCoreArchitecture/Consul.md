- 使用 docker 获取 consul

```
docker pull consul
```

- 启动 consul

```
docker run --name consul -d -p 8500:8500 consul
```

- 浏览器打开 http://localhost:8500/

# 集群

- node1(server)

```
docker run -d --name=node1 --restart=always \
    -e 'CONSUL_LOCAL_CONFIG={"skip_leave_on_interrupt":true}' \
    -p 8300:8300 \
    -p 8301:8301 \
    -p 8301:8301/udp \
    -p 8302:8302 \
    -p 8302:8302/udp \
    -p 8400:8400 \
    -p 8500:8500 \
    -p 8600:8600 \
    -h node1 \
    consul agent -server -bind=0.0.0.0 -bootstrap-expect=3 -node=node1 \
    -data-dir=/tmp/data-dir -client 0.0.0.0 -ui
```

- node2(server)

```
docker run -d --name=node2 --restart=always \
    -e 'CONSUL_LOCAL_CONFIG={"skip_leave_on_interrupt":true}' \
    -p 9300:8300 \
    -p 9301:8301 \
    -p 9301:8301/udp \
    -p 9302:8302 \
    -p 9302:8302/udp \
    -p 9400:8400 \
    -p 9500:8500 \
    -p 9600:8600 \
    -h node2 \
    consul agent -server -bind=0.0.0.0 \
    -join=8.131.87.88 -node-id=$(uuidgen | awk '{print tolower($0)}') \
    -node=node2 \
    -data-dir=/tmp/data-dir -client 0.0.0.0 -ui
```

- node3(server)

```
docker run -d --name=node3 --restart=always \
    -e 'CONSUL_LOCAL_CONFIG={"skip_leave_on_interrupt":true}' \
    -p 10300:8300 \
    -p 10301:8301 \
    -p 10301:8301/udp \
    -p 10302:8302 \
    -p 10302:8302/udp \
    -p 10400:8400 \
    -p 10500:8500 \
    -p 10600:8600 \
    -h node3 \
    consul agent -server -bind=0.0.0.0 \
    -join=8.131.87.88 -node-id=$(uuidgen | awk '{print tolower($0)}') \
    -node=node3 \
    -data-dir=/tmp/data-dir -client 0.0.0.0 -ui
```

- node4(client)

```
docker run -d --name=node4 --restart=always \
    -e 'CONSUL_LOCAL_CONFIG={"leave_on_terminate":true}' \
    -p 11300:8300 \
    -p 11301:8301 \
    -p 11301:8301/udp \
    -p 11302:8302 \
    -p 11302:8302/udp \
    -p 11400:8400 \
    -p 11500:8500 \
    -p 11600:8600 \
    -h node4 \
    consul agent -bind=0.0.0.0 -retry-join=8.131.87.88 \
    -node-id=$(uuidgen | awk '{print tolower($0)}') \
    -node=node4 -client 0.0.0.0 -ui
```

- node5(client)

```
docker run -d --name=node5 --restart=always \
    -e 'CONSUL_LOCAL_CONFIG={"leave_on_terminate":true}' \
    -p 12300:8300 \
    -p 12301:8301 \
    -p 12301:8301/udp \
    -p 12302:8302 \
    -p 12302:8302/udp \
    -p 12400:8400 \
    -p 12500:8500 \
    -p 12600:8600 \
    -h node5 \
    consul agent -bind=0.0.0.0 -retry-join=8.131.87.88 \
    -node-id=$(uuidgen | awk '{print tolower($0)}') \
    -node=node5 -client 0.0.0.0 -ui
```

- 集群

```
docker run -d -p 8090:80 -v /project/ConsulWithGateway/consulnginx:/var/log/nginx/ -v /project/ConsulWithGateway/consulnginx/consulnginx.conf:/etc/nginx/nginx.conf --name consulnginx nginx
```

## 状态查看

- 查看节点

```
docker exec -t node1 consul members
```

- 查看主从信息

```
docker exec -t node1 consul operator raft list-peers
```
