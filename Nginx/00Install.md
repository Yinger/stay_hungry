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

# openresty

```bash
dnf -y install wget
wget 'https://openresty.org/package/centos/openresty.repo'
mv openresty.repo /etc/yum.repos.d/
dnf check-update
dnf -y install openresty
which openresty # /usr/bin/openresty
file `which openresty` # /usr/bin/openresty: symbolic link to /usr/local/openresty/nginx/sbin/nginx
openresty -V
systemctl start openresty
ps aux|grep nginx
```

```
gmake[2]: Leaving directory '/install/openresty-1.19.3.1/build/nginx-1.19.3'
gmake[1]: Leaving directory '/install/openresty-1.19.3.1/build/nginx-1.19.3'
mkdir -p /usr/local/openresty/site/lualib /usr/local/openresty/site/pod /usr/local/openresty/site/manifest
ln -sf /usr/local/openresty/nginx/sbin/nginx /usr/local/openresty/bin/openresty
```

# goaccess

https://goaccess.io/

```
goaccess host.access.log -o ../html/report.html --real-time-html --time-format='%H:%M:S' --date-format='%b/%d/%Y' --log-format=COMBINED --daemonize
```
