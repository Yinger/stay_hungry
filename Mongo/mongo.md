# Installation

## Linux 上安装 MongoDB

```
tar -xvf ./mongodb-linux-x86_64-rhel80-4.4.4.tgz
export PATH=$PATH:/data/mongodb-linux-x86_64-rhel80-4.4.4/bin
mongod --dbpath /data/db --port 27017 --logpath /data/db/mongod.log --fork
```

## 使用 Atlas 免费账号

- https://cloud.mongodb.com/

```
mongo "mongodb+srv://cmongo.x8nbf.mongodb.net/firstDB" --username sa
```

## 导入样本数据

```
tar -xvf dump.tar.gz
mongorestore -h localhost:27017
```

## 下载并安装 MongoDB Compass

- https://www.mongodb.com/products/compass

## docker

```
docker run -it -v data:/data/db -p 27018:27017 --name mongodb -d mongo --auth
docker cp /data/dump/. 8600a1a820ff:/data/dump
docker exec -it mongodb bash
```

# CRUD

## insert

```
db.<集合>.insertOne(<JSON对象>)
db.<集合>.insertMany([<JSON 1>, <JSON 2>, ...<JSON n>])

db.fruit.insertOne({name: "apple"})
db.fruit.insertMany([
    {name: "apple"},
    {name: "pear"},
    {name: "orange"}
])
```

## find

- find 是 MongoDB 中查询数据的基本指令，相当于 SQL 中的 SELECT
- find 返回的是游标

```
db.movies.find( { "year" : 1975 } ) //单条件查询
db.movies.find( { "year" : 1989, "title" : "Batman" } ) //多条件and查询
db.movies.find( { $and : [ {"title" : "Batman"}, { "category" : "action" }] } ) // and的另一种形式
db.movies.find( { $or: [{"year" : 1989}, {"title" : "Batman"}] } ) //多条件or查询
db.movies.find( { "title" : /^B/} ) //按正则表达式查找
```

### 查询条件对照表

| SQL    | MQL            |
| ------ | -------------- |
| a=1    | {a: 1}         |
| a <> 1 | {a: {$ne: 1}}  |
| a>1    | {a: {$gt: 1}}  |
| a >= 1 | {a: {$gte: 1}} |
| a<1    | {a: {$lt: 1}}  |
| a <= 1 | {a: {$lte: 1}} |

### 查询逻辑对照表

| SQL             | MQL                                    |
| --------------- | -------------------------------------- |
| a = 1 AND b = 1 | {a: 1, b: 1}或{$and: [{a: 1}, {b: 1}]} |
| a = 1 OR b = 1  | {$or: [{a: 1}, {b: 1}]}                |
| a IS NULL       | {a: {$exists: false}}                  |
| a IN (1, 2, 3)  | {a: {$in: [1, 2, 3]}}                  |

### 查询逻辑运算符

- $lt: 存在并小于
- $lte: 存在并小于等于
- $gt: 存在并大于
- $gte: 存在并大于等于
- $ne: 不存在或存在但不等于
- $in: 存在并在指定数组中
- $nin: 不存在或不在指定数组中
- $or: 匹配两个或多个条件中的一个
- $and: 匹配全部条件

### 使用 find 搜索子文档

find 支持使用“field.sub_field”的形式查询子文档。假设有一个文档:

```
db.fruit.insertOne({
        name: "apple",
        from: {
            country: "China",
            province: "Guangdon"
} })
```

考虑以下查询的意义:

```
db.fruit.find( { "from.country" : "China" } )
db.fruit.find( { "from" : {country: "China"} } )
```

### 使用 find 搜索数组

find 支持对数组中的元素进行搜索。假设有一个文档:

```
db.fruit.insert([
      { "name" : "Apple", color: ["red", "green" ] },
      { "name" : "Mango", color: ["yellow", "green"] }
])
```

考虑以下查询的意义:

```
db.fruit.find({color: "red"})
db.fruit.find({$or: [{color: "red"}, {color: "yellow"}]} )
```

### 使用 find 搜索数组中的对象

考虑以下文档，在其中搜索

```
db.movies.insertOne( {
          "title" : "Raiders of the Lost Ark",
          "filming_locations" : [
      "USA" },{ "city" : "Los Angeles", "state" : "CA", "country" :
              { "city" : "Rome", "state" : "Lazio", "country" : "Italy" },
              { "city" : "Florence", "state" : "SC", "country" : "USA" }
          ]
})
```

```
// 查找城市是 Rome 的记录
db.movies.find({"filming_locations.city": "Rome"})
```

### 使用 find 搜索数组中的对象

在数组中搜索子对象的多个字段时，如果使用 $elemMatch，它表示必须是同一个 子对象满足多个条件。考虑以下两个查询:

```
db.getCollection('movies').find({
       "filming_locations.city": "Rome",
       "filming_locations.country": "USA"
      })
      db.getCollection('movies').find({
       "filming_locations": {
       $elemMatch:{"city":"Rome", "country": "USA"}
} })
```

### 控制 find 返回的字段

- find 可以指定只返回指定的字段
- \_id 字段必须明确指明不返回，否则默认返回
- 在 MongoDB 中我们称这为投影(projection)

```
//不返回_id 返回title
db.movies.find({"category": "action"},{"_id":0, title:1})
```

## remove

- remove 命令需要配合查询条件使用
- 匹配查询条件的的文档会被删除
- 指定一个空文档条件会删除所有文档

```
db.testcol.remove( { a : 1 } ) // 删除a 等于1的记录
db.testcol.remove( { a : { $lt : 5 } } ) // 删除a 小于5的记录
db.testcol.remove( { } ) // 删除所有记录
db.testcol.remove() //报错
```

## update

### 使用 update 更新文档

- Update 操作执行格式:db.<集合>.update(<查询条件>, <更新字段>)

```
db.fruit.insertMany([
        {name: "apple"},
        {name: "pear"},
        {name: "orange"}
])
```

```
// 查询 name 为 apple 的记录,将找到记录的 from 设置为 China
db.fruit.updateOne({name: "apple"}, {$set: {from: "China"}})
```

- 使用 updateOne 表示无论条件匹配多少条记录，始终只更新第一条
- 使用 updateMany 表示条件匹配多少条就更新多少条
- updateOne/updateMany 方法要求更新条件部分必须具有以下之一，否则将报错
  - $set/$unset
  - $push/$pushAll/$pop
  - $pull/$pullAll
  - $addToSet

```
// 报错
db.fruit.updateOne({name: "apple"}, {from: "China"})
```

### 使用 update 更新数组

- $push: 增加一个对象到数组底部
- $pushAll: 增加多个对象到数组底部
- $pop: 从数组底部删除一个对象
- $pull: 如果匹配指定的值，从数组中删除相应的对象
- $pullAll: 如果匹配任意的值，从数据中删除相应的对象
- $addToSet: 如果不存在则增加一个值到数组

## drop

### 使用 drop 删除集合

- 使用 db.<集合>.drop() 来删除一个集合
- 集合中的全部文档都会被删除
- 集合相关的索引也会被删除

```
db.colToBeDropped.drop()
```

### 使用 dropDatabase 删除数据库

- 使用 db.dropDatabase() 来删除数据库
- 数据库相应文件也会被删除，磁盘空间将被释放

```
 use tempDB
 db.dropDatabase()
 show collections // No collections
 show dbs // The db is gone
```

# 文档模型设计

## 一. 建立基础文档模型

### 1-1 关系建模: portraits

- 基本原则:
  - 一对一关系以内嵌为主
  - 作为子文档形式 或者直接在顶级
  - 不涉及到数据冗余
- 例外情况
  - 如果内嵌后导致文档大小超过 16MB

```json
name: "Rin",
company: "ABC",
title: "CTO",
portraits: {
    mimetype: xxx,
    data: xxxx }
```

### 1-N 关系建模: Addresses

- 基本原则:
  - 一对多关系同样以内嵌为主
  - 用数组来表示一对多
  - 不涉及到数据冗余
- 例外情况
  - 如果内嵌后导致文档大小超过 16MB
  - 数组长度太大(数万或更多)
  - 数组长度不确定

```json
name: "Rin",
company: "ABC",
title: "CTO",
portraits: {
    mimetype: xxx,
    data: xxxx
    },
addresses: [
     { type: home, ... },
     { type: work, ... }
]
```

### N-N 关系建模:内嵌数组模式:groups

- 基本原则:
  - 不需要映射表
  - 一般用内嵌数组来表示一对多
  - 通过冗余来实现 N-N
- 例外情况
  - 如果内嵌后导致文档大小超过 16MB
  - 数组长度太大(数万或更多)
  - 数组长度不确定

```json
name: "Rin",
company: "ABC",
title: "CTO",
portraits: {
    mimetype: xxx,
    data: xxxx
    },
addresses: [
     { type: home, ... },
     { type: work, ... }
],
groups: [
      {name:  ”Friends” },
      {name:  ”Surfers” },
   ]
```

### 小结

- 90:10 规则: 大部分时候你会使用内嵌来表示 1-1，1-N，N-N
- 内嵌类似于预先聚合(关联)
- 内嵌后对读操作通常有优势(减少关联)

## 二. 根据读写工况细化

### 例： 一个分组信息的改动意味着 百万级的 DB 操作

- 解决方案: Group 使用单独的集合

  - 类似于关系型设计
  - 用 id 或者唯一键关联
  - 使用 `$lookup` 来提供一次查询多表 的能力(类似关联)

  ```json
    name: "Rin",
    company: "ABC",
    title: "CTO",
    portraits: {
        mimetype: xxx,
        data: xxxx
        },
    addresses: [
        { type: home, ... },
        { type: work, ... }
    ],
    group_ids: [1,2,3...]
  ```

  ```json
  group_id:1,
  group_name:"student"
  ```

  ```nosql
  db.contacts.aggregate([
   {
        $lookup:
            {
                from: "groups",
                localField: "group_ids",
                foreignField: "group_id",
                as: "groups"
            }
    }])
  ```

### 联系人的头像: 引用模式

- 头像使用高保真，大小在 5MB- 10MB
- 头像一旦上传，一个月不可更换
- 基础信息查询(不含头像)和 头 像查询的比例为 9 :1
- 建议: 使用引用方式，把头像数 据放到另外一个集合，可以显著提 升 90% 的查询效率

Contact:

```json
name: "Rin",
company: "ABC",
title: "CTO",
portraits_id:123,
addresses: [
    { type: home, ... },
    { type: work, ... }
]
```

Contact_Portrait:

```json
_id: 123,
mimetype: “xxx”,
data: ”xxxx...”
```

### 小结：什么时候该使用引用方式

- 内嵌文档太大，数 MB 或者超过 16MB
- 内嵌文档或数组元素会频繁修改
- 内嵌数组元素会持续增长并且没有封顶

## MongoDB 引用设计的限制

- MongoDB 对使用引用的集合之间并无主外键检查
- MongoDB 使用聚合框架的 `$lookup` 来模仿关联查询
- `$lookup` 只支持 `left outer join`
- `$lookup` 的关联目标(from)不能是分片表

# 文档模型设计之模式套用

## 分桶设计模式

| 场景                              | 痛点                       | 设计模式的方案及优点                                                                       |
| --------------------------------- | -------------------------- | ------------------------------------------------------------------------------------------ |
| 时序数据,物联网,智慧城市,智慧交通 | 数据点采集频繁，数据量太多 | 利用文档内嵌数组，将一个时间段的数据聚合到一个文档里。大量减少文档数量大量减少索引占用空间 |

## 列转行

| 场景                                                           | 痛点                                                      | 设计模式的方案及优点                 |
| -------------------------------------------------------------- | --------------------------------------------------------- | ------------------------------------ |
| 产品属性 ‘color’, ‘size’, ‘dimensions’, ... 多语言(多国家)属性 | 文档中有很多类似的字段.会用于组合查询搜索，需要建很多索引 | 转化为数组，一个索引解决所有查询问题 |

## 版本字段

## 近似计算

## 预聚合

# 事务开发

## 写操作事务

### `writeConcern`

`writeConcern` 决定一个写操作落到多少个节点上才算成功。`writeConcern` 的取值包括：

- 0：发起写操作，不关心是否成功
- 1~集群最大数据节点数：写操作需要被复制到指定节点数才算成功
- majority：写操作需要被复制到大多数节点上才算成功
  发起写操作的程序将阻塞到写操作到达指定的节点数为止

| 默认行为                                                  | w: “majority”                                               | w: “all”                                          |
| --------------------------------------------------------- | ----------------------------------------------------------- | ------------------------------------------------- |
| 复制集不作任何特别设定                                    | 大多数节点确认模式                                          | 全部节点确认模式                                  |
| ![default](/Mongo/Documents/pic/writeConcern-default.png) | ![majority](/Mongo/Documents/pic/writeConcern-majority.png) | ![all](/Mongo/Documents/pic/writeConcern-all.png) |

### `j:true`

`writeConcern` 可以决定写操作到达多少个节点才算成功，journal 则定义如何才算成功。取值包括

- true: 写操作落到 journal 文件中才算成功
- false: 写操作到达内存即算作成功。

### writeConcern 实验

```
在复制集测试writeConcern参数
db.test.insert( {count: 1}, {writeConcern: {w: "majority"}})
db.test.insert( {count: 1}, {writeConcern: {w: 3 }})
db.test.insert( {count: 1}, {writeConcern: {w: 4 }})
配置延迟节点，模拟网络延迟（复制延迟）
conf=rs.conf()
conf.members[2].slaveDelay = 5
conf.members[2].priority = 0
rs.reconfig(conf)
观察复制延迟下的写入，以及timeout参数
db.test.insert( {count: 1}, {writeConcern: {w: 3}})
db.test.insert( {count: 1}, {writeConcern: {w: 3, wtimeout:3000 }})
```

### 注意事项

- 虽然多于半数的 `writeConcern` 都是安全的，但通常只会设置 `majority`，因为这是等待写入延迟时间最短的选择
- 不要设置 `writeConcern` 等于总节点数，因为一旦有一个节点故障，所有写操作都将失败
- `writeConcern` 虽然会增加写操作延迟时间，但并不会显著增加集群压力，因此无论是否等待，写操作最终都会复制到所有节点上。设置 `writeConcern` 只是让写操作等待复制后再返回而已
- 应对重要数据应用 `{w: “majority”}`，普通数据可以应用 `{w: 1}` 以确保最佳性能

## 读操作事务

在读取数据的过程中我们需要关注以下两个问题

- 从哪里读 (`readPreference`)
- 什么样的数据可以读 (`readConcern`)

### readPreference

readPreference 决定使用哪一个节点来满足正在发起的读请求。可选值包括

- `primary`: 只选择主节点
- `primaryPreferred`：优先选择主节点，如果不可用则选择从节点
- `secondary`：只选择从节点
- `secondaryPreferred`：优先选择从节点，如果从节点不可用则选择主节点
- `nearest`：选择最近的节点

场景举例

- 用户下订单后马上将用户转到订单详情页——primary/primaryPreferred。因为此时从节点可能还没复制到新订单；
- 用户查询自己下过的订单——secondary/secondaryPreferred。查询历史订单对时效性通常没有太高要求；
- 生成报表——secondary。报表对时效性要求不高，但资源需求大，可以在从节点单独处理，避免对线上用户造成影响；
- 将用户上传的图片分发到全世界，让各地用户能够就近读取——nearest。每个地区的应用选择最近的节点读取数据。

#### readPreference 与 Tag

`readPreference` 只能控制使用一类节点。`Tag` 则可以将节点选择控制到一个或几个节点

- `{purpose: "online"}`
- `{purpose: "analyse"}`

#### readPreference 配置

- 通过 MongoDB 的连接串参数
  - `mongodb://host1:27107,host2:27107,host3:27017/?replicaSet=rs&readPreference=secondary`
- 通过 MongoDB 驱动程序 API
  - `MongoCollection.withReadPreference(ReadPreference readPref)`
- Mongo Shell
  - `db.collection.find({}).readPref( “secondary” )`

#### 实验: 从节点读

```
主节点写入 {x:1}, 观察该条数据在各个节点均可见
在两个从节点分别执行 db.fsyncLock() 来锁定写入（同步）
主节点写入 {x:2}
  db.test.find({a: 123})
  db.test.find({a: 123}).readPref(“secondary”)
解除从节点锁定 db.fsyncUnlock()
  db.test.find({a: 123}).readPref(“secondary”)
```

#### 注意事项

- 指定 readPreference 时也应注意高可用问题。例如将 readPreference 指定 primary，则发生故障转移不存在 primary 期间将没有节点可读。如果业务允许，则应选择 primaryPreferred
- 使用 Tag 时也会遇到同样的问题，如果只有一个节点拥有一个特定 Tag，则在这个节点失效时将无节点可读。这在有时候是期望的结果，有时候不是。例如
  - 如果报表使用的节点失效，即使不生成报表，通常也不希望将报表负载转移到其他节点上，此时只有一个节点有报表 Tag 是合理的选择
  - 如果线上节点失效，通常希望有替代节点，所以应该保持多个节点有同样的 Tag
- Tag 有时需要与优先级、选举权综合考虑。例如做报表的节点通常不会希望它成为主节点，则优先级应为 0

### readConcern

在 readPreference 选择了指定的节点后，readConcern 决定这个节点上的数据哪些是可读的，类似于关系数据库的隔离级别。可选值包括

- `available`：读取所有可用的数据
- `local`：读取所有可用且属于当前分片的数据 (默认设置)
- `majority`：读取在大多数节点上提交完成的数据 (数据读一致性的充分保证)
- `linearizable`：可线性化读取文档 (可能导致非常慢的读，因此总是建议配合使用 maxTimeMS)
- `snapshot`：读取最近快照中的数据 (最高隔离级别，接近于 Seriazable)

## 多文档事务

- 事务默认必须在 60 秒（可调）内完成，否则将被取消
- 涉及事务的分片不能使用仲裁节点
- 事务会影响 chunk 迁移效率。正在迁移的 chunk 也可能造成事务提交失败（重试即可）
- 多文档事务中的读操作必须使用主节点读
- readConcern 只应该在事务级别设置，不能设置在每次读写操作上

# Change Stream

Change Stream 是 MongoDB 用于实现变更追踪的解决方案，类似于关系数据库的触发器，但原理不完全相同
||Change Stream|触发器|
|--|--|--|
|触发方式|异步|同步（事务保证）|
|触发位置|应用回调事件|数据库触发器|
|触发次数|每个订阅事件的客户端|1 次（触发器）|
|故障恢复|从上次断点重新触发|事务回滚|

Change Stream 是基于 oplog 实现的。它在 oplog 上开启一个 tailable cursor 来追踪所有复制集上的变更操作，最终调用应用中定义的回调函数。被追踪的变更事件主要包括

- insert/update/delete：插入、更新、删除
- drop：集合被删除
- rename：集合被重命名
- dropDatabase：数据库被删除
- invalidate：drop/rename/dropDatabase 将导致 invalidate 被触发,并关闭 change stream

### Change Stream 与可重复读

Change Stream 只推送已经在大多数节点上提交的变更操作。即“可重复读”的变更。这个验证是通过`{readConcern: “majority”}`实现的。因此

- 未开启 majority readConcern 的集群无法使用 Change Stream
- 当集群无法满足 `{w: “majority”}` 时，不会触发 Change Stream（例如 PSA 架构中的 S 因故障宕机）

### Change Stream 变更过滤

如果只对某些类型的变更事件感兴趣，可以使用使用聚合管道的过滤步骤过滤事件

```
var cs = db.collection.watch([{
  $match: {
    operationType: {
      $in: ['insert', 'delete']
    }
  }
}])
```

### Change Stream 故障恢复

想要从上次中断的地方继续获取变更流，只需要保留上次变更通知中的 `_id` 即可

```
var cs = db.collection.watch([], {resumeAfter: <_id>})
```

即可从上一条通知中断处继续获取后续的变更通知

### Change Stream 使用场景

- 跨集群的变更复制——在源集群中订阅 Change Stream，一旦得到任何变更立即写入目标集群。
- 微服务联动——当一个微服务变更数据库时，其他微服务得到通知并做出相应的变更
- 其他任何需要系统联动的场景

### 注意事项

- Change Stream 依赖于 oplog，因此中断时间不可超过 oplog 回收的最大时间窗
- 在执行 update 操作时，如果只更新了部分数据，那么 Change Stream 通知的也是增量部分
- 同理，删除数据时通知的仅是删除数据的 `_id`

# MongoDB 开发最佳实践

### 连接到 MongoDB

- 关于驱动程序：总是选择与所用之 MongoDB 相兼容的驱动程序。这可以很容易地从[驱动兼容对照表](https://docs.mongodb.com/drivers/driver-compatibility-reference/)中查到
- 关于连接对象 MongoClient：使用 MongoClient 对象连接到 MongoDB 实例时总是应该保证它单例，并且在整个生命周期中都从它获取其他操作对象
- 关于连接字符串：连接字符串中可以配置大部分连接选项，建议总是在连接字符串中配置这些选项

  ```
  // 连接到复制集
  mongodb://节点1,节点2,节点3…/database?[options]
  // 连接到分片集
  mongodb://mongos1,mongos2,mongos3…/database?[options]
  ```

### 常见连接字符串参数

- maxPoolSize
  - 连接池大小
- Max Wait Time
  - 建议设置，自动杀掉太慢的查询
- Write Concern
  - 建议 majority 保证数据安全
- Read Concern
  - 对于数据一致性要求高的场景适当使用

### 连接字符串节点和地址

- 无论对于复制集或分片集，连接字符串中都应尽可能多地提供节点地址，建议全部列出；
  - 复制集利用这些地址可以更有效地发现集群成员；
  - 分片集利用这些地址可以更有效地分散负载；
- 连接字符串中尽可能使用与复制集内部配置相同的域名或 IP

### 使用域名连接集群

在配置集群时使用域名可以为集群变更时提供一层额外的保护。例如需要将集群整体迁移到新网段，直接修改域名解析即可

另外，MongoDB 提供的 `mongodb+srv://` 协议可以提供额外一层的保护。该协议允许通过域名解析得到所有 mongos 或节点的地址，而不是写在连接字符串中

### 不要在 mongos 前面使用负载均衡

### 游标使用

如果一个游标已经遍历完，则会自动关闭；如果没有遍历完，则需要手动调用 `close()`方法，否则该游标将在服务器上存在 10 分钟（默认值）后超时释放，造成不必要的资源浪费。

但是，如果不能遍历完一个游标，通常意味着查询条件太宽泛，更应该考虑的问题是
如何将条件收紧

### 关于查询及索引

- 每一个查询都必须要有对应的索引
- 尽量使用覆盖索引 Covered Indexes （可以避免读数据文件）
- 使用 projection 来减少返回到客户端的的文档的内容

### 关于写入

- 在 update 语句里只包括需要更新的字段
- 尽可能使用批量插入来提升写入性能
- 使用 TTL 自动过期日志类型的数

### 关于文档结构

- 防止使用太长的字段名（浪费空间）
- 防止使用太深的数组嵌套（超过 2 层操作比较复杂）
- 不使用中文，标点符号等非拉丁字母作为字段名

### 处理分页问题 – 避免使用 count

### 处理分页问题 – 巧分页

避免使用 skip/limit 形式的分页，特别是数据量大的时候；替代方案：使用查询条件+唯一排序条件；

- 例如：
  ```
  第一页：db.posts.find({}).sort({_id: 1}).limit(20);
  第二页：db.posts.find({_id: {$gt: <第一页最后一个_id>}}).sort({_id: 1}).limit(20);
  第三页：db.posts.find({_id: {$gt: <第二页最后一个_id>}}).sort({_id: 1}).limit(20);
  ...
  ```

### 关于事务

使用事务的原则：

- 无论何时，事务的使用总是能避免则避免；
- 模型设计先于事务，尽可能用模型设计规避事务；
- 不要使用过大的事务（尽量控制在 1000 个文档更新以内）；
- 当必须使用事务时，尽可能让涉及事务的文档分布在同一个分片上，这将有效地提高效率；
