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
