# Nginx学习笔记
## 一、什么是Nginx
Nginx是高性能web服务器，主要两大用途：
1. 静态资源服务器：访问html、图片等静态文件
2. 反向代理：接收客户端请求，转发给后端程序（Go、Java等），对外统一暴露80端口，隐藏后端真实端口。

## 二、Windows常用命令（cmd进入nginx根目录执行）
```
taskkill /f /im nginx.exe   # 杀死全部nginx进程
start nginx                 # 启动nginx
nginx -t                    # 检查配置文件语法是否正确
nginx -s reload             # 热重载配置，修改conf后执行
nginx -s stop               # 快速停止
```
## 三、核心配置文件：conf/nginx.conf
1. 配置块简单说明
worker_processes：工作进程数量
events：连接相关配置
http{}：http 服务总块，大部分配置写在这里面
server{}：代表一个虚拟站点，写端口、域名
listen：监听端口
server_name：访问域名 / 主机名
location / {}：匹配请求路径，写转发规则
2. 反向代理核心配置
nginx
server {
    listen 80;
    server_name localhost;

    location / {
        proxy_pass http://127.0.0.1:8080;  # 转发到后端服务地址
        proxy_set_header Host $host;
        proxy_set_header X‑Real‑IP $remote_addr; # 传递客户端真实IP
        proxy_http_version 1.1;
        proxy_set_header Connection "";
    }
}
proxy_pass：最重要指令，指定要转发的后端地址
## 四、反向代理工作流程
客户端请求 → Nginx (80 端口) → location 匹配路径 → proxy_pass 转发 → 后端 Go 服务 → 响应原路返回客户端。
## 五、常见报错
404：①路径匹配错误；②旧 nginx 进程没退出读旧配置
502 Bad Gateway：Nginx 找不到后端服务，后端程序没启动
nginx‑t 报错：conf 配置语法写错，括号、分号不能漏
## 六、负载均衡简单了解
可以配置多个后端地址，Nginx 把请求分发给多台后端，提高并发能力，常见策略：轮询、加权轮询、ip_hash。