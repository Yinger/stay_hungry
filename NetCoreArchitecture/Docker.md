# 容器

容器本身没有价值，有价值的是“容器编排”。
容器，其实是一种特殊的进程而已。

## Linux 容器的具体实现方式

一个“容器”，实际上是一个由 Linux Namespace、Linux Cgroups 和 rootfs 三种技术构建出来的进程的隔离环境。

1. Namespace 的作用是“隔离”，它让应用进程只能看到该 Namespace 内的“世界”；修改进程视图的主要方法
2. Cgroups 的作用是“限制”，它给这个“世界”围上了一圈看不见的墙。制造约束的主要手段
3. 一组联合挂载在 /var/lib/docker/aufs/mnt 上的 rootfs，这一部分我们称为“容器镜像”（Container Image），是容器的静态视图；一个由 Namespace+Cgroups 构成的隔离环境，这一部分我们称为“容器运行时”（Container Runtime），是容器的动态视图。

### namespace 机制

在容器中看到只有自己一个进程
但是这个进程其实是在宿主机上的 只不过是一个独立的 namespace
定义了 5 个命名空间结构体，多个进程可以使用同一个 namespace

1. UTS： 运行内核的名称、版本、底层体系结构类型等信息（UNIX Timesharing System）
2. IPC： 与进程间通信（IPC)有关
3. MNT： 已经装载的文件系统的视图 Mount Namespace，用于让被隔离进程只看到当前 Namespace 里的挂载点信息；
4. PID：有关进程 ID 的信息
5. NET：网络相关的命名空间参数 Network Namespace，用于让被隔离进程看到当前 Namespace 里的网络设备和配置。
   在 Linux 内核中，有很多资源和对象是不能被 Namespace 化的，最典型的例子就是：时间。

### cgroup 技术

Linux Cgroups 就是 Linux 内核中用来为进程设置资源限制的一个重要功能。
Linux Cgroups 的全称是 Linux Control Group。它最主要的作用，就是限制一个进程组能够使用的资源上限，包括 CPU、内存、磁盘、网络带宽等等。
此外，Cgroups 还能够对进程进行优先级设置、审计，以及将进程挂起和恢复等操作
一个正在运行的 Docker 容器，其实就是一个启用了多个 Linux Namespace 的应用进程，而这个进程能够使用的资源量，则受 Cgroups 配置的限制。这也是容器技术中一个非常重要的概念，即：容器是一个“单进程”模型。

# 基本操作（CentOS）

## 安装

https://docs.docker.com/engine/install/centos/

## 卸载

- 查询 docker 安装过的包

```
yum list installed | grep docker
```

- 删除安装包

```
yum remove docker-ce.x86_64 docker-ce-cli.x86_64 -y
```

- 删除镜像/容器等

```
rm -rf /var/lib/docker
```

## 启动、设置开启开机启动

```
# service docker start
sudo systemctl start docker
sudo systemctl enable docker
```

## 查看状态

```
systemctl status docker
```

# 常用命令

```
docker -v / version /info         查看docker版本信息
docker images                     查看所有镜像
docker ps                         查看所有正在运行的容器
docker ps -a                      查看所有容器
docker rmi 镜像ID                 删除某个镜像
docker stop 容器ID                停止某个容器
docker start 容器ID               启动某个容器
docker restart 容器ID             重启某个容器
docker rm 容器ID                  删除某个容器
docker rm -f 容器ID               删除正在运行的容器
docker rmi $(docker images -q)    删除所有镜像
docker rm $(docker ps -a -q)      删除所有容器
docker logs 容器ID                查看容器日志信息找原因
netstat -tanlp                    查看端口
kill 端口                         杀死某个端口
docker search mysql               镜像检索
docker pull mysql                 镜像下载

docker stop $(docker ps -q) & docker rm $(docker ps -aq) #停用並刪除所有容器
docker rmi -f $(docker images -qa) #删除所有镜像
docker rmi $(docker images -f "dangling=true" -q) #删除不需要的中间镜像  (<none> 中间镜像（intermediate images）)
docker exec -it CONTAINERID /bin/bash  # 进入到容器中操作
exit  # 退出
docker ps -a | grep Exit | cut -d ' ' -f 1 | xargs sudo docker rm #remove all Exited containers

docker images|grep hello

for i in `docker ps -a|grep -i exit|awk '{print $1}'`;do docker rm -f $i; done
```

# 构建 .net core 镜像

## 方式一：基于 Dockerfile

- 项目里添加 Docker 支持
- 将项目 **源代码** ftp 至服务器 （/project）
- 服务器 cd 至项目目录 `cd /project/HelloDocker`
- `docker build -t hello-docker -f Dockerfile .`
- 创建一个新的容器并运行 `docker run -itd -p 9001:80 hello-docker`

## 方式二：docker 基础镜像 + 挂载文件

- 发布项目
- ftp 到服务器
- pull 基础镜像
  ```
  docker pull mcr.microsoft.com/dotnet/core/sdk
  docker pull microsoft/dotnet
  ```
- 挂载运行 `docker run -d -p 9002:80 -v /project/publish/HelloDocker:/app --workdir /app mcr.microsoft.com/dotnet/core/aspnet dotnet /app/HelloDocker.dll`

## 方式三：直接用 DockerHub 上的基础镜像启动

- DockerHub 拉取 Nginx 镜像
- 单独的 ngnix.conf (对应版本)
- 修改配置文件，挂载进去
- Run `docker run -d -p 9006:80 -v /project/enginx:/var/log/nginx/ -v /project/enginx/nginx.conf:/etc/nginx/nginx.conf --name netcore-nginx nginx`

```
<语法>
docker run [OPTIONS] IMAGE [COMMAND] [ARG...]
OPTIONS说明：

-a stdin: 指定标准输入输出内容类型，可选 STDIN/STDOUT/STDERR 三项；

-d: 后台运行容器，并返回容器ID；

-i: 以交互模式运行容器，通常与 -t 同时使用；

-p: 端口映射，格式为：主机(宿主)端口:容器端口

-t: 为容器重新分配一个伪输入终端，通常与 -i 同时使用；

--name="nginx-lb": 为容器指定一个名称；

--dns 8.8.8.8: 指定容器使用的DNS服务器，默认和宿主一致；

--dns-search example.com: 指定容器DNS搜索域名，默认和宿主一致；

-h "mars": 指定容器的hostname；

-e username="ritchie": 设置环境变量；

--env-file=[]: 从指定文件读入环境变量；

--cpuset="0-2" or --cpuset="0,1,2": 绑定容器到指定CPU运行；

-m :设置容器使用内存最大值；

--net="bridge": 指定容器的网络连接类型，支持 bridge/host/none/container: 四种类型；

--link=[]: 添加链接到另一个容器；

--expose=[]: 开放一个端口或一组端口；

--rm: 运行完成后自己退出

bin/sleep 300
ps aux|grep sleep|grep -v grep
```

## docker hub

```
docker login docker.io
cat /root/.docker/config.json
echo "cmlu..."|base64 -d
```

## 镜像管理

registry_name/repository_name/image_name:tag_name

```
docker search alpine
docker pull alpine
docker pull alpine:3.10.1
docker pull docker.io/library/alpine:3.10.1

# 给镜像打标签
docker tag ...image_id... docker.io/yourdockerhubname/alpine:v3.10.3

# 推送镜像
docker push docker.io/yourdockerhubname/alpine:v3.10.3
```

## 导入导出

```
docker commit -p myalpine yourdockerhubname/alpine:v3.10.3_with1.txt
docker export containerid > myalpine_with_1.tar

docker same containerid > alpine:v3.10.3_with_1.tar
docker load < alpine\:v3.10.3_with_1.tar
```

## 日志

```
docker run 2>&1 >> /dev/null
docker logs -f containerid
```

# 重点操作

## 映射端口

- docker run -p 容器外端口:容器内端口

```
docker run --rm --name mynginx -d p81:80 imagename:version
netstat -luntp|grep 81
curl 127.0.0.1:81
```

## 挂载数据卷

- docker run -v 容器外目录:容器内目录

```
mkdir html
wget www.baidu.com -O index.html
docker run -d --rm --name nginx_with_baidu -d -p 82:80 -v /root/html:/usr/share/nginx/html imagename
docker inspect containerid|grep share
```

## 传递环境变量

- docker run -e 环境变零 key:环境变量 value

```
docker run --rm -e E_OPTS=abc -e C_OPTS=123 imagename printenv
```

## 容器内安装软件（工具）

- yum/apt-get/apt 等

```
docker exec -it containername /bin/bash

tee /etc/apt/source.list <<EOF
deb http://mirrors.163.com/debian/ jessie main non-free contrib
deb http://mirrors.163.com/debian/ jessie-updates main non-free contrib
EOF

apt -get update && apt -get install curl -y
curl -k https://www.baidu.com

docker commit -p containerid name:tag
docker push contianername:tag
```

# 容器的生命周期

- 检查本地是否存在镜像，如果不存在即从远端仓库检索
- 利用镜像启动容器
- 分配一个文件系统，并在只读的镜像层外挂载一层可读写层
- 从宿主机配置的网桥接口中桥接一个虚拟接口到容器
- 从地址池配置一个 ip 地址给容器
- 执行用户指定的指令
- 执行完毕后容器终止

# Dockerfile

> Dockerfile 是一个文本文件，其内包含了一条条的 指令(Instruction)，每一条指令构建一层，因此每一条指令的内容，就是描述该层应当如何构建。

## 4 组核心的 Dockerfile 指令

- USER/WORKDIR
- ADD/EXPOSE
- RUN/ENV
- CMD/ENTRYPOINT

## dockerfile 的编写

在一个空白目录中，建立一个文本文件，并命名为 Dockerfile：

```
$ mkdir mynginx
$ cd mynginx
$ touch Dockerfile
```

其内容为

```
FROM nginx
RUN echo '<h1>Hello, Docker!</h1>' > /usr/share/nginx/html/index.html
```

### 注意事项

由于 dockerfile 中每一个指令都会建立一层，每一个 RUN 的行为，会新建立一层，在其上执行这些命令，执行结束后，`commit` 这一层的修改，构成新的镜像。镜像是多层存储，**每一层的东西并不会在下一层被删除 s**，会一直跟随着镜像。因此镜像构建时，一定要确保每一层只添加真正需要添加的东西，任何无关的东西都应该清理掉。(安装包、缓存等)

Dockerfile 支持 Shell 类的行尾添加 `\` 的命令换行方式，以及行首 `#` 进行注释的格式。良好的格式，比如换行、缩进、注释等，会让维护、排障更为容易，这是一个比较好的习惯。

```
FROM debian:stretch
RUN buildDeps='gcc libc6-dev make wget' \
    && apt-get update \
    && apt-get install -y $buildDeps \
    && wget -O redis.tar.gz "http://download.redis.io/releases/redis-5.0.3.tar.gz" \
    && mkdir -p /usr/src/redis \
    && tar -xzf redis.tar.gz -C /usr/src/redis --strip-components=1 \
    && make -C /usr/src/redis \
    && make -C /usr/src/redis install \
    && rm -rf /var/lib/apt/lists/* \
    && rm redis.tar.gz \
    && rm -r /usr/src/redis \
    && apt-get purge -y --auto-remove $buildDeps
```

## 构建镜像

命令格式为 `docker build [选项] <上下文路径/URL/->`

在 Dockerfile 文件所在目录执行：

```
[root@supercomputer]# docker build -t nginx:v3 .
Sending build context to Docker daemon 2.048 kB
Step 1 : FROM nginx
 ---> e43d811ce2f4
Step 2 : RUN echo '<h1>Hello, Docker!</h1>' > /usr/share/nginx/html/index.html
 ---> Running in 9cdc27646c7b
 ---> 44aa4490ce2c
Removing intermediate container 9cdc27646c7b
Successfully built 44aa4490ce2c
```

在这里我们指定了最终镜像的名称 `-t nginx:v3`

### 上下文路径

docker build 命令最后有一个 `.`。`.` 表示当前目录，但是这里的当前目录指的并非是 dockerfile 所在的路径 `docker build -t nginx:v3 .` 中的这个 .，实际上是在指定上下文的目录，docker build 命令会将该目录下的内容打包交给 Docker 引擎以帮助构建镜像。
一般来说，应该会将 Dockerfile 置于一个空目录下，或者项目根目录下。如果该目录下没有所需文件，那么应该把所需文件复制一份过来。如果目录下有些东西确实不希望构建时传给 Docker 引擎，那么可以用 `.gitignore` 一样的语法写一个 `.dockerignore`，该文件是用于剔除不需要作为上下文传递给 Docker 引擎的。

## Dockerfile 指令详解

### FROM 指定基础镜像

`FROM` 就是指定 基础镜像，因此一个 Dockerfile 中 `FROM` 是必备的指令，并且 **必须是第一条指令**。
除了选择现有镜像为基础镜像外，Docker 还存在一个特殊的镜像，名为 `scratch`。这个镜像是虚拟的概念，并不实际存在，它表示一个空白的镜像。

```
FROM scratch
...
```

如果你以 scratch 为基础镜像的话, 意味着你不以任何镜像为基础, 接下来所写的指令将作为镜像第一层开始存在.
不以任何系统为基础, 直接将可执行文件复制进镜像的做法并不罕见, 比如 swarm, etcd. 对于 Linux 下静态编译的程序来说, 并不需要有操作系统提供运行时支持, 所需的一切库都已经在可执行文件里了, 因此直接 FROM scratch 会让镜像体积更加小巧. 使用 Go 语言 开发的应用很多会使用这种方式来制作镜像, 这也是为什么有人认为 Go 是特别适合容器微服务架构的语言的原因之一.

### RUN 执行命令

shell 格式：`RUN <命令>`，就像直接在命令行中输入的命令一样
注意要简化命令 避免多次使用 run 并且在最后清理安装包等

### COPY 复制文件

```
COPY [--chown=<user>:<group>] <源路径>... <目标路径>
COPY [--chown=<user>:<group>] ["<源路径1>",... "<目标路径>"]
```

`COPY` 指令将从构建上下文目录中 `<源路径>` 的文件/目录复制到新的一层的镜像内的 `<目标路径>` 位置。比如：

```
COPY package.json /usr/src/app/
COPY hom* /mydir/
COPY hom?.txt /mydir/
```

- `<源路径>` 可以是多个，甚至可以是通配符
- `<目标路径>` 可以是容器内的绝对路径，也可以是相对于工作目录的相对路径（工作目录可以用 WORKDIR 指令来指定）。目标路径不需要事先创建，如果目录不存在会在复制文件前先行创建缺失目录。

**注:** 使用 `COPY` 指令，源文件的各种元数据都会保留。比如读、写、执行权限、文件变更时间等。

### ADD 更高级的复制文件

`ADD` 指令和 `COPY` 的格式和性质基本一致。如果 `<源路径>` 为一个 `tar` 压缩文件的话，压缩格式为 `gzip`, `bzip2` 以及 `xz` 的情况下，`ADD` 指令将会自动解压缩这个压缩文件到 `<目标路径>` 去。

因此在 `COPY` 和 `ADD` 指令中选择的时候，可以遵循这样的原则，所有的文件复制均使用 `COPY` 指令，仅在需要自动解压缩的场合使用 `ADD`
在使用该指令的时候还可以加上 `--chown=<user>:<group>` 选项来改变文件的所属用户及所属组。

```
ADD --chown=55:mygroup files* /mydir/
ADD --chown=bin files* /mydir/
ADD --chown=1 files* /mydir/
ADD --chown=10:11 files* /mydir/
```

### CMD 容器启动命令

`CMD` 指令的格式和 `RUN` 相似，也是两种格式：

- shell 格式：`CMD <命令>`
- exec 格式：`CMD ["可执行文件", "参数1", "参数2"...]`
- 参数列表格式：`CMD ["参数1", "参数2"...]`。在指定了 `ENTRYPOINT` 指令后，用 `CMD` 指定具体的参数。

Docker 不是虚拟机，容器就是进程。既然是进程，那么在启动容器的时候，需要指定所运行的程序及参数。`CMD` 指令就是用于指定默认的容器主进程的启动命令的。

在运行时可以指定新的命令来替代镜像设置中的这个默认命令，比如，ubuntu 镜像默认的 CMD 是 /bin/bash，如果我们直接 docker run -it ubuntu 的话，会直接进入 bash。我们也可以在运行时指定运行别的命令，如 docker run -it ubuntu cat /etc/os-release。这就是用 cat /etc/os-release 命令替换了默认的 /bin/bash 命令了，输出了系统版本信息。
在指令格式上，一般推荐使用 exec 格式，这类格式在解析时会被解析为 JSON 数组，因此一定要使用双引号 "，而不要使用单引号。

如果使用 shell 格式的话，实际的命令会被包装为 `sh -c` 的参数的形式进行执行。比如：
`CMD echo $HOME`

在实际执行中，会将其变更为：

```
CMD [ "sh", "-c", "echo $HOME" ]
```

```
# 错误
CMD service nginx start

# 正确
CMD ["nginx", "-g", "daemon off;"]
```

### ENTRYPOINT 入口点

`ENTRYPOINT` 的格式和 `RUN` 指令格式一样，分为 `exec` 格式和 `shell` 格式。

`ENTRYPOINT` 的目的和 `CMD` 一样，都是在指定容器启动程序及参数。`ENTRYPOINT` 在运行时也可以替代，不过比 `CMD` 要略显繁琐，需要通过 `docker run` 的参数 `--entrypoint` 来指定。

当指定了 `ENTRYPOINT` 后，`CMD` 的含义就发生了改变，不再是直接的运行其命令，而是将 `CMD` 的内容作为参数传给 `ENTRYPOINT` 指令

`ENTRYPOINT`的两种用法

```
ENTRYPOINT [ "curl", "-s", "https://ip.cn" ]

docker run myip -i # 相当于在后面加了参数 不会改变原来的命令
```

```
ENTRYPOINT ["docker-entrypoint.sh"] #此脚本要add进去并且添加执行权限
CMD [ "redis-server" ]

# 执行时候就是相当于执行docker-entrypoint.sh redis-server
# 相当于带参数的脚本 比如 mysql 类的数据库，可能需要一些数据库配置、初始化的工作，这些工作要在最终的 mysql 服务器运行之前解决。
```

### ENV 设置环境变量

格式有两种：

- `ENV <key> <value>`
- `ENV <key1>=<value1> <key2>=<value2>...`

这个指令很简单，就是设置环境变量而已，无论是后面的其它指令，如 RUN，还是运行时的应用，都可以直接使用这里定义的环境变量。

```
ENV VERSION=1.0 DEBUG=on \
    NAME="Happy Feet" #有空格用引号
```

### ARG 构建参数

### VOLUME 定义匿名卷

### EXPOSE 暴露端口

格式为 `EXPOSE <端口1> [<端口2>...]`。

`EXPOSE` 指令是声明运行时容器提供服务端口，这只是一个声明，在运行时并不会因为这个声明应用就会开启这个端口的服务。在 Dockerfile 中写入这样的声明有两个好处，一个是帮助镜像使用者理解这个镜像服务的守护端口，以方便配置映射；另一个用处则是在运行时使用随机端口映射时，也就是 `docker run -P` 时，会自动随机映射 `EXPOSE` 的端口。

### WORKDIR 指定工作目录

格式为 `WORKDIR <工作目录路径>`。

使用 `WORKDIR` 指令可以来指定工作目录（或者称为当前目录），以后各层的当前目录就被改为指定的目录，如该目录不存在，`WORKDIR` 会帮你建立目录。

### USER 指定当前用户

格式：`USER <用户名>[:<用户组>]`

`USER` 指令和 `WORKDIR` 相似，都是改变环境状态并影响以后的层。`WORKDIR` 是改变工作目录，`USER` 则是改变之后层的执行 `RUN`, `CMD` 以及 `ENTRYPOINT` 这类命令的身份。

当然，和 `WORKDIR` 一样，`USER` 只是帮助你切换到指定用户而已，这个用户必须是事先建立好的，否则无法切换。

```
RUN groupadd -r redis && useradd -r -g redis redis
USER redis
RUN [ "redis-server" ]
```

如果以 `root` 执行的脚本，在执行期间希望改变身份，比如希望以某个已经建立好的用户来运行某个服务进程，不要使用 `su` 或者 `sudo`，这些都需要比较麻烦的配置，而且在 `TTY` 缺失的环境下经常出错。建议使用 `gosu`。

```
# 建立 redis 用户，并使用 gosu 换另一个用户执行命令
RUN groupadd -r redis && useradd -r -g redis redis
# 下载 gosu
RUN wget -O /usr/local/bin/gosu "https://github.com/tianon/gosu/releases/download/1.7/gosu-amd64" \
    && chmod +x /usr/local/bin/gosu \
    && gosu nobody true
# 设置 CMD，并以另外的用户执行
CMD [ "exec", "gosu", "redis", "redis-server" ]
```

为什么要用 gosu？

- gosu 启动命令时只有一个进程，所以 docker 容器启动时使用 gosu，那么该进程可以做到 PID 等于 1；
- sudo 启动命令时先创建 sudo 进程，然后该进程作为父进程去创建子进程，1 号 PID 被 sudo 进程占据；

### HEALTHCHECK 健康检查

### ONBUILD 为他人作嫁衣裳

## 镜像优化

### 第一步优化：使用轻量化基础镜像

### 第二步优化：多阶段构建
