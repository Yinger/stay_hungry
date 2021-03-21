# Mac Docker 下配置 Nginx

- 搜索 Nginx 镜像列表

```
docker search nginx
```

- 拉取镜像

```
docker pull nginx
```

- 创建目录

```
mkdir -p /Users/linying/Documents/Kitematic/nginx /Users/linying/Documents/Kitematic/nginx/conf.d
```

- 创建配置文件

```
touch /Users/linying/Documents/Kitematic/nginx/nginx.conf
touch /Users/linying/Documents/Kitematic/nginx/conf.d/default.conf
```

- 创建 Nginx 临时容器，用于拷贝所需配置文件

```
docker run --name tmp-nginx-container -d nginx
```

- 拷贝 Nginx 配置

```
docker cp tmp-nginx-container:/etc/nginx/nginx.conf /Users/linying/Documents/Kitematic/nginx/nginx.conf
docker cp tmp-nginx-container:/etc/nginx/conf.d/default.conf /Users/linying/Documents/Kitematic/nginx/conf.d/default.conf
```

- 删除 Nginx 临时容器

```
docker rm -f tmp-nginx-container
```

- 创建 Nginx 容器

```
docker run --name nginx -p 80:80 -v /Users/linying/Documents/Kitematic/nginx/nginx.conf:/etc/nginx/nginx.conf -v /Users/linying/Documents/Kitematic/nginx/conf.d:/etc/nginx/conf.d -v /Users/linying/Documents/Kitematic/nginx/www:/www -d nginx
```

## 测试

- 拷贝 default.conf 为 test.conf ，并修改 server_name 和 root 目录

```
server {
    listen       80;
    server_name  www.test.com;

    location / {
        root   /www/test;
        index  index.html index.htm;
    }

    error_page   500 502 503 504  /50x.html;
    location = /50x.html {
        root   /usr/share/nginx/html;
    }
}
```

- 修改本地 /etc/hosts 文件，添加：

```
127.0.0.1 www.test.com
```

- 在本地 /Users/linying/Documents/Kitematic/nginx/www 下新建 test 目录，并编写一个 html 测试文件

```
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>

<body>
    <h1>Test for Nginx</h1>
</body>

</html>
```

- 重启 Nginx,访问 http://www.test.com/
