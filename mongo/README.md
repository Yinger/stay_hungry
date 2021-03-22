# mongo

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
