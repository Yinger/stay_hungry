# 基本架构

## 基本操作集合

- PUT：新写入或更新一个 key-value 对；
- GET：根据一个 key 读取相应的 value 值；
- DELETE：根据一个 key 删除整个 key-value 对。
- SCAN: 根据一段 key 的范围返回相应的 value 值

有些键值数据库的新写 / 更新操作叫 SET

## 键值对保存在内存还是外存

- 保存在内存的好处是读写很快，毕竟内存的访问速度一般都在百 ns 级别。但是，潜在的风险是一旦掉电，所有的数据都会丢失。
- 保存在外存，虽然可以避免数据丢失，但是受限于磁盘的慢速读写（通常在几 ms 级别），键值数据库的整体性能会被拉低。

## 采用什么访问模式

- 通过函数库调用的方式供外部应用使用
- 通过网络框架以 Socket 通信的形式对外提供键值对操作

## 如何定位键值对的位置

索引的作用是让键值数据库根据 key 找到相应 value 的存储位置，进而执行操作

Memcached 和 Redis 采用哈希表作为 key-value 索引，而 RocksDB 则采用跳表作为内存中 key-value 的索引

## 不同操作的具体逻辑是怎样的

- 对于 GET/SCAN 操作而言，此时根据 value 的存储位置返回 value 值即可；
- 对于 PUT 一个新的键值对数据而言，SimpleKV 需要为该键值对分配内存空间；
- 对于 DELETE 操作，SimpleKV 需要删除键值对，并释放相应的内存空间，这个过程由分配器完成。

## 如何实现重启后快速提供服务

采用了常用的内存分配器 glibc 的 malloc 和 free

# 数据结构

## 数据类型

- String
- List
- Hash
- Set
- Sorted Set

## 底层数据结构

- 简单动态字符串 (String) O(1)
- 双向链表 (List) O(n)
- 压缩列表 (List,Hash) O(n)
- 哈希表 (Hash,Set) O(1)
- 跳表 (Sorted Set) O(logN)
- 整数数组 (Set) O(n)
