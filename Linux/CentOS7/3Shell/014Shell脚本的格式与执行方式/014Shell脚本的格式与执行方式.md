# Shell 脚本
* UNIX 的哲学：一条命令只做一件事
* 为了组合命令和多次执行，使用脚本文件来保存需要执行的命令
* 赋予该文件执行权限（chmod u+rx filename）

```
cd /var/ ; ls ; pwd ; du -sh ; du -sh *

cd
vim helloworld.sh
chmod u+x helloworld.sh
bash hellowrold.sh
```

## bash声明
```
#!/bin/bash
```

## 其他运行脚本的方法
```
# 使用当前默认解释器解释
./helloworld.sh
```

## 注释
```
# demo
# 也可以写在命令后面，但是不推荐（阅读不方便）
```

# 标准的 Shell 脚本要包含哪些元素
* Sha-Bang（#!/bin/bash）
* 命令
* “#” 号开头的注释
* chmod u+rx filename 可执行权限
* 执行命令
  * bash ./filename.sh （可以不用赋予执行权限，执行完毕后返回到当前目录）
  > 在当前终端下产生一个新的叫做bash的子进程，用这个子进程运行脚本
  * ./filename.sh （如果不赋予执行权限会提示权限不够，执行完毕后返回到当前目录）
  > 也是产生子进程，由子进程运行脚本
  * source ./filename.sh （可以不用赋予执行权限，执行完毕后不返回到当前目录）
  > 在当前进程里执行脚本
  * . filename.sh （可以不用赋予执行权限，执行完毕后不返回到当前目录）
  > 在当前进程里执行脚本

  ## 内建命令和外部命令的区别
  * 内建命令不需要创建子进程（ex cd，pwd）
  * 内建命令对当前 Shell 生效 （source ./filename.sh，. filename.sh ）
