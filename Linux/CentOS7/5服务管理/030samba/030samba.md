# samba

## Samba 和 NFS
* 常见共享服务的区别
  * 协议不同
  * 对操作系统的支持程度不同
  * 交互的便利性不同

## Samba
* Samba 服务的安装
  * yum install samba
* Samba 服务配置文件
  * /etc/samba/smb.conf
  ```
  [share]
    comment = my share
    path = /data/share
    read only = No
  ```
* Samba 用户的设置
  * smbpasswd 命令
    * -a 添加用户
    * -x 删除用户
  * pdbedit
    * -L 查看用户
* Samba 服务的启动和停止
  * systemctl start | stop smb.service
  * Linux 客户端
  ```
  mount -t cifs -o username=user1 //127.0.0.1/user1/mnt
  ```
  * Windows 客户端
    * 资源管理器访问共享
    * 映射网络驱动器
* NFS 服务的配置
* NFS 服务的启动和停止

```
cd /etc/samba/
ls

man 5 smb.conf

vim smb.conf

# 建立用户
# 1.建立系统用户
useradd user1
# 2.建立samba同名用户
smbpasswd -a user1

# 查看samba数据库里有哪些用户
pdbedit -L

# 删除samba用户
smbpasswd -x user1

# 启动samba
systemctl start smb.service

# Linux 访问samba
mount -t cifs -o username=user1 //127.0.0.1/user1 /mnt
# 输入samba密码
mount | tail -l
ls /mnt
ls /home/user1

# 共享share
umount /mnt
mount -t cifs -o username=user1 //127.0.0.1/share /mnt
ls /mnt
ls /data/share/
```