# Linux 的启动过程
* BIOS-MBR-BootLoader(grub)-kernel-systemd-系统初始化-shell

## BIOS
F2

## MBR
```
dd if=/dev/sda of=mbr.bin bs=446 count=1
hexdump -C mbr.bin

dd if=/dev/sda of=mbr2.bin bs=512 count=1
hexdump -C mbr2.bin | more

## 55 aa
```

## grub
```
cd /boot/grub2
ls

# 查看默认引导内核
grub2-editenv list

# 查看当前使用内核
uname -r
```

## kernel

## systemd
```
cd /etc/systemd/system
```

## 系统初始化
```
cd /usr/lib/systemd/system
```