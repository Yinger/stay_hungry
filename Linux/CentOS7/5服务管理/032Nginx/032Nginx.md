# Nginx
* Nginx 和 Web 服务介绍
  * Nginx（engine x）是一个高性能的 Web 和反向代理服务器
  * Nginx 支持 HTTP、HTTPS和电子邮件代理协议
  * OpenResty 是基于 Nginx 和 Lua 实现的 web 应用网关，集成了大量的第三方模块
* OpenResty 软件的下载和安装
  * yum-config-manager --add-repo https://openresty.org/package/centos/openresty.repo
  * yum install openresty
* OpenResty 的配置文件
  * /usr/local/openresty/nginx/conf/nginx.conf
  * service openresty start reload
  ```
  cd /usr/local/openresty
  ls
  cd nginx
  ls
  ls sbin
  ls conf/
  ls logs/
  ls html/

  ls
  cd conf/
  ls
  vim nginx.conf
  worker_processess # 占用多少个进程（占用几个CPU）
  worker_connections # 每一个进程支持多少个并发进程 over：HTTP503
  ```
* 使用 OpenResty 配置域名虚拟主机
  ```
  # 基于域名的虚拟主机
  pwd
  vim nginx.conf

  server {
      listen 8000;
      server_name www.servera.com;
      location / {
          root html/servera;
          index index.html index.htm;
      }
  }

  cd ..
  ls
  cd sbin
  ./nginx -t # 检测语法错误
  ./nginx

  ps -ef | grep nginx

  # 重新启动
  ./nginx -s stop
  ./nginx

  # reload(不用重启，现有连接不会断)
  ./nginx -s reload
  netstat -ntpl | grep nginx

  vim /etc/hosts
  127.0.0.1 www.servera.com www.serverb.com
  :wq

  # 进入html目录
  cd ..
  ls
  cd html/
  ls
  mkdir sservera serverb
  echo servera > servera/index.html
  echo serverb > serverb/index.html

  curl http://www.servera.com:8000
  -> servera

  ```