# install

## 编译 Ngnix

- 1.下载 http://nginx.org/en/download.html (Stable version)
- 2.Configure
- 3.编译

```bash
yum install pcre pcre-devel -y
mkdir /nginx
cd /nginx
wget http://nginx.org/download/nginx-1.18.0.tar.gz
tar -xzf nginx-1.18.0.tar.gz
cd nginx-1.18.0/
cp -r contrib/vim* ~/.vim/
mkdir /nginx/install
./configure --prefix=/nginx/install
make
make install
```
