# 启动过程

ConfigureWebHostDefaults(eg:配置的组件，容器的组件)
⬇️
ConfigureHostConfiguration(监听端口，监听的 URL 地址)
⬇️
ConfigureAppConfiguration（嵌入自己的配置文件）
⬇️
ConfigureServices
ConfigureLogging （往容器中注入应用组件）
Startup
Startup.ConfigureServices
⬇️
Startup.Configure（注入中间件，处理 HttpContext 整个请求）
