# 实验：搭建复制集

## 准备工作

- 安装 Mongo

## 搭建复制集

### 创建数据目录

MongoDB 启动时将使用一个数据目录存放所有数据文件。我们将为 3 个复制集节点创建各自的数据目录。

Linux/MacOS:

```bash
mkdir -p /data/db{1,2,3}
```

### 准备配置文件

- `/data/db1/mongod.conf`
- `/data/db2/mongod.conf`
- `/data/db3/mongod.conf`

```yaml
# /data/db1/mongod.conf
systemLog:
  destination: file
  path: /data/db1/mongod.log # 日志文件路径
  logAppend: true
storage:
  dbPath: /data/db1 # 数据目录
net:
  bindIp: 0.0.0.0
  port: 28017 # 端口
replication:
  replSetName: rs0
processManagement:
  fork: true
```

### 执行进程

```bash
mongod -f /data/db1/mongod.conf
mongod -f /data/db2/mongod.conf
mongod -f /data/db3/mongod.conf
```

```
ps -f | grep mongod
```

### 配置复制集

```bash
mongo --port 28017
rs.initiate()

rs.add("rin:28018")
rs.add("rin:28019")

# 查看复制集状态
rs.status()
```

验证：

```bash
mongo rin:28018
db.test.find() # error
rs.secondaryOk()
db.test.find()
...
```
