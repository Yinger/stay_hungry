# LNMP
* 什么是 LNMP 环境
  * LAMP（Linux+Apache+PHP+MySQL）
  * LNMP（~~Apache~~ -> Nginx）
* LNMP 环境的搭建
  * MySQL安装
    * 可以使用 mariadb 替代
    * yum install mariadb mariadb-server
    * 修改默认编码
    ```
    vim /etc/my.cnf
    character_set_server=utf8
    init_connect='SET NAMES utf8'
    :wq
    ```
    * systemctl start mariadb.service
    ```
    # 进入mysql
    mysql
    show variables like '%character_set%';
    quit
    ```
* php 安装
  * yum install php-mysql php-fpm
  ```
  systemctl start php-frm.service
  ps -ef | grep php
  ```
* Nginx 配置
  ```
  cd /usr/local/openresty/nginx/
  ls
  cd conf
  ls
  vim nginx.conf

  location ~ \.php$ {
      root html;
      fastcgi_pass 127.0.0.1:9000;
      fastcgi_index index.php;
      fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
      include fastcgi_params;
  }
  :wq

  cd..
  cd sbin
  ./nginx -t
  ./nginx -s reload

  # php.index
  cd ..
  ls
  cd html/
  ls
  cat index.php
  <?php
  phpinfo();
  >

  # 请求
  ifconfig eth0 
  -> ip
  open browser 
  type ip
  type ip/index.php
  ```