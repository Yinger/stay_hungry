# 行编辑器
* Vim 和 sed，AWK 的区别
  * 交互式与非交互式
  * 文件操作模式与行操作模式

## sed
* 基本用法
  * sed 一般用于对文本内容做替换
  ```
  sed '/user1/s/user1/u1' /etc/passwd
  ```
## awk
* 基本用法
  * awk 一般用于对文本内容进行统计，按需要的格式进行输出
    * cut 命令： ``` cut -d : -f 1 /etc/passwd ``` 
    * awk 命令：``` awk -F : '/wd$/{print$1}' /etc/passwd ```
